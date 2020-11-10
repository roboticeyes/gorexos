// Package admin provides ...
package admin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/roboticeyes/gorexos"
)

const (
	apiUsersPlain  = "/administration/v1/users"
	apiUsers       = "/administration/v1/users?sort=lastLogin,desc"
	apiSearchUsers = "/administration/v1/users/search/findAllByNameOrEmailContaining?partialString="
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

type UserDetail struct {
	Urn                           string `json:"urn"`
	Address                       string `json:"address"`
	City                          string `json:"city"`
	Company                       string `json:"company"`
	Country                       string `json:"country"`
	Countryname                   string `json:"countryname"`
	DateCreated                   string `json:"dateCreated"`
	Disabled                      bool   `json:"disabled"`
	Email                         string `json:"email"`
	FirstName                     string `json:"firstName"`
	LastLogin                     string `json:"lastLogin"`
	LastName                      string `json:"lastName"`
	LastUpdated                   string `json:"lastUpdated"`
	Locked                        bool   `json:"locked"`
	MaxNumberOfProjects           int    `json:"maxNumberOfProjects"`
	MaxNumberOfPublicShareActions int    `json:"maxNumberOfPublicShareActions"`
	MaxTotalUsedDiskSpace         int    `json:"maxTotalUsedDiskSpace"`
	NumberOfProjects              int    `json:"numberOfProjects"`
	NumberOfPublicShareActions    int    `json:"numberOfPublicShareActions"`
	State                         string `json:"state"`
	TotalUsedDiskSpace            int    `json:"totalUsedDiskSpace"`
	UserID                        string `json:"userID"`
	Zip                           string `json:"zip"`
	Subscriptions                 []struct {
		ActivationDate string `json:"activationDate"`
		ExpirationDate string `json:"expirationDate"`
		Name           string `json:"name"`
		Urn            string `json:"urn"`
	} `json:"subscriptions"`
}

func GetUserDetails(handler gorexos.RequestHandler, urn string) (UserDetail, error) {

	query := fmt.Sprintf("%s/%s", apiUsersPlain, urn)
	resp, err := handler.Get(query)

	if err != nil {
		return UserDetail{}, err
	}
	if resp.StatusCode() != http.StatusOK {
		return UserDetail{}, fmt.Errorf("request responded with error code %s", resp.Status())
	}

	result := UserDetail{}
	err = json.Unmarshal(resp.Body(), &result)
	return result, nil
}

func FindUsers(handler gorexos.RequestHandler, partialNameOrEmail string) ([]User, error) {

	users := []User{}

	pageIdx := 0
	for {
		var usersResponse UsersResponse
		query := fmt.Sprintf("%s%s&page=%d", apiSearchUsers, partialNameOrEmail, pageIdx)
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
