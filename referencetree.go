package gorexos

import (
	"encoding/json"
	"fmt"
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

// GetTransformations fetches all transformations for all references and project files
func (t *ReferenceTree) GetTransformations(handler RequestHandler) error {
	for i, ref := range t.References {
		resp, err := handler.Get(apiReferences + "/" + StripPrefix(ref.Urn))

		if err != nil {
			return err
		}

		var refResponse Reference
		err = json.Unmarshal(resp.Body(), &refResponse)
		t.References[i].LocalTransformation = refResponse.LocalTransformation
		t.References[i].WorldTransformation = refResponse.WorldTransformation
	}

	for i, projectFile := range t.ProjectFiles {
		resp, err := handler.Get(apiProjectFiles + "/" + StripPrefix(projectFile.Urn))

		if err != nil {
			return err
		}

		var projectFileResponse ProjectFile
		err = json.Unmarshal(resp.Body(), &projectFileResponse)
		t.ProjectFiles[i].DataTransformation = projectFileResponse.DataTransformation
	}
	return nil
}

// Beautify modifies the tree and adds attributes to the graph
func (t *ReferenceTree) Beautify() {

	// add project node
	root := tree.NewNode(t.ProjectType + "\n" + StripPrefix(t.ProjectUrn))
	root.Attributes["shape"] = "octagon"
	root.Attributes["color"] = "azure2"
	root.Children = append(root.Children, t.Tree)
	t.Tree = root

	for _, v := range t.References {

		node := tree.FindByID(t.Tree, StripPrefix(v.Urn))
		if node != nil {
			switch v.Type {
			case "portal":
				node.Attributes["color"] = "goldenrod1"
				node.Attributes["shape"] = "circle"
			case "root":
				node.Attributes["shape"] = "doublecircle"
				node.Attributes["color"] = "firebrick"
			// case "group":
			// see below with categories
			// 	node.Attributes["color"] = "darkolivegreen1"
			case "file":
				node.Attributes["color"] = "dodgerblue3"
			}

			switch v.Category {
			case "activity":
				node.ID = "Activity\n" + node.ID
				node.Attributes["color"] = "chocolate2"
			case "inspection":
				node.ID = "Inspection\n" + node.ID
				node.Attributes["color"] = "seagreen1"
			case "track":
				node.ID = "Track\n" + node.ID
				node.Attributes["color"] = "orchid3"
			case "file":
				node.ID = "File\n" + node.ID
			case "route":
				node.ID = "Route\n" + node.ID
				node.Attributes["color"] = "hotpink4"
			case "data":
				node.ID = "Data\n" + node.ID
				node.Attributes["color"] = "aquamarine4"
			}

			// attach project file
			for _, p := range t.ProjectFiles {
				if p.Urn == v.ProjectFileUrn {
					// found project file
					fileSize := fmt.Sprintf("~%.2fmb", float32(p.FileSize)/1000.0/1000.0)
					pfNode := &tree.Node{
						ID:         p.Type + "\n" + StripPrefix(p.Urn) + "\n" + fileSize,
						Name:       p.Name,
						Attributes: make(map[string]string),
					}
					pfNode.Attributes["shape"] = "box"
					pfNode.Attributes["color"] = "powderblue"
					node.Children = append(node.Children, pfNode)
				}
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
		relations = append(relations, tree.Relation{ID: StripPrefix(v.Urn), ParentID: StripPrefix(v.ParentReferenceUrn)})
	}

	return tree.Deserialize(relations)
}

// StripPrefix removes the prefix of the URN and only returns the ID
func StripPrefix(urn string) string {
	parts := strings.Split(urn, ":")
	return parts[len(parts)-1]
}
