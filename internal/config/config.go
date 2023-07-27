package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	DBuser  string `yaml:"DBuser"`
	DBpass  string `yaml:"DBpass"`
	DBnet   string `yaml:"DBnet"`
	DBaddr  string `yaml:"DBaddr"`
	DBport  string `yaml:"DBport"`
	DBname  string `yaml:"DBname"`
	WebAddr string `yaml:"WebAddr"`
	WebPort string `yaml:"WebPort"`
}

func LoadConfigFromYAML(f string) (*ServerConfig, error) {
	file, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	var data ServerConfig

	err = yaml.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
