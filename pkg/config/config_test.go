package config

import (
	"bytes"
	"testing"
)

func TestReadConfig(t *testing.T) {

	r := bytes.NewReader([]byte(sampleConfig))

	c := ReadConfig(r)

	if c == nil {
		t.Fatal("ReadConfig returned nil")
	}

	if c.APIVersion != "v2" {
		t.Fatal("APIversion does not match")
	}

	if len(c.Instances) < 1 {
		t.Fatal("No instances found")
	}

	id, secret, err := GetInstanceCredentials(c, "production", "testuser")

	if err != nil {
		t.Fatal("Cannot get instance credentials: ", err)
	}

	if id != "12345" {
		t.Fatalf("ID does not match: expected %s, got %s", "12345", id)
	}
	if secret != "67890" {
		t.Fatalf("ID does not match: expected %s, got %s", "67890", secret)
	}
}

const sampleConfig = `
api-version: v2

instances:
    - name: production
      url: api.rexos.cloud
      users:
        - name: testuser
          client-id: 12345
          client-secret: 67890
`
