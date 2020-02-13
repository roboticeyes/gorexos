package rexos

import (
	"bytes"
	"encoding/json"
	"io"
	"time"
)

// Session stores the information which gets retrieved after successful authentication. It also
// contains the token.
type Session struct {
	Domain          string    `json:"domain"`
	AccessToken     string    `json:"access_token"`
	TokenType       string    `json:"token_type"`
	ExpiresIn       int       `json:"expires_in"`
	Scope           string    `json:"scope"`
	UserID          string    `json:"user_id"`
	UserName        string    `json:"user_name"`
	UserDisplayName string    `json:"user_display_name"`
	Jti             string    `json:"jti"`
	Expires         time.Time `json:expires`
}

// NewSession creates a new session of the authentication response
func NewSession(domain string, body []byte) Session {

	var session Session
	err := json.Unmarshal(body, &session)
	if err != nil {
		panic(err)
	}

	// TODO set the expires date using the expires_in value in seconds
	session.Expires = time.Now().Add(time.Second * time.Duration(session.ExpiresIn))
	session.Domain = domain
	return session
}

// Valid checks if the session has already expired
func (s *Session) Valid() bool {
	// TODO check if token is still valid
	return true
}

func (s *Session) Write(w io.Writer) error {

	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		_, err := w.Write(data)
		return err
	}
	_, err = w.Write(out.Bytes())
	return err

}
