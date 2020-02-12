package commands

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/roboticeyes/gorexos/pkg/config"
	"github.com/roboticeyes/gorexos/pkg/http/rexos"
	"github.com/urfave/cli/v2"
)

// LoginCommand performs a login to the REXos platform with the given credentials
var LoginCommand = &cli.Command{
	Name:   "login",
	Usage:  "Login into REXos",
	Action: loginAction,
}

func loginAction(ctx *cli.Context) error {

	config := ctx.App.Metadata["config"].(*config.Config)

	app := tview.NewApplication()
	instanceList := tview.NewList()
	userList := tview.NewList()

	for k, i := range config.Instances {
		if len(i.Users) > 0 {
			instanceList.AddItem(i.Name, i.URL, rune(48+k), nil)
		}
	}

	var selectedInstance int

	instanceList.SetSelectedFunc(func(idx int, mainText string, secondaryText string, shortcut rune) {
		userList.Clear()
		selectedInstance = idx
		for k, u := range config.Instances[idx].Users {
			userList.AddItem(u.Name, u.ClientID, rune(48+k), nil)
		}
		app.SetRoot(userList, true)
	})

	userList.SetSelectedFunc(func(idx int, mainText string, secondaryText string, shortcut rune) {
		app.Stop()
		handler := rexos.NewRequestHandler(config.Instances[selectedInstance].URL, "/rex-gateway/api/v2")
		clientID := config.Instances[selectedInstance].Users[idx].ClientID
		clientSecret := config.Instances[selectedInstance].Users[idx].ClientSecret
		userID := handler.Authenticate(clientID, clientSecret)
		fmt.Println("UserID: ", userID)
	})

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc || event.Rune() == 'q' {
			app.Stop()
		}
		return event
	})

	if err := app.SetRoot(instanceList, true).Run(); err != nil {
		panic(err)
	}

	return nil
}
