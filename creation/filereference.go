package creation

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/roboticeyes/gorexos"
)

const (
	apiFileReferences = "/filereferences"
)

type FileReference struct {
	FileUrn        string                 `json:"fileUrn"`
	Transformation gorexos.Transformation `json:"transformation"`
	Urn            string                 `json:"urn"`
}

// CreateFileReference creates a 3D reference in space which points to a project file being
// referenced by the fileUrn. The project file needs to be created first.
func CreateFileReference(handler gorexos.RequestHandler, projectUrn, fileUrn string, localTransformation gorexos.Transformation) (FileReference, error) {

	if fileUrn == "" {
		return FileReference{}, fmt.Errorf("fileUrn cannot be empty")
	}

	p := FileReference{
		FileUrn:        fileUrn,
		Transformation: localTransformation,
	}

	resp, err := handler.Post(apiProjects+"/"+projectUrn+apiFileReferences, p)

	if err != nil {
		return p, err
	}
	if resp.StatusCode() != http.StatusCreated {
		return p, fmt.Errorf("request responded with error code %s", resp.Status())
	}

	err = json.Unmarshal(resp.Body(), &p)
	return p, err
}
