package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config Struct will hold all the configurations for the App
type Config struct {
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Initialize will load a config file from the specified path and will set all the appropiate values
func (c *Config) Initialize(configFile string) error {
	f, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	json.Unmarshal(f, &c)
	return nil
}
