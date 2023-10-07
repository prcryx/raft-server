package config

import (
	"log"

	godotenv "github.com/joho/godotenv"
	"github.com/prcryx/raft-server/internal/common/constants"
)

type EnvConfig struct {
	Host                  string
	Port                  string
	ServiceAccountKeyFile string
}

func LoadConfig() (*EnvConfig, error) {
	var env *EnvConfig
	godotenv.Load(".env")

	envMap, err := godotenv.Read()

	if err != nil {
		log.Fatal(1)
	}

	env = &EnvConfig{
		Port:                  envMap[constants.Port],
		Host:                  envMap[constants.Host],
		ServiceAccountKeyFile: envMap[constants.ServiceAccountKeyFilePath],
	}

	return env, nil

}
