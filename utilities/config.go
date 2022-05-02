package utilities

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Smtp struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	To       string `yaml:"to"`
	From     string `yaml:"from"`
}

type Config struct {
	SnapraidBin string `yaml:"snapraid_bin"`
	Smtp        Smtp
}

func LoadConfigFile() (config *Config, err error) {
	// Load config.yml
	yfile, err := ioutil.ReadFile("config.yaml")

	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(yfile, &config); err != nil {
		return nil, err
	}

	return config, nil
}
