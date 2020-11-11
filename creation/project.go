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
	"strconv"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/roboticeyes/gorexos"
)

const (
	apiProjects = "/creation/v1/projects"
)

type ProjectsPaged struct {
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
	Anchored        string
	IsOwnedBy       string
	IsReadSharedTo  string
	IsWriteSharedTo string
	Legacy          string
	IsHiddenFor     string
	IsFavoriteOf    string
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

func GetProjects(handler gorexos.RequestHandler, page int64, params *ProjectParameters) (ProjectsPaged, error) {
	paramString := ""
	if params != nil {
		if params.Anchored != "" {
			paramString += "&anchored=" + params.Anchored
		}
		if params.IsOwnedBy != "" {
			paramString += "&isOwnedBy=" + params.IsOwnedBy
		}
		if params.IsReadSharedTo != "" {
			paramString += "&isReadSharedTo=" + params.IsReadSharedTo
		}
		if params.IsWriteSharedTo != "" {
			paramString += "&isWriteSharedTo=" + params.IsWriteSharedTo
		}
		if params.Legacy != "" {
			paramString += "&legacy=" + params.Legacy
		}
		if params.IsHiddenFor != "" {
			paramString += "&isHiddenFor=" + params.IsHiddenFor
		}
		if params.IsFavoriteOf != "" {
			paramString += "&isFavoriteOf=" + params.IsFavoriteOf
		}
	}
	resp, err := handler.Get(apiProjects + "?page=" + strconv.FormatInt(page, 10) + paramString)
	if err != nil {
		return ProjectsPaged{}, err
	}
	if resp.StatusCode() != http.StatusOK {
		return ProjectsPaged{}, fmt.Errorf("request responded with error code %s", resp.Status())
	}
	var response ProjectsPaged
	err = json.Unmarshal(resp.Body(), &response)
	return response, nil
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
	if strings.ToLower(filepath.Ext(fileName)) == ".rex" {
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
