package rexos

import (
	b64 "encoding/base64"
	"fmt"
	"io"

	"github.com/go-resty/resty/v2"
)

const (
	pathPrefix = "/rex-gateway"
)

// RequestHandler interface
type RequestHandler interface {
	Authenticate(domain, id, secret string) Session
	AuthenticateWithSession(session Session) error
	Get(path string) (*resty.Response, error)
	GetFullyQualified(path string) (*resty.Response, error)
	Post(path string, body interface{}) (*resty.Response, error)
	Delete(path string, body interface{}) (*resty.Response, error)
	PostMultipartFile(path string, fileName string, fileData io.Reader) (*resty.Response, error)
}

type requestHandler struct {
	client  *resty.Client
	session Session
}

// NewRequestHandler creates a new handler
func NewRequestHandler() RequestHandler {

	return &requestHandler{
		client: resty.New(),
	}
}

// AuthenticateWithSession uses an existing session
func (r *requestHandler) AuthenticateWithSession(session Session) error {
	if session.Valid() {
		r.session = session
		return nil
	}
	return fmt.Errorf("Session has expired, please login again")
}

// Authenticate uses clientID and clientSecret for authentication and returns a session
func (r *requestHandler) Authenticate(domain, clientID, clientSecret string) Session {

	payload := clientID + ":" + clientSecret
	encodedToken := b64.StdEncoding.EncodeToString([]byte(payload))
	resp, err := r.client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Basic "+encodedToken).
		SetBody([]byte(`grant_type=client_credentials`)).
		Post("https://" + domain + "/oauth/token")

	if err != nil {
		panic(err)
	}
	r.session = NewSession(domain, resp.Body())
	return r.session
}

func (r *requestHandler) Get(path string) (*resty.Response, error) {

	return r.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetAuthToken(r.session.AccessToken).
		Get("https://" + r.session.Domain + pathPrefix + path)
}

func (r *requestHandler) GetFullyQualified(path string) (*resty.Response, error) {

	return r.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetAuthToken(r.session.AccessToken).
		Get(path)
}

func (r *requestHandler) Post(path string, body interface{}) (*resty.Response, error) {

	return r.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetAuthToken(r.session.AccessToken).
		SetBody(body).
		Post("https://" + r.session.Domain + pathPrefix + path)
}

func (r *requestHandler) Delete(path string, body interface{}) (*resty.Response, error) {

	return r.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetAuthToken(r.session.AccessToken).
		SetBody(body).
		Delete("https://" + r.session.Domain + pathPrefix + path)
}

func (r *requestHandler) PostMultipartFile(path string, fileName string, fileData io.Reader) (*resty.Response, error) {

	return r.client.R().
		SetFileReader("file", fileName, fileData).
		SetAuthToken(r.session.AccessToken).
		Post("https://" + r.session.Domain + pathPrefix + path)
}
