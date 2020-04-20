package rexos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/roboticeyes/gorexos/pkg/math"
)

// Target is the top level entity for rex-inspection API
type Target struct {
	Name            string `json:"name"`
	Owner           string `json:"owner"`
	PortalReference struct {
		LocalTransformation math.Transformation `json:"localTransformation"`
		RexTagURL           string              `json:"rexTagUrl"`
	} `json:"portalReference"`
	Type string `json:"type"`
	Urn  string `json:"urn"`
}

type TargetFile struct {
	LocalFile   string     `json:"localFile"`
	Type        string     `json:"type"`
	Translation math.Vec3f `json:"translation"`
}

func DefaultTarget() Target {
	return Target{
		Name: uuid.New().String(),
		PortalReference: struct {
			LocalTransformation math.Transformation `json:"localTransformation"`
			RexTagURL           string              `json:"rexTagUrl"`
		}{
			LocalTransformation: math.NewTransformation(),
			RexTagURL:           "",
		},
	}
}

// GetUserStatistics fetches all statistics for a given user
func CreateTarget(handler RequestHandler, target Target) (Target, error) {

	resp, err := handler.Post(apiInspectionTargets, target)

	if err != nil {
		return target, err
	}

	err = json.Unmarshal(resp.Body(), &target)

	if resp.StatusCode() != http.StatusCreated {
		return target, fmt.Errorf("Failed, got back %d: %s", resp.StatusCode(), resp.Body())
	}
	return target, nil
}

// CreateTargetFile creates a new target file and returns the URN
func CreateTargetFile(handler RequestHandler, targetUrn string, targetFile TargetFile) error {

	body := struct {
		Rotation    math.Vec4f `json:"rotation"`
		Translation math.Vec3f `json:"translation"`
		Type        string     `json:"type"`
		Urn         string     `json:"urn"`
	}{
		Rotation:    math.Vec4f{X: 0, Y: 0, Z: 0, W: 1},
		Translation: targetFile.Translation,
		Type:        targetFile.Type,
	}

	resp, err := handler.Post(apiInspectionTargets+"/"+targetUrn+"/"+"files", body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(resp.Body(), &body)

	if resp.StatusCode() != http.StatusCreated {
		return fmt.Errorf("Failed, got back %d: %s", resp.StatusCode(), resp.Body())
	}

	err = uploadTargetFile(handler, targetUrn, body.Urn, targetFile.LocalFile)
	if err != nil {
		return fmt.Errorf("Failed to upload binary file, got back: %v", err)
	}

	return nil
}

func uploadTargetFile(handler RequestHandler, targetUrn, targetFileUrn, localFile string) error {

	r, err := os.Open(localFile)
	if err != nil {
		return fmt.Errorf("Cannot open local file: %s", localFile)
	}
	defer r.Close()

	url := apiInspectionTargets + "/" + targetUrn + "/files/" + targetFileUrn + "/data"
	resp, err := handler.PostMultipartFile(
		url,
		"file.rex",
		r,
	)

	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusCreated {
		return fmt.Errorf("Failed, got back %d: %s", resp.StatusCode(), resp.Body())
	}

	return nil
}
