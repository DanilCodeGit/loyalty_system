package conf

import "fmt"

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

var dsn = fmt.Sprintf("postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full")

type ServerConfig struct {
	Host string
	Post string
}
