package coreapi

import (
	"encoding/json"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/roboticeyes/gorexos"
	"github.com/tidwall/gjson"
)

// Project is the data container for top level REXos information
type Project struct {
	DateCreated string `json:"dateCreated"`
	LastUpdated string `json:"lastUpdated"`

	Urn      string `json:"urn"`
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	Public   bool   `json:"public"`
	Type     string `json:"type"`
	Scheme   string `json:"scheme"`
	SelfLink string

	NumberOfProjectFiles uint   `json:"numberOfProjectFiles"`
	TotalProjectFileSize uint   `json:"totalProjectFileSize"`
	RootRexReferenceKey  string `json:"rootRexReferenceKey"`
}

// Statistics is the user general statistics in REXos
type Statistics struct {
	NumberOfProjects              int `json:"numberOfProjects"`
	MaxNumberOfProjects           int `json:"maxNumberOfProjects"`
	TotalUsedDiskSpace            int `json:"totalUsedDiskSpace"`
	MaxTotalUsedDiskSpace         int `json:"maxTotalUsedDiskSpace"`
	NumberOfPublicShareActions    int `json:"numberOfPublicShareActions"`
	MaxNumberOfPublicShareActions int `json:"maxNumberOfPublicShareActions"`
}

// GetUserStatistics fetches all statistics for a given user
func GetUserStatistics(handler gorexos.RequestHandler, userID string) (Statistics, error) {

	var stats Statistics
	resp, err := handler.Get(apiStatisticsByUserID + userID)

	if err != nil {
		return stats, err
	}
	err = json.Unmarshal(resp.Body(), &stats)
	return stats, err
}

// GetAllProjectsByOwner fetches all project for a given owner
func GetAllProjectsByOwner(handler gorexos.RequestHandler, owner string) ([]Project, error) {

	var projects []Project
	resp, err := handler.Get(apiProjectByOwner + "owner=" + owner + "&projection=detailedList")

	if err != nil {
		return projects, err
	}
	projectList := []byte(gjson.GetBytes(resp.Body(), "_embedded.projects").Raw)
	err = json.Unmarshal(projectList, &projects)
	return projects, err
}

// GetProjectByUrn fetches the project by a given URN
func GetProjectByUrn(handler gorexos.RequestHandler, urn string) (Project, error) {

	var project Project
	resp, err := handler.Get(apiProjectsByUrn + "urn=" + urn + "&projection=detailedList")

	if err != nil {
		return project, err
	}
	err = json.Unmarshal(resp.Body(), &project)
	project.SelfLink = gorexos.GetSelfLinkFromHal(resp.Body())

	if project.Urn != urn {
		return project, fmt.Errorf("%s", gjson.Get(string(resp.Body()), "message").String())
	}
	return project, err
}

func (p Project) String() string {

	var writer strings.Builder
	tw := tabwriter.NewWriter(&writer, 0, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "Urn\t%s\n", p.Urn)
	fmt.Fprintf(tw, "Name\t%s\n", p.Name)
	fmt.Fprintf(tw, "Owner\t%s\n", p.Owner)
	fmt.Fprintf(tw, "Public\t%v\n", p.Public)
	fmt.Fprintf(tw, "Type\t%s\n", p.Type)
	fmt.Fprintf(tw, "Scheme\t%s\n", p.Scheme)
	fmt.Fprintf(tw, "RootRexReferenceKey\t%s\n", p.RootRexReferenceKey)
	fmt.Fprintf(tw, "\t\n")
	fmt.Fprintf(tw, "DateCreated\t%s\n", p.DateCreated)
	fmt.Fprintf(tw, "LastUpdated\t%s\n", p.LastUpdated)
	fmt.Fprintf(tw, "\t\n")
	fmt.Fprintf(tw, "NumberOfProjectFiles\t%d\n", p.NumberOfProjectFiles)
	fmt.Fprintf(tw, "Size\t%.2f mb\n", float32(p.TotalProjectFileSize)/1000.0/1000.0)
	fmt.Fprintf(tw, "\t\n")
	fmt.Fprintf(tw, "Link\t%s\n", p.SelfLink)
	tw.Flush()
	return writer.String()
}
