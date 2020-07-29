// Package creation provides ...
package creation

import "github.com/roboticeyes/gorexos"

type ProjectsResponse struct {
	Page struct {
		Number        int `json:"number"`
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
	} `json:"page"`
	Projects []Project `json:"projects"`
}

type Project struct {
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

func GetProjects(handler gorexos.RequestHandler, userID string, p ProjectParameters) ([]Project, error) {

	projects := []Project{}

	return projects, nil
}
