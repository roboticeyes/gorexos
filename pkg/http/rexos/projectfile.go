package rexos

// ProjectFile structure of REXos
type ProjectFile struct {
	LastModified string `json:"lastModified"`
	ContentType  string `json:"contentType"`
	Urn          string `json:"urn"`
	FileSize     int    `json:"fileSize"`
	ContentHash  string `json:"contentHash"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Links        struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		FileDownload struct {
			Href string `json:"href"`
		} `json:"file.download"`
	} `json:"_links"`
}

type projectFilesHal struct {
	Embedded struct {
		ProjectFiles []ProjectFile `json:"projectFiles"`
	} `json:"_embedded"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}
