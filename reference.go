package gorexos

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

// Reference is representing a simple form of a REX reference potentially including all its children
type Reference struct {
	Urn                 string         `json:"urn"`
	ParentReferenceUrn  string         `json:"parentReferenceUrn"`
	RootReference       bool           `json:"rootReference"`
	Name                string         `json:"name"`
	Key                 string         `json:"key"`
	Category            string         `json:"category"`
	ProjectFileUrn      string         `json:"projectFileUrn"`
	WorldTransformation Transformation `json:"worldTransformation"`
	LocalTransformation Transformation `json:"localTransformation"`
	Type                string         `json:"type"`
	Links               struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Project struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"project"`
		ProjectFile struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"projectFile"`
		ParentReference struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"parentReference"`
	} `json:"_links"`
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

	return ref, err
}

func (r Reference) String() string {

	var str string
	str += fmt.Sprintf("%s\n", r.Urn)
	return str
}
