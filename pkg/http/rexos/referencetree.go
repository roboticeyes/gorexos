package rexos

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tidwall/gjson"
)

// ReferenceTree represents the complete reference tree of one project
type ReferenceTree struct {
	Head *Reference
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

	var tree ReferenceTree

	project, err := GetProjectByUrn(handler, projectUrn)
	if err != nil {
		return tree, err
	}

	resp, err := handler.GetFullyQualified(project.SelfLink + "/rexReferences?projection=linkedList")
	if err != nil {
		return tree, err
	}

	var halTree referenceTreeHal
	err = json.Unmarshal(resp.Body(), &halTree)

	f, _ := os.Create("/tmp/reftree.dot")
	defer f.Close()

	fmt.Fprintf(f, "digraph {")
	for _, v := range halTree.Embedded.RexReferences {
		if len(v.ParentReferenceUrn) > 0 {
			fmt.Fprintf(f, "%s -> %s;\n",
				v.ParentReferenceUrn[len(v.ParentReferenceUrn)-4:],
				v.Urn[len(v.Urn)-4:])
		}

		// check if projectfile exists
		pflink := StripTemplateParameter(v.Links.ProjectFile.Href)
		pfresp, err := handler.GetFullyQualified(pflink)
		if err != nil {
			fmt.Println("Error during getting project file")
		} else {
			if pfresp.StatusCode() < 400 {
				name := gjson.GetBytes(pfresp.Body(), "name").String()
				fmt.Fprintf(f, "%s -> \"%s\";\n", v.Urn[len(v.Urn)-4:], name[0:8])
				fmt.Fprintf(f, "\"%s\" [shape=box,style=filled,color=yellow]\n", name[0:8])
			}
		}
	}

	for _, v := range halTree.Embedded.RexReferences {
		if v.Type == "portal" {
			fmt.Fprintf(f, "%s [shape=doublecircle,style=filled,color=lightblue]\n", v.Urn[len(v.Urn)-4:])
		} else if v.Type == "group" {
			fmt.Fprintf(f, "%s [style=filled,color=cyan]\n", v.Urn[len(v.Urn)-4:])
		} else if v.Type == "file" {
			fmt.Fprintf(f, "%s [style=filled,color=green]\n", v.Urn[len(v.Urn)-4:])
		}
	}
	fmt.Fprintln(f, "}")

	return tree, err
}
