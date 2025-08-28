package config

import (
	"log"
	"os"

	"github.com/beeploop/sylvie/internal/utils"
	"gopkg.in/yaml.v3"
)

const (
	DEFAULT_CONFIG_FILE = "sylvieconfig.yaml"
)

type Config struct {
	OutDir string `yaml:"out_dir"`
}

func Init(configFile *string) *Config {
	f, err := os.ReadFile(utils.Ternary(configFile == nil, *configFile, DEFAULT_CONFIG_FILE))
	if err != nil {
		log.Fatal("configuration file not found")
	}

	var config Config
	if err := yaml.Unmarshal(f, &config); err != nil {
		log.Fatal(err.Error())
	}

	if err := os.MkdirAll(config.OutDir, 0777); err != nil {
		log.Fatal(err.Error())
	}

	return &config
}
