package creation

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/roboticeyes/gorexos"
)

const (
	apiPois = "/pois"
)

type Poi struct {
	FileUrn        string                 `json:"fileUrn"`
	Transformation gorexos.Transformation `json:"transformation"`
	Urn            string                 `json:"urn"`
}

// CreatPoi creates a 3D reference in space which points to a project file being
// referenced by the fileUrn. The project file needs to be created first.
func CreatePoi(handler gorexos.RequestHandler, projectUrn, fileUrn string, translation gorexos.Vec3f, rotation gorexos.Vec4f) (Poi, error) {

	if fileUrn == "" {
		return Poi{}, fmt.Errorf("fileUrn cannot be empty")
	}

	p := Poi{
		FileUrn: fileUrn,
		Transformation: gorexos.Transformation{
			Translation: translation,
			Rotation:    rotation,
			Scale:       1.0,
		},
	}

	resp, err := handler.Post(apiProjects+"/"+projectUrn+apiPois, p)

	if err != nil {
		return p, err
	}
	if resp.StatusCode() != http.StatusCreated {
		return p, fmt.Errorf("request responded with error code %s", resp.Status())
	}

	err = json.Unmarshal(resp.Body(), &p)
	return p, err
}
