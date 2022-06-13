package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

const (
	defaultConfigFile = "config.yaml"
)

type ServiceConfig struct {
	Client string `yaml:"rss_feed_url"`
}

func Load() (ServiceConfig, error) {
	var cfg ServiceConfig
	yamlFile, err := ioutil.ReadFile(defaultConfigFile)
	if err != nil {
		return ServiceConfig{}, fmt.Errorf("unable to load config file: %w", err)
	}
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return ServiceConfig{}, fmt.Errorf("unable to parse config file: %w", err)
	}

	return cfg, nil
}
