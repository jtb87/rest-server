package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	c := Config{}
	err := c.Initialize("example-config.json")
	if err != nil {
		t.Errorf("got error on init: %s", err)
	}
	if c.Password != "password" {
		t.Errorf("password expected: %s got %s", "password", c.Password)
	}
	if c.Username != "user" {
		t.Errorf("username expected: %s got %s", "user", c.Password)
	}
}
