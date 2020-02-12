package rexos

import (
	b64 "encoding/base64"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

// RequestHandler interface
type RequestHandler interface {
	Authenticate(id, secret string) string
	Post(path string, body interface{}) (*resty.Response, error)
	Delete(path string, body interface{}) (*resty.Response, error)
}

type requestHandler struct {
	client     *resty.Client
	token      string
	userID     string
	domain     string
	pathPrefix string
}

// NewRequestHandler creates a new handler
func NewRequestHandler(domain, pathPrefix string) RequestHandler {

	return &requestHandler{
		client:     resty.New(),
		domain:     domain,
		pathPrefix: pathPrefix,
	}
}

func (r *requestHandler) Authenticate(clientID, clientSecret string) string {

	payload := clientID + ":" + clientSecret
	encodedToken := b64.StdEncoding.EncodeToString([]byte(payload))
	resp, err := r.client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Basic "+encodedToken).
		SetBody([]byte(`grant_type=client_credentials`)).
		Post("https://" + r.domain + "/oauth/token")

	if err != nil {
		panic(err)
	}
	body := string(resp.Body())
	r.token = gjson.Get(body, "access_token").String()
	r.userID = gjson.Get(body, "user_id").String()

	return r.userID
}

func (r *requestHandler) Post(path string, body interface{}) (*resty.Response, error) {

	return r.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetAuthToken(r.token).
		SetBody(body).
		Post("https://" + r.domain + r.pathPrefix + path)
}

func (r *requestHandler) Delete(path string, body interface{}) (*resty.Response, error) {

	return r.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetAuthToken(r.token).
		SetBody(body).
		Delete("https://" + r.domain + r.pathPrefix + path)
}
