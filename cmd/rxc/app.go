package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/roboticeyes/gorexos/pkg/config"
	"github.com/roboticeyes/gorexos/pkg/http/rexos"
	"github.com/urfave/cli/v2"
)

// App is the main app
type App struct {
	app              *tview.Application
	status           *tview.TextView
	counter          int
	clientID         string
	clientSecret     string
	authURL          string
	session          rexos.Session
	selectedUser     string
	selectedInstance string
}

const refreshInterval = time.Second

// UIRunner wrapps the function to run the UI
type UIRunner interface {
	Run() error
}

func createList(title string) *tview.List {
	l := tview.NewList()
	l.SetBorder(true)
	l.SetTitle(" " + title + " ")
	l.SetBorderPadding(1, 1, 1, 1)
	return l
}

func (a *App) getStatusText() string {
	s := "You are logged in REXos\n\n"
	s += "   Instance:      " + a.selectedInstance + "\n"
	s += "   User:          " + a.selectedUser + "\n"
	s += "   Token refresh: "
	return fmt.Sprintf("%s%d", s, a.counter)
}

func (a *App) updateTime() {
	for {
		time.Sleep(refreshInterval)
		if !a.session.Valid() {
			continue
		}
		a.app.QueueUpdateDraw(func() {
			a.status.SetText(a.getStatusText())
		})
		a.counter--
		if a.counter <= 10 { // 10 seconds before token expires
			a.refreshToken()
			a.counter = a.session.ExpiresIn
		}
	}
}

// NewApp creates a new app
func NewApp(ctx *cli.Context) UIRunner {

	cfg := ctx.App.Metadata["config"].(*config.Config)

	a := App{
		app: tview.NewApplication(),
	}

	a.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc || event.Rune() == 'q' {
			a.app.Stop()
		}
		return event
	})

	listInstance := createList("Select instance")
	listUser := createList("Select user")
	a.status = tview.NewTextView()
	a.status.SetText("You are now logged in ...")

	for k, i := range cfg.Instances {
		if len(i.Users) > 0 {
			listInstance.AddItem(i.Name, i.URL, rune(48+k), nil)
		}
	}

	var selectedInstance int

	listInstance.SetSelectedFunc(func(idx int, mainText string, secondaryText string, shortcut rune) {
		listUser.Clear()
		selectedInstance = idx
		a.selectedInstance = cfg.Instances[idx].Name
		for k, u := range cfg.Instances[idx].Users {
			listUser.AddItem(u.Name, u.ClientID, rune(48+k), nil)
		}
		a.app.SetRoot(center(0, 0, listUser), true)
	})

	listUser.SetSelectedFunc(func(idx int, mainText string, secondaryText string, shortcut rune) {
		a.clientID = cfg.Instances[selectedInstance].Users[idx].ClientID
		a.clientSecret = cfg.Instances[selectedInstance].Users[idx].ClientSecret
		a.selectedUser = cfg.Instances[selectedInstance].Users[idx].Name
		a.authURL = cfg.Instances[selectedInstance].URL
		a.app.SetRoot(center(0, 0, a.status), true)
		a.refreshToken()
	})

	a.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc || event.Rune() == 'q' {
			a.app.Stop()
		}
		return event
	})

	go a.updateTime()
	a.app.SetRoot(center(0, 0, listInstance), true)
	return &a
}

func (a *App) refreshToken() {
	handler := rexos.NewRequestHandler()
	a.session = handler.Authenticate(a.authURL, a.clientID, a.clientSecret)
	a.counter = a.session.ExpiresIn
	f, _ := os.Create(filepath.Join(config.UserRexOSDir(), ".session.json"))
	defer f.Close()
	a.session.Write(f)
}

// Run starts the user interface
func (a *App) Run() error {
	return a.app.Run()
}

func center(width, height int, p tview.Primitive) tview.Primitive {
	return tview.NewFlex().
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(p, height, 1, true).
			AddItem(tview.NewBox(), 0, 1, false), width, 1, true).
		AddItem(tview.NewBox(), 0, 1, false)
}
