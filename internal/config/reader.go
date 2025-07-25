package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Read(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	if err := yaml.Unmarshal([]byte(file), c); err != nil {
		return nil, err
	}

	return c, nil
}
