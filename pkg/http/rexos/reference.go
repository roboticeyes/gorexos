package rexos

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

// Reference is representing a simple form of a REX reference potentially including all its children
type Reference struct {
	Name     string `json:"name"`
	Urn      string `json:"urn"`
	Key      string `json:"key"`
	Type     string `json:"type"`
	Children []Reference
}

// GetReferenceByKey returns a REX reference by the given key
func GetReferenceByKey(handler RequestHandler, key string) (Reference, error) {

	var ref Reference

	resp, err := handler.Get(apiReferenceByKey + "key=" + key)

	if err != nil {
		return ref, err
	}
	err = json.Unmarshal(resp.Body(), &ref)

	if ref.Key != key {
		return ref, fmt.Errorf("%s", gjson.Get(string(resp.Body()), "message").String())
	}

	// Get child references
	var childReferences []Reference
	childReferencesRaw := []byte(gjson.GetBytes(resp.Body(), "_embedded.childReferences").Raw)
	err = json.Unmarshal(childReferencesRaw, &childReferences)

	for _, r := range childReferences {
		ref.Children = append(ref.Children, Reference{
			Name: r.Name,
			Urn:  r.Urn,
			Key:  r.Key,
			Type: r.Type,
		})
	}
	return ref, err
}

// GetReferenceTreeByProjectUrn returns the full reference tree for the given project URN
func GetReferenceTreeByProjectUrn(handler RequestHandler, projectUrn string) (Reference, error) {

	var root Reference

	// Get project
	project, err := GetProjectByUrn(handler, projectUrn)
	if err != nil {
		return root, err
	}

	// Get root reference
	root, err = GetReferenceByKey(handler, project.RootRexReferenceKey)
	if err != nil {
		return root, err
	}

	// Iterate over all children
	for _, r := range root.Children {
		ref, err := GetReferenceByKey(handler, r.Key)
		if err != nil {
			panic(err)
		}
		root.Children = append(root.Children, ref)
	}

	return root, err
}

func indent(level int) string {
	str := ""
	for i := 0; i < level; i++ {
		str += "  "
	}
	return str
}

func printReference(level int, r Reference) string {

	var str string
	str += "\n"
	str += fmt.Sprintf("%sType: %s\n", indent(level), r.Type)
	str += fmt.Sprintf("%sName: %s\n", indent(level), r.Name)
	str += fmt.Sprintf("%sUrn:  %s\n", indent(level), r.Urn)
	str += fmt.Sprintf("%sKey:  %s\n", indent(level), r.Key)

	for i := 0; i < len(r.Children); i++ {
		str += printReference(i+1, r.Children[i])
	}
	return str
}

func (r Reference) String() string {

	return printReference(0, r)
}
