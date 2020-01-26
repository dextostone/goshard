package config

import "testing"

func TestValidateConfig(t *testing.T) {
	nilConfig := Config{}
	err := validateConfig(nilConfig)
	if err == nil {
		t.Error("validateConfig should have returned error")
	}

}
