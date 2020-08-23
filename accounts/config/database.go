package config

// DBConfig holds the credentials for the database connection.
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
}
type dbConfig struct {
	host     string `yaml:"host"`
	port     int    `yaml:"port"`
	user     string `yaml:"user"`
	password string `yaml:"password"`
	dBname   string `yaml:"d_bname"`
}
