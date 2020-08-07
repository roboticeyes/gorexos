package admin

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/roboticeyes/gorexos"
)

const (
	apiCoreLicenses = "/api/v2/licenses"
	apiLicenses     = "/administration/v1/licenses"
)

type LicensesResponse struct {
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
	Licenses []License `json:"licenses"`
}

// LicenseItem is a container for license items
type LicenseItem struct {
	DateCreated  string  `json:"dateCreated" example:"2019-07-09T13:03:57.575+0000"`
	LastUpdated  string  `json:"lastUpdated" example:"2019-07-09T13:03:57.575+0000"`
	Key          string  `json:"key" example:"positionProjectFileBudget"`
	ValueString  string  `json:"valueString"`
	ValueLong    int64   `json:"valueLong" example:"20"`
	ValueDouble  float64 `json:"valueDouble"`
	ValueBoolean bool    `json:"valueBoolean"`
	ValueDate    string  `json:"valueDate"`
	Urn          string  `json:"urn" example:"robotic-eyes:license-item:1102"`
}

// License is a container for licenses
type License struct {
	DateCreated    string        `json:"dateCreated" example:"2019-07-09T13:03:57.557+0000"`
	LastUpdated    string        `json:"lastUpdated" example:"2019-07-09T13:03:57.557+0000"`
	Key            string        `json:"key" example:"rex-addon-positioning-tags-20"`
	Name           string        `json:"name" example:"20 positioning tags"`
	Description    string        `json:"description" example:"addon sample description"`
	ExpirationDate string        `json:"expirationDate" example:"2019-07-09T13:03:57.557+0000"`
	Active         bool          `json:"active"`
	DefaultLicense bool          `json:"defaultLicense"`
	Urn            string        `json:"urn" example:"robotic-eyes:license:1101"`
	LicenseItems   []LicenseItem `json:"licenseItems"`
}

// CreateLicense creates a new REX license where the license is provided by a JSON file reader
func CreateLicense(handler gorexos.RequestHandler, r io.Reader) error {

	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	handler.Post(apiCoreLicenses, body)
	return nil
}

// DeleteLicense deletes a REX license permanently!
func DeleteLicense(handler gorexos.RequestHandler, urn string) error {

	resp, err := handler.Delete(apiCoreLicenses, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("Received unexpected response code %d", resp.StatusCode())
	}
	return nil
}

// GetLicenses gets all the available licenses
func GetLicenses(handler gorexos.RequestHandler) ([]License, error) {

	licenses := []License{}

	pageIdx := 0
	for {
		var licResponse LicensesResponse
		query := fmt.Sprintf("%s?page=%d", apiLicenses, pageIdx)
		resp, err := handler.Get(query)

		if err != nil {
			return licResponse.Licenses, err
		}
		if resp.StatusCode() != http.StatusOK {
			return licenses, fmt.Errorf("request responded with error code %s", resp.Status())
		}

		err = json.Unmarshal(resp.Body(), &licResponse)
		licenses = append(licenses, licResponse.Licenses...)

		nrOfPages := licResponse.Page.TotalPages
		if nrOfPages == 0 {
			break
		}
		if licResponse.Page.Number > nrOfPages-1 {
			break
		} else {
			pageIdx++
		}
	}

	return licenses, nil
}
