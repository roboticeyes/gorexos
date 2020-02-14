package rexos

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"time"

	"github.com/gookit/color"
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
	Expires         time.Time `json:"expires"`
}

// NewSession creates a new session of the authentication response
func NewSession(domain string, body []byte) Session {

	var session Session
	err := json.Unmarshal(body, &session)
	if err != nil {
		panic(err)
	}

	session.Expires = time.Now().Add(time.Second * time.Duration(session.ExpiresIn))
	session.Domain = domain
	return session
}

// Valid checks if the session has already expired
func (s *Session) Valid() bool {

	if time.Now().Sub(s.Expires) > 0 {
		return false
	}
	return true
}

// OpenSession opens a session from a given file reader
func OpenSession(r io.Reader) (Session, error) {

	var session Session
	body, err := ioutil.ReadAll(r)

	if err != nil {
		return session, err
	}

	err = json.Unmarshal(body, &session)

	color.Cyan.Println("============================================================")
	color.Cyan.Println("Instance:", session.Domain)
	color.Cyan.Println("Username:", session.UserName)
	color.Cyan.Println("============================================================")
	return session, err
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
