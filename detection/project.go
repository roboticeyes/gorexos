// Package detection provides ...
package detection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/roboticeyes/gorexos"
)

const (
	apiProjects = "/detection/v1/projects"
)

type ProjectBase struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
	Urn   string `json:"urn"`

	PortalReference *PortalReference `json:"portalReference"`
}

type Project struct {
	ProjectBase

	Type        string `json:"type"`
	LastUpdated string `json:"lastUpdated"`
	Rooms       []Room `json:"rooms"`
}

type RoomBase struct {
	Name            string           `json:"name"`
	Urn             string           `json:"urn"`
	PhysicalObjects []PhysicalObject `json:"physicalObjects"`
}

type Room struct {
	RoomBase

	Geometries []Geometry `json:"geometries"`
}

type Geometry struct {
	Name           string                 `json:"name"`
	Urn            string                 `json:"urn"`
	Transformation gorexos.Transformation `json:"transformation"`
	FileUrn        string                 `json:"fileUrn"`
}

type PhysicalObject struct {
	Name           string                 `json:"name"`
	Category       string                 `json:"category"`
	Authored       string                 `json:"authored"`
	Confidence     float32                `json:"confidence"`
	Transformation gorexos.Transformation `json:"transformation"`
}

// PortalReference is the entry to a project
type PortalReference struct {
	RexTagURL           string                 `json:"rexTagUrl,omitempty"`
	LocalTransformation gorexos.Transformation `json:"localTransformation,omitempty"`
}

func GetProjectByUrn(handler gorexos.RequestHandler, urn string) (Project, error) {

	p := Project{}
	resp, err := handler.Get(apiProjects + "/" + urn)
	if err != nil {
		return p, err
	}
	if resp.StatusCode() != http.StatusOK {
		return p, fmt.Errorf("request responded with error code %s", resp.Status())
	}
	err = json.Unmarshal(resp.Body(), &p)
	return p, nil
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

func CreateProject(handler gorexos.RequestHandler, p ProjectBase) (ProjectBase, error) {

	if p.PortalReference == nil {
		p.PortalReference = &PortalReference{
			LocalTransformation: gorexos.NewTransformation(),
		}
	}
	if p.Name == "" {
		return p, fmt.Errorf("Project name cannot be empty")
	}
	resp, err := handler.Post(apiProjects, p)

	if err != nil {
		return p, err
	}
	if resp.StatusCode() != http.StatusCreated {
		return p, fmt.Errorf("request responded with error code %s", resp.Status())
	}

	err = json.Unmarshal(resp.Body(), &p)
	return p, err
}

func CreateRoom(handler gorexos.RequestHandler, projectUrn string, r RoomBase) (RoomBase, error) {

	if projectUrn == "" {
		return r, fmt.Errorf("ProjectUrn cannot be empty")
	}
	// Generate a random name for the room if not set
	if r.Name == "" {
		r.Name = uuid.New().String()
	}
	resp, err := handler.Post(apiProjects+"/"+projectUrn+"/rooms", r)

	if err != nil {
		return r, err
	}
	if resp.StatusCode() != http.StatusCreated {
		return r, fmt.Errorf("request responded with error code %s", resp.Status())
	}

	err = json.Unmarshal(resp.Body(), &r)
	return r, err
}

// CreateRoomGeometery uploads a new REXfile representing the room. Only
// REXfiles are allowed!
func CreateRoomGeometry(handler gorexos.RequestHandler, projectUrn, roomUrn string, r io.Reader, transform gorexos.Transformation) (Geometry, error) {

	geom := Geometry{}

	dat, err := ioutil.ReadAll(r)
	if err != nil {
		return geom, fmt.Errorf("error reading content from file %v", err)
	}
	reader := bytes.NewReader(dat)

	geom.Transformation = transform
	geom.Name = uuid.New().String()

	url := apiProjects + "/" + projectUrn + "/rooms/" + roomUrn + "/geometries"
	resp, err := handler.Post(url, geom)

	if err != nil {
		return geom, err
	}
	if resp.StatusCode() != http.StatusCreated {
		return geom, fmt.Errorf("request responded with error code %s", resp.Status())
	}
	err = json.Unmarshal(resp.Body(), &geom)

	// Upload file content
	url = apiProjects + "/" + projectUrn + "/rooms/" + roomUrn + "/geometries/" + geom.Urn + "/data"
	resp2, err := handler.PostMultipartFile(url, geom.Name+".rex", reader)
	if err != nil {
		return geom, err
	}

	if resp2.StatusCode() != http.StatusCreated {
		// delete file
		handler.Delete(apiProjects+"/"+projectUrn+"/rooms/"+roomUrn+"/geometries/"+geom.Urn, nil)
		return geom, fmt.Errorf("request responded with error code %s", resp2.Status())
	}

	return geom, nil
}
