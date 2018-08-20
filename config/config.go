package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config Struct will hold all the configurations for the App
type Config struct {
	Port       string `json:"port"`
	DBUsername string `json:"username"`
	DBPassword string `json:"password"`
}

// Initialize will load a config file from the specified path and will set all the appropiate values
func (c *Config) Initialize(configFile string) (Config, error) {
	f, err := os.Open(configFile)
	if err != nil {
		return err
	}
	defer f.Close()
	jP := json.NewDecoder(f)
	err = jP.Decode(&c)
	if err != nil {
		return err
	}
	return nil
}
