package gorexos

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gookit/color"
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

// UserRexOSDir is the local directory which should contain the config.yml file
func UserRexOSDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return filepath.Join(home, ".config", "rexos")
	} else if runtime.GOOS == "linux" {
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return filepath.Join(home, "rexos")
		}
	}
	return filepath.Join(os.Getenv("HOME"), ".config", "rexos")
}

// GetInstanceCredentials searches through the list of instances
// and returns the config where the defaultValue matches one of
// the instance with the given username (e.g. rex-production/user1)
// Returns clientID, clientSecret, error
func GetInstanceCredentials(config *Config, instanceName, userName string) (string, string, error) {

	if config == nil {
		return "", "", fmt.Errorf("No config found")
	}
	if instanceName == "" {
		return "", "", fmt.Errorf("Instance name cannot be empty")
	}
	if userName == "" {
		return "", "", fmt.Errorf("Username name cannot be empty")
	}

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

// ReadConfig reads the config from a reader
func ReadConfig(r io.Reader) *Config {

	var config Config

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		color.Red.Println("Cannot read content from config file reader")
		return &config
	}
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		color.Red.Println("Cannot unmarshal config file, file format may be wrong")
	}
	return &config
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
