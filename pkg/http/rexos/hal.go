package rexos

import (
	"strings"

	"github.com/tidwall/gjson"
)

// StripTemplateParameter removes the trailing template parameters of an HATEOAS URL
// For example: "https://rex.robotic-eyes.com/rex-gateway/api/v2/rexReferences/1000/project{?projection}"
func StripTemplateParameter(templateURL string) string {
	return strings.Split(templateURL, "{")[0]
}

// GetSelfLinkFromHal returns the stripped self link of a HAL resource. The input is the JSON
// response as string
func GetSelfLinkFromHal(json []byte) string {
	return StripTemplateParameter(gjson.Get(string(json), "_links.self.href").String())
}
