package rexos

import (
	"encoding/json"
	"strings"

	"github.com/breiting/tree"
)

// referenceTreeHal is a serialized linked list of all references belonging to one project.
// this can be retrieved by /projects/:id/rexReferences?projection=linkedList
type referenceTreeHal struct {
	Embedded struct {
		RexReferences []Reference `json:"rexReferences"`
	} `json:"_embedded"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

// GetReferenceTreeByProjectUrn returns the full reference tree for the given project URN
func GetReferenceTreeByProjectUrn(handler RequestHandler, projectUrn string) (*tree.Node, error) {

	project, err := GetProjectByUrn(handler, projectUrn)
	if err != nil {
		return nil, err
	}

	resp, err := handler.GetFullyQualified(project.SelfLink + "/rexReferences?projection=linkedList")
	if err != nil {
		return nil, err
	}

	var halTree referenceTreeHal
	err = json.Unmarshal(resp.Body(), &halTree)
	if err != nil {
		return nil, err
	}

	return reconstructReferenceTreefromJSON(halTree.Embedded.RexReferences)
}

// strip removes the prefix and only returns the ID
func strip(urn string) string {
	parts := strings.Split(urn, ":")
	return parts[len(parts)-1]
}

func reconstructReferenceTreefromJSON(refs []Reference) (*tree.Node, error) {

	var relations []tree.Relation

	for _, v := range refs {
		relations = append(relations, tree.Relation{ID: strip(v.Urn), ParentID: strip(v.ParentReferenceUrn)})
	}

	return tree.Deserialize(relations)
}
