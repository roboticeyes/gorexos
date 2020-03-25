package rexos

import (
	"encoding/json"
	"fmt"
	"time"
)

// ReferenceTree represents the complete reference tree of one project
// The nodes are stored in a map structure, and the hierarchy is kept in a simple
// structure with UIDs
type ReferenceTree struct {
	Head       *Node
	References map[string]Reference
}

// Node just consists of the URN and its children
type Node struct {
	Urn      string
	Children []*Node
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

	project, err := GetProjectByUrn(handler, projectUrn)
	if err != nil {
		return ReferenceTree{}, err
	}

	resp, err := handler.GetFullyQualified(project.SelfLink + "/rexReferences?projection=linkedList")
	if err != nil {
		return ReferenceTree{}, err
	}

	var halTree referenceTreeHal
	err = json.Unmarshal(resp.Body(), &halTree)
	if err != nil {
		return ReferenceTree{}, err
	}

	return reconstructReferenceTreefromJSON(halTree.Embedded.RexReferences)
}

func (t *ReferenceTree) findAndAddNode(node *Node, parentUrn, childUrn string) bool {

	fmt.Printf("CURRENT NODE URN %s\n", node.Urn)
	if node.Urn == parentUrn {
		fmt.Printf("- found parent %s for child %s\n", parentUrn, childUrn)
		node.Children = append(node.Children, &Node{Urn: childUrn})
		return true
	}
	for i := 0; i < len(node.Children); i++ {
		fmt.Printf("- looking for a parent of child %s\n", childUrn)
		return t.findAndAddNode(node.Children[i], parentUrn, childUrn)
	}
	fmt.Printf("- not found anything\n")
	return false
}

func unlinkedNodes(visited map[string]bool) int {

	var c int
	for _, v := range visited {
		if v == true {
			c++
		}
	}
	return c
}

func reconstructReferenceTreefromJSON(refs []Reference) (ReferenceTree, error) {

	var tree ReferenceTree
	tree.References = make(map[string]Reference)

	visited := make(map[string]bool)
	for i := 0; i < len(refs); i++ {
		tree.References[refs[i].Urn] = refs[i]
		if refs[i].RootReference {
			tree.Head = &Node{Urn: refs[i].Urn}
			visited[refs[i].Urn] = true
			break
		}
	}
	// no root found
	if len(visited) < 1 {
		return tree, fmt.Errorf("Cannot find root reference")
	}

	for {
		fmt.Println(tree)
		if unlinkedNodes(visited) < 1 {
			break // we are done
		}
		for i := 0; i < len(refs); i++ {
			if visited[refs[i].Urn] {
				continue
			}
			success := tree.findAndAddNode(tree.Head, refs[i].ParentReferenceUrn, refs[i].Urn)
			fmt.Println(refs[i].Urn, success)
			visited[refs[i].Urn] = success
		}
		fmt.Println(unlinkedNodes(visited), len(refs))
		// fmt.Println(nodeMap)
		time.Sleep(time.Second)
	}

	// f, _ := os.Create("/tmp/reftree.dot")
	// defer f.Close()
	//
	// fmt.Fprintf(f, "digraph {")
	// for _, v := range halTree.Embedded.RexReferences {
	// 	fmt.Println("Category:", v.Category)
	// 	if len(v.ParentReferenceUrn) > 0 {
	// 		fmt.Fprintf(f, "%s -> %s;\n",
	// 			v.ParentReferenceUrn[len(v.ParentReferenceUrn)-4:],
	// 			v.Urn[len(v.Urn)-4:])
	// 	}

	// check if projectfile exists
	// pflink := StripTemplateParameter(v.Links.ProjectFile.Href)
	// pfresp, err := handler.GetFullyQualified(pflink)
	// if err != nil {
	// 	fmt.Println("Error during getting project file")
	// } else {
	// 	if pfresp.StatusCode() < 400 {
	// 		name := gjson.GetBytes(pfresp.Body(), "name").String()
	// 		fmt.Fprintf(f, "%s -> \"%s\";\n", v.Urn[len(v.Urn)-4:], name[0:8])
	// 		fmt.Fprintf(f, "\"%s\" [shape=box,style=filled,color=yellow]\n", name[0:8])
	// 	}
	// }
	// }

	// for _, v := range halTree.Embedded.RexReferences {
	// 	if v.Type == "portal" {
	// 		fmt.Fprintf(f, "%s [shape=doublecircle,style=filled,color=lightblue]\n", v.Urn[len(v.Urn)-4:])
	// 	} else if v.Type == "group" {
	// 		fmt.Fprintf(f, "%s [style=filled,color=cyan]\n", v.Urn[len(v.Urn)-4:])
	// 	} else if v.Type == "file" {
	// 		fmt.Fprintf(f, "%s [style=filled,color=green]\n", v.Urn[len(v.Urn)-4:])
	// 	}
	// }
	// fmt.Fprintln(f, "}")

	return tree, nil
}

func indent(level int) string {
	str := ""
	for i := 0; i < level; i++ {
		str += "  "
	}
	return str
}

func printNode(level int, n *Node) string {

	var str string
	str += fmt.Sprintf("%sUrn: %s\n", indent(level), n.Urn)
	for i := 0; i < len(n.Children); i++ {
		str += printNode(i+1, n.Children[i])
	}
	return str
}

func (t ReferenceTree) String() string {

	var str string
	str += fmt.Sprintf("ROOT: %s\n", t.Head.Urn)
	for i := 0; i < len(t.Head.Children); i++ {
		str += fmt.Sprintf("%sUrn: %s\n", indent(1), t.Head.Children[i].Urn)
		for j := 0; j < len(t.Head.Children[i].Children); j++ {
			str += fmt.Sprintf("%sUrn: %s\n", indent(2), t.Head.Children[i].Children[j].Urn)
		}

	}

	// str += "\n"
	// str += fmt.Sprintf("%sType: %s\n", indent(level), r.Type)
	// str += fmt.Sprintf("%sName: %s\n", indent(level), r.Name)
	// str += fmt.Sprintf("%sUrn:  %s\n", indent(level), r.Urn)
	// str += fmt.Sprintf("%sKey:  %s\n", indent(level), r.Key)
	// str += fmt.Sprintf("%sRoot: %v\n", indent(level), r.RootReference)

	// return printNode(0, t.Head)
	return str
}
