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
	apiFiles    = "/files"
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
	Name            string          `json:"name"`
	Owner           string          `json:"owner"`
	Urn             string          `json:"urn"`
	PortalReference PortalReference `json:"portalReference"`
}

type PortalReference struct {
	Key            string                 `json:"key"`
	Name           string                 `json:"name"`
	Positioned     bool                   `json:"positioned"`
	Urn            string                 `json:"urn"`
	Transformation gorexos.Transformation `json:"transformation"`
}

type ProjectFile struct {
	Name           string                          `json:"name"`
	Urn            string                          `json:"urn"`
	Transformation gorexos.TransformationWithScale `json:"transformation"`
	Type           string                          `json:"type"`
}

type ProjectFileDetail struct {
	ProjectFile
	ContentType  string `json:"contentType"`
	DownloadLink string `json:"downloadLink"`
	FileSize     int    `json:"fileSize"`
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
	resp, err := handler.Get(apiProjects + "?sort=lastUpdated&page=" + strconv.FormatInt(page, 10) + paramString)
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

func GetProjectFiles(handler gorexos.RequestHandler, projectUrn string) ([]ProjectFileDetail, error) {

	files := []ProjectFileDetail{}
	resp, err := handler.Get(apiProjects + "/" + projectUrn + apiFiles)
	if err != nil {
		return files, err
	}
	if resp.StatusCode() != http.StatusOK {
		return files, fmt.Errorf("request responded with error code %s", resp.Status())
	}
	response := struct {
		Files []ProjectFileDetail `json:"files"`
	}{}
	err = json.Unmarshal(resp.Body(), &response)
	return response.Files, nil
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

func CreateProject(handler gorexos.RequestHandler, name string, portalReference PortalReference) (Project, error) {

	project := Project{
		Name:            name,
		PortalReference: portalReference,
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

func UploadProjectFile(handler gorexos.RequestHandler, urn, fileName string, dataTransform *gorexos.TransformationWithScale) (ProjectFile, error) {

	// Make sure to create a valid transformation if nothing is applied
	t := gorexos.NewTransformationWithScale()
	if dataTransform != nil {
		t = *dataTransform
	}

	var pf ProjectFile

	r, err := os.Open(fileName)
	if err != nil {
		return pf, err
	}
	defer r.Close()

	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		return pf, fmt.Errorf("error reading file %v", err)
	}
	reader := bytes.NewReader(dat)

	var fileType string
	if strings.ToLower(filepath.Ext(fileName)) == ".rex" {
		fileType = "rex"
	} else {
		mime, err := mimetype.DetectReader(r)
		if err != nil {
			return pf, fmt.Errorf("Cannot detect MIME type: %v", err)
		}
		fileType = mime.String()
	}

	pf = ProjectFile{
		Name:           filepath.Base(fileName),
		Transformation: t,
		Type:           fileType,
	}

	// Create project file
	url := apiProjects + "/" + urn + "/files"
	resp, err := handler.Post(url, pf)

	if err != nil {
		return pf, err
	}
	if resp.StatusCode() != http.StatusCreated {
		return pf, fmt.Errorf("request responded with error code %s", resp.Status())
	}
	err = json.Unmarshal(resp.Body(), &pf)

	// Upload file content
	url = apiProjects + "/" + urn + "/files/" + pf.Urn + "/data"
	resp2, err := handler.PostMultipartFile(url, fileName, reader)
	if err != nil {
		return pf, err
	}

	if resp2.StatusCode() != http.StatusCreated {
		// delete file
		handler.Delete(apiProjects+"/"+urn+"/files/"+pf.Urn, nil)
		return pf, fmt.Errorf("request responded with error code %s", resp2.Status())
	}

	return pf, nil
}
