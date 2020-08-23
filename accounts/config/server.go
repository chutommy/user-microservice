package config

import (
	"time"
)

// ServerConfig holds the server's attributs.
type ServerConfig struct {
	Address           string
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
}
type serverConfig struct {
	address           string `yaml:"address"`
	readTimeout       string `yaml:"read_timeout"`
	readHeaderTimeout string `yaml:"read_header_timeout"`
	writeTimeout      string `yaml:"write_timeout"`
	idleTimeout       string `yaml:"idle_timeout"`
	maxHeaderBytes    int    `yaml:"max_header_bytes"`
}
