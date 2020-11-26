package creation

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/roboticeyes/gorexos"
)

const (
	apiGroupReferences = "/groupreferences"
	apiFileReferences  = "/filereferences"
)

type GroupReference struct {
	Transformation gorexos.TransformationWithScale `json:"transformation"`
	Urn            string                          `json:"urn"`
}

type FileReference struct {
	FileUrn        string                 `json:"fileUrn"`
	Transformation gorexos.Transformation `json:"transformation"`
	Urn            string                 `json:"urn"`
}

// CreateGroupReference creates a 3D group reference in space. This is required
// for a valid REXcad project. Every file reference needs to go underneath this
// group reference.
func CreateGroupReference(handler gorexos.RequestHandler, projectUrn string, localTransformation gorexos.TransformationWithScale) (GroupReference, error) {

	if projectUrn == "" {
		return GroupReference{}, fmt.Errorf("projectUrn cannot be empty")
	}

	p := GroupReference{
		Transformation: localTransformation,
	}

	resp, err := handler.Post(apiProjects+"/"+projectUrn+apiGroupReferences, p)

	if err != nil {
		return p, err
	}
	if resp.StatusCode() != http.StatusCreated {
		return p, fmt.Errorf("request responded with error code %s", resp.Status())
	}

	err = json.Unmarshal(resp.Body(), &p)
	return p, err
}

// CreateFileReference creates a 3D reference in space which points to a project file being
// referenced by the fileUrn. The group reference needs to be created first.
func CreateFileReference(handler gorexos.RequestHandler, projectUrn, groupUrn, fileUrn string, localTransformation gorexos.Transformation) (FileReference, error) {

	if projectUrn == "" || groupUrn == "" || fileUrn == "" {
		return FileReference{}, fmt.Errorf("projectUrn, groupUrn, and fileUrn cannot be empty")
	}

	p := FileReference{
		FileUrn:        fileUrn,
		Transformation: localTransformation,
	}

	resp, err := handler.Post(apiProjects+"/"+projectUrn+
		apiGroupReferences+"/"+groupUrn+
		apiFileReferences, p)

	if err != nil {
		return p, err
	}
	if resp.StatusCode() != http.StatusCreated {
		return p, fmt.Errorf("request responded with error code %s", resp.Status())
	}

	err = json.Unmarshal(resp.Body(), &p)
	return p, err
}
