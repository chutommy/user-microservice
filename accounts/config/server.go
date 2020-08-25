package config

import (
	"time"
)

// ServerConfig holds the server's attributs.
type ServerConfig struct {
	Address           string        `yaml:"address"`
	ReadTimeout       time.Duration `yaml:"read_timeout"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout"`
	WriteTimeout      time.Duration `yaml:"write_timeout"`
	IdleTimeout       time.Duration `yaml:"idle_timeout"`
}
