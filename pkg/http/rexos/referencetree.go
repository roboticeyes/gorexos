package rexos

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/breiting/tree"
)

// ReferenceTree contains everything for getting information about the reference tree in REXos
type ReferenceTree struct {
	ProjectUrn   string
	ProjectType  string
	Tree         *tree.Node
	References   []Reference
	ProjectFiles []ProjectFile
}

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
func GetReferenceTreeByProjectUrn(handler RequestHandler, projectUrn string) (ReferenceTree, error) {

	var refTree ReferenceTree

	// Get project
	project, err := GetProjectByUrn(handler, projectUrn)
	if err != nil {
		return refTree, err
	}
	refTree.ProjectUrn = project.Urn
	refTree.ProjectType = project.Type

	// Get project file list
	refTree.ProjectFiles, err = GetProjectFilesByProjectSelfLink(handler, project.SelfLink)

	// Get reference tree
	resp, err := handler.GetFullyQualified(project.SelfLink + "/rexReferences?projection=linkedList")
	if err != nil {
		return refTree, err
	}
	var halTree referenceTreeHal
	err = json.Unmarshal(resp.Body(), &halTree)
	if err != nil {
		return refTree, err
	}
	refTree.References = halTree.Embedded.RexReferences
	refTree.Tree, err = reconstructReferenceTreefromJSON(halTree.Embedded.RexReferences)
	return refTree, err
}

// Beautify modifies the tree and adds attributes to the graph
func (t *ReferenceTree) Beautify() {

	// add project node
	root := tree.NewNode(t.ProjectType + "\n" + strip(t.ProjectUrn))
	root.Attributes["shape"] = "octagon"
	root.Attributes["color"] = "gray"
	root.Children = append(root.Children, t.Tree)
	t.Tree = root

	for _, v := range t.References {

		node := tree.FindByID(t.Tree, strip(v.Urn))
		if node != nil {
			switch v.Type {
			case "portal":
				node.Attributes["color"] = "lightpink1"
				node.Attributes["shape"] = "doublecircle"
			case "group":
				node.Attributes["color"] = "darkolivegreen1"
			case "file":
				node.Attributes["color"] = "darkorange"
			}

			// shape is resource type
			// color is type von reference
			// file size und file name

			switch v.Category {
			case "activity":
				node.ID = "Activity\n" + node.ID
			case "inspection":
				node.ID = "Inspection\n" + node.ID
			case "track":
				node.ID = "Track\n" + node.ID
			case "file":
				node.ID = "File\n" + node.ID
			case "route":
				node.ID = "Route\n" + node.ID
			case "data":
				node.ID = "Data\n" + node.ID
			}
		}
	}
}

// WriteToDot gets the reference tree of the project and dumps out the structure as DOT file (graphviz)
func (t *ReferenceTree) WriteToDot(w io.Writer) error {
	return tree.WriteToDot(t.Tree, w)
}

func reconstructReferenceTreefromJSON(refs []Reference) (*tree.Node, error) {

	var relations []tree.Relation

	for _, v := range refs {
		relations = append(relations, tree.Relation{ID: strip(v.Urn), ParentID: strip(v.ParentReferenceUrn)})
	}

	return tree.Deserialize(relations)
}

// strip removes the prefix and only returns the ID
func strip(urn string) string {
	parts := strings.Split(urn, ":")
	return parts[len(parts)-1]
}
