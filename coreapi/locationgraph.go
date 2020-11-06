package coreapi

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/breiting/tree"
	"github.com/roboticeyes/gorexos"
)

// LocationGraph contains everything for getting information about the location graph in REXos
type LocationGraph struct {
	ProjectUrn   string
	ProjectType  string
	Tree         *tree.Node
	References   []Reference
	ProjectFiles []ProjectFile
}

// locationGraphHal is a serialized linked list of all references belonging to one project.
// this can be retrieved by /projects/:id/rexReferences?projection=linkedList
type locationGraphHal struct {
	Embedded struct {
		RexReferences []Reference `json:"rexReferences"`
	} `json:"_embedded"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

// GetLocationGraphByProjectUrn returns the full location graph for the given project URN
func GetLocationGraphByProjectUrn(handler gorexos.RequestHandler, projectUrn string) (LocationGraph, error) {

	var locTree LocationGraph

	// Get project
	project, err := GetProjectByUrn(handler, projectUrn)
	if err != nil {
		return locTree, err
	}
	locTree.ProjectUrn = project.Urn
	locTree.ProjectType = project.Type

	// Get project file list
	locTree.ProjectFiles, err = GetProjectFilesByProjectSelfLink(handler, project.SelfLink)

	// Get references
	resp, err := handler.GetFullyQualified(project.SelfLink + "/rexReferences?projection=linkedList")
	if err != nil {
		return locTree, err
	}
	var halTree locationGraphHal
	err = json.Unmarshal(resp.Body(), &halTree)
	if err != nil {
		return locTree, err
	}
	locTree.References = halTree.Embedded.RexReferences
	locTree.Tree, err = reconstructLocationGraphfromJSON(halTree.Embedded.RexReferences)
	return locTree, err
}

// GetTransformations fetches all transformations for all references and project files
func (t *LocationGraph) GetTransformations(handler gorexos.RequestHandler) error {
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
func (t *LocationGraph) Beautify() {

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

// WriteToDot gets the location graph of the project and dumps out the structure as DOT file (graphviz)
func (t *LocationGraph) WriteToDot(w io.Writer) error {
	return tree.WriteToDot(t.Tree, w)
}

func reconstructLocationGraphfromJSON(refs []Reference) (*tree.Node, error) {

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
