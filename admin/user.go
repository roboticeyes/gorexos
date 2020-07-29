// Package admin provides ...
package admin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/roboticeyes/gorexos"
)

const (
	apiUsers = "/administration/v1/users?sort=lastLogin,desc"
)

type UsersResponse struct {
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
	Users []User `json:"users"`
}

type User struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Locked      bool   `json:"locked"`
	Disabled    bool   `json:"disabled"`
	DateCreated string `json:"dateCreated"`
	LastUpdated string `json:"lastUpdated"`
	LastLogin   string `json:"lastLogin"`
	Urn         string `json:"urn"`
}

func GetUsers(handler gorexos.RequestHandler) ([]User, error) {

	users := []User{}

	pageIdx := 0
	for {
		var usersResponse UsersResponse
		query := fmt.Sprintf("%s&page=%d", apiUsers, pageIdx)
		resp, err := handler.Get(query)

		if err != nil {
			return usersResponse.Users, err
		}
		if resp.StatusCode() != http.StatusOK {
			return users, fmt.Errorf("request responded with error code %s", resp.Status())
		}

		err = json.Unmarshal(resp.Body(), &usersResponse)
		users = append(users, usersResponse.Users...)

		nrOfPages := usersResponse.Page.TotalPages
		if nrOfPages == 0 {
			break
		}
		if usersResponse.Page.Number > nrOfPages-1 {
			break
		} else {
			pageIdx++
		}
	}

	return users, nil
}
