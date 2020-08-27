package config

import (
	"io/ioutil"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Config defines the configuration for the microservice.
type Config struct {
	Server *ServerConfig `yaml:"server"`
	Db     *DBConfig     `yaml:"db"`
}

// GetConfig retrieves the data from the config.yml file.
func GetConfig(filename string) (*Config, error) {

	// get config file's content
	path := filepath.Join(rootDir(), filename)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read file")
	}

	// decode the yaml values
	var cfg Config
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "faile to parse the config file")
	}

	return &cfg, nil
}

// rootDir returns the caller's root directory
func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "..")
}
