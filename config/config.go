package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v3"
)

// Config struct holds service configuration.
type Config struct {
	Port          int    `yaml:"port"`
	Logging       bool   `yaml:"logging"`
	JSONParamPath string `yaml:"json_param_path"`
	URLParamPath  string `yaml:"url_param_path"`
}

// NewDefault return default config.
func NewDefault() Config {
	return Config{
		Port:          8090,
		Logging:       true,
		JSONParamPath: "/json",
		URLParamPath:  "/url",
	}
}

// NewFromFile read configuration from YML file, return default on failure.
func NewFromFile(path string) (Config, error) {
	conf := NewDefault()

	raw, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return conf, fmt.Errorf("failed to read config file %s, %w", path, err)
	}

	if err := yaml.Unmarshal(raw, &conf); err != nil {
		return conf, fmt.Errorf("failed to read config as YML: %w", err)
	}

	return conf, nil
}
