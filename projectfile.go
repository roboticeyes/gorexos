package gorexos

import (
	"encoding/json"
)

// ProjectFile structure of REXos
type ProjectFile struct {
	LastModified       string         `json:"lastModified"`
	ContentType        string         `json:"contentType"`
	Urn                string         `json:"urn"`
	FileSize           int            `json:"fileSize"`
	ContentHash        string         `json:"contentHash"`
	Name               string         `json:"name"`
	Type               string         `json:"type"`
	DataTransformation Transformation `json:"dataTransformation"`
	Links              struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		FileDownload struct {
			Href string `json:"href"`
		} `json:"file.download"`
	} `json:"_links"`
}

// GetProjectFilesByProjectSelfLink fetches all project files for a given self link
func GetProjectFilesByProjectSelfLink(handler RequestHandler, projectSelfLink string) ([]ProjectFile, error) {

	resp, err := handler.GetFullyQualified(projectSelfLink + "/projectFiles")
	if err != nil {
		return []ProjectFile{}, err
	}
	var result projectFilesHal
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return []ProjectFile{}, err
	}

	return result.Embedded.ProjectFiles, err
}

type projectFilesHal struct {
	Embedded struct {
		ProjectFiles []ProjectFile `json:"projectFiles"`
	} `json:"_embedded"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}
