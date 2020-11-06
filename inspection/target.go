package inspection

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/roboticeyes/gorexos"
)

const (
	apiTargets = "/inspection/v1/targets"
)

type TargetListElement struct {
	Urn                    string `json:"urn"`
	Name                   string `json:"name"`
	Owner                  string `json:"owner"`
	Type                   string `json:"type"`
	Public                 bool   `json:"public"`
	LastUpdated            string `json:"lastUpdated"`
	NrOfReadPermittedUsers string `json:"numberOfReadPermittedUsers"`
}

type Target struct {
	Urn             string           `json:"urn"`
	Name            string           `json:"name"`
	Owner           string           `json:"owner"`
	PortalReference *PortalReference `json:"portalReference"`
}

type TargetsPaged struct {
	Page struct {
		Number        int `json:"number"`
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
	} `json:"page"`
	Targets []TargetListElement `json:"targets"`
}

type PortalReference struct {
	LocalTransformation gorexos.Transformation `json:"localTransformation"`
	RexTagURL           string                 `json:"rexTagUrl"`
}

func GetTargets(handler gorexos.RequestHandler, page int64) (TargetsPaged, error) {
	resp, err := handler.Get(apiTargets + "?page=" + strconv.FormatInt(page, 10))
	if err != nil {
		return TargetsPaged{}, err
	}
	if resp.StatusCode() != http.StatusOK {
		return TargetsPaged{}, fmt.Errorf("request responded with error code %s", resp.Status())
	}
	var response TargetsPaged
	err = json.Unmarshal(resp.Body(), &response)
	return response, nil
}

func CreateTarget(handler gorexos.RequestHandler, name string, portalReference *PortalReference) (Target, error) {

	target := Target{
		Name:            name,
		PortalReference: portalReference,
	}
	resp, err := handler.Post(apiTargets, target)

	if err != nil {
		return target, err
	}
	if resp.StatusCode() != http.StatusCreated {
		return target, fmt.Errorf("request responded with error code %s", resp.Status())
	}

	err = json.Unmarshal(resp.Body(), &target)
	return target, err
}

func DeleteTarget(handler gorexos.RequestHandler, urn string) error {

	resp, err := handler.Delete(apiTargets+"/"+urn, nil)

	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("request responded with error code %s", resp.Status())
	}
	return nil
}
