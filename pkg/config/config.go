package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config is the general config file
type Config struct {
	APIVersion string     `yaml:"api-version"`
	Instances  []Instance `yaml:"instances"`
}

// Instance defines one environment
type Instance struct {
	Name  string `yaml:"name"`
	URL   string `yaml:"url"`
	Users []User `yaml:"users"`
}

// User defines a user login
type User struct {
	Name         string `yaml:"name"`
	ClientID     string `yaml:"client-id"`
	ClientSecret string `yaml:"client-secret"`
}

// GetInstanceCredentials searches through the list of instances
// and returns the config where the defaultValue matches one of
// the instance with the given username (e.g. rex-production/user1)
// Returns clientID, clientSecret, error
func GetInstanceCredentials(config *Config, instanceName, userName string) (string, string, error) {

	for _, i := range config.Instances {

		if i.Name == instanceName {
			for _, u := range i.Users {
				if u.Name == userName {
					return u.ClientID, u.ClientSecret, nil
				}
			}
			return "", "", fmt.Errorf("Cannot find username")
		}
	}
	return "", "", fmt.Errorf("Cannot find instance")
}

// ReadConfig reads the config from the file
func ReadConfig(fileName string) *Config {
	r, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	buf, err := ioutil.ReadAll(r)

	c := Config{}

	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		panic(err)
	}
	return &c
}

func (c Config) String() string {
	s := "REXos Configuration (API version " + c.APIVersion + ")\n"
	s += "====================================\n"

	for _, i := range c.Instances {
		s += "Name: " + i.Name + "\n"
		s += "URL:  " + i.URL + "\n"
		s += "Users:  " + "\n"

		for _, u := range i.Users {
			s += "\tName:     " + u.Name + "\n"
			s += "\tClientID: " + u.ClientID + "\n"
		}
		s += "\n"
	}
	s += "\n"
	return s
}
