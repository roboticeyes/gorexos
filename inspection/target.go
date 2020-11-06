package inspection

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/roboticeyes/gorexos"
)

const (
	apiTargets = "/inspection/v1/targets"
)

type Target struct {
	Urn             string           `json:"urn"`
	Name            string           `json:"name"`
	Owner           string           `json:"owner"`
	PortalReference *PortalReference `json:"portalReference"`
}

type PortalReference struct {
	LocalTransformation gorexos.Transformation `json:"localTransformation"`
	RexTagURL           string                 `json:"rexTagUrl"`
}

func GetTargets(handler gorexos.RequestHandler) ([]Target, error) {
	// TODO lisa
	return []Target{}, nil
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
