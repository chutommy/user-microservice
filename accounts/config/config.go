package config

// Config defines the configuration for the microservice.
type Config struct {
	Server *ServerConfig
	Db     *DBConfig
}
