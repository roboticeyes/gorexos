// Package creation provides ...
package creation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/roboticeyes/gorexos"
)

const (
	apiProjects = "/creation/v1/projects"
)

type ProjectsResponse struct {
	Page struct {
		Number        int `json:"number"`
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
	} `json:"page"`
	Projects []ProjectDescription `json:"projects"`
}

type ProjectDescription struct {
	Name                 string `json:"name"`
	NumberOfProjectFiles int    `json:"numberOfProjectFiles"`
	Public               bool   `json:"public"`
	TotalProjectFileSize int    `json:"totalProjectFileSize"`
	Type                 string `json:"type"`
	Urn                  string `json:"urn"`
}

type ProjectParameters struct {
	Anchored        bool
	IsOwnedBy       bool
	IsReadSharedTo  bool
	IsWriteSharedTo bool
}

type Project struct {
	Name   string `json:"name"`
	Owner  string `json:"owner"`
	Urn    string `json:"urn"`
	Anchor Anchor `json:"anchor"`
}

type Anchor struct {
	Key            string                 `json:"key"`
	Name           string                 `json:"name"`
	Positioned     bool                   `json:"positioned"`
	Urn            string                 `json:"urn"`
	Transformation gorexos.Transformation `json:"transformation"`
}

type ProjectFile struct {
	Name           string                 `json:"name"`
	Urn            string                 `json:"urn"`
	Transformation gorexos.Transformation `json:"transformation"`
	Type           string                 `json:"type"`
}

func GetProjects(handler gorexos.RequestHandler, userID string, p ProjectParameters) ([]ProjectDescription, error) {

	resp, err := handler.Get(apiProjects)
	if err != nil {
		return []ProjectDescription{}, err
	}
	if resp.StatusCode() != http.StatusOK {
		return []ProjectDescription{}, fmt.Errorf("request responded with error code %s", resp.Status())
	}
	var response ProjectsResponse
	err = json.Unmarshal(resp.Body(), &response)
	return response.Projects, nil
}

func DeleteProject(handler gorexos.RequestHandler, urn string) error {

	resp, err := handler.Delete(apiProjects+"/"+urn, nil)

	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("request responded with error code %s", resp.Status())
	}
	return nil
}

func CreateProject(handler gorexos.RequestHandler, name string) (Project, error) {

	project := Project{
		Name: name,
		Anchor: Anchor{
			Transformation: gorexos.NewTransformation(),
		},
	}
	resp, err := handler.Post(apiProjects, project)

	if err != nil {
		return project, err
	}
	if resp.StatusCode() != http.StatusCreated {
		return project, fmt.Errorf("request responded with error code %s", resp.Status())
	}

	err = json.Unmarshal(resp.Body(), &project)
	return project, err
}

func UploadProjectFile(handler gorexos.RequestHandler, urn, fileName string, transformation gorexos.Transformation) error {

	r, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer r.Close()

	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("error reading file %v", err)
	}
	reader := bytes.NewReader(dat)

	var fileType string
	if strings.ToLower(filepath.Ext(fileName)) == "rex" {
		fileType = "rex"
	} else {
		mime, err := mimetype.DetectReader(r)
		if err != nil {
			return fmt.Errorf("Cannot detect MIME type: %v", err)
		}
		fileType = mime.String()
	}

	pf := ProjectFile{
		Name:           filepath.Base(fileName),
		Transformation: transformation,
		Type:           fileType,
	}
	fmt.Printf("%+v\n", pf)

	// Create project file
	url := apiProjects + "/" + urn + "/files"
	resp, err := handler.Post(url, pf)

	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusCreated {
		return fmt.Errorf("request responded with error code %s", resp.Status())
	}
	err = json.Unmarshal(resp.Body(), &pf)
	fmt.Printf("%+v\n", pf)

	// Upload file content
	url = apiProjects + "/" + urn + "/files/" + pf.Urn + "/data"
	resp2, err := handler.PostMultipartFile(url, fileName, reader)
	if err != nil {
		return err
	}

	if resp2.StatusCode() != http.StatusCreated {
		// delete file
		handler.Delete(apiProjects+"/"+urn+"/files/"+pf.Urn, nil)
		return fmt.Errorf("request responded with error code %s", resp2.Status())
	}

	return nil
}
