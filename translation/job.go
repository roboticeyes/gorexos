package translation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/roboticeyes/gorexos"
)

// Job contains the specification of a translation job
type Job struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type StringList struct {
	Names []string `json:"names"`
}

// CreateJob creates a new translation job
func CreateJob(handler gorexos.RequestHandler, job Job) (Job, error) {

	resp, err := handler.Post("/translation/v1/jobs", job)
	if err != nil {
		return job, err
	}

	err = json.Unmarshal(resp.Body(), &job)

	if resp.StatusCode() != http.StatusCreated {
		return job, fmt.Errorf("Failed, got back %d: %s", resp.StatusCode(), resp.Body())
	}
	return job, nil
}

func UploadFile(handler gorexos.RequestHandler, jobId string, fileName string, r io.Reader) error {

	url := "/translation/v1/jobs/" + jobId + "/files"
	resp, err := handler.PostMultipartFile(url, filepath.Base(fileName), r)
	if resp.StatusCode() != http.StatusCreated {
		return fmt.Errorf("Failed, got back %d: %s", resp.StatusCode(), resp.Body())
	}
	return err
}

// GetPipelines get all the available pipelines
func GetPipelines(handler gorexos.RequestHandler) ([]string, error) {

	var res StringList
	url := "/translation/v1/pipelines"
	resp, err := handler.Get(url)
	if resp.StatusCode() != http.StatusOK {
		return []string{}, fmt.Errorf("Failed, got back %d: %s", resp.StatusCode(), resp.Body())
	}

	var names []string
	err = json.Unmarshal(resp.Body(), &res)
	for _, v := range res.Names {
		names = append(names, v)
	}
	return names, err
}

// GetSupportedFormats get all the supported file formats
func GetSupportedFormats(handler gorexos.RequestHandler) ([]string, error) {

	var res StringList
	url := "/translation/v1/formats"
	resp, err := handler.Get(url)
	if resp.StatusCode() != http.StatusOK {
		return []string{}, fmt.Errorf("Failed, got back %d: %s", resp.StatusCode(), resp.Body())
	}

	var names []string
	err = json.Unmarshal(resp.Body(), &res)
	for _, v := range res.Names {
		names = append(names, v)
	}
	return names, err
}

// GetJob gets the current status of the job
func GetJob(handler gorexos.RequestHandler, jobId string) (Job, error) {

	var job Job
	url := "/translation/v1/jobs/" + jobId
	resp, err := handler.Get(url)
	if resp.StatusCode() != http.StatusOK {
		return job, fmt.Errorf("Failed, got back %d: %s", resp.StatusCode(), resp.Body())
	}

	err = json.Unmarshal(resp.Body(), &job)
	return job, err
}

// GetJobResult gets the data from the job
func GetJobResult(handler gorexos.RequestHandler, jobId string, w io.Writer) error {

	url := "/translation/v1/results/" + jobId
	resp, err := handler.Get(url)
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("Failed, got back %d: %s", resp.StatusCode(), resp.Body())
	}

	io.Copy(w, bytes.NewReader(resp.Body()))
	return err
}

// StartJob starts a job
func StartJob(handler gorexos.RequestHandler, jobId string, pipeline string) error {

	url := "/translation/v1/jobs/" + jobId + "/start"
	payload := struct {
		Pipeline string `json:"pipeline"`
	}{
		Pipeline: pipeline,
	}
	resp, err := handler.Post(url, payload)
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("Failed, got back %d: %s", resp.StatusCode(), resp.Body())
	}
	return err
}
