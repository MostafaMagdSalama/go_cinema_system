package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type EnvireonmentConfig struct {
	DB_CONNECTION string `envvonfig:"DB_CONNECTION" required:"true"`
}

var (
	loadENV    = godotenv.Load
	processENV = envconfig.Process
)

func ReadEnvironmentVars(envFilePath string) (*EnvireonmentConfig, error) {

	if err := loadENV(envFilePath); err != nil {
		return nil, errors.Wrap(err, "loading env vars")
	}

	envConfig := new(EnvireonmentConfig)

	if err := processENV("", envConfig); err != nil {
		return nil, errors.Wrap(err, "processing env vars")
	}
	return envConfig, nil

}
