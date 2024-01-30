package conf

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

const DSN = "postgresql://user:user@localhost:5432/postgres"

type Config struct {
	Server struct {
		Host string `yaml:"host" env-description:"server host" env-default:"localhost"`
		Port string `yaml:"port" env-description:"server port" env-default:"8080"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host" env-description:"db host" env-default:"localhost"`
		Port     uint16 `yaml:"port" env-description:"db port" env-default:"5432"`
		Database string `yaml:"database" env-description:"db name" env-default:"postgres"`
		User     string `yaml:"user" env-description:"db user" env-default:"postgres"`
		Password string `yaml:"password" env-description:"db password" env-default:"postgres"`
	} `yaml:"database"`
}

var cfg Config

func InitConfig() error {
	err := cleanenv.ReadConfig("./internal/conf/conf.yml", &cfg)
	if err != nil {
		return errors.WithMessage(err, "read config file")
	}

	return nil
}
