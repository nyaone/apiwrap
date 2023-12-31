package inits

import (
	"apiwrap/config"
	"gopkg.in/yaml.v3"
	"os"
)

func Config() error {
	// Read config file
	configFileBytes, err := os.ReadFile("config.yml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configFileBytes, &config.Config)
	if err != nil {
		return err
	}

	return nil
}
