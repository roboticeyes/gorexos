package translation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/roboticeyes/gorexos/pkg/http/rexos"
)

// Job contains the specification of a translation job
type Job struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// CreateJob creates a new translation job
func CreateJob(handler rexos.RequestHandler, job Job) (Job, error) {

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

func UploadFile(handler rexos.RequestHandler, jobId string, fileName string, r io.Reader) error {

	url := "/translation/v1/jobs/" + jobId + "/files"
	resp, err := handler.PostMultipartFile(url, filepath.Base(fileName), r)
	if resp.StatusCode() != http.StatusCreated {
		return fmt.Errorf("Failed, got back %d: %s", resp.StatusCode(), resp.Body())
	}
	return err
}

// GetJob gets the current status of the job
func GetJob(handler rexos.RequestHandler, jobId string) (Job, error) {

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
func GetJobResult(handler rexos.RequestHandler, jobId string, w io.Writer) error {

	url := "/translation/v1/results/" + jobId
	resp, err := handler.Get(url)
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("Failed, got back %d: %s", resp.StatusCode(), resp.Body())
	}

	io.Copy(w, bytes.NewReader(resp.Body()))
	return err
}

// StartJob starts a job
func StartJob(handler rexos.RequestHandler, jobId string, pipeline string) error {

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
