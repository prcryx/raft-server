package config

import (
	"log"
	"sync"

	godotenv "github.com/joho/godotenv"
	"github.com/prcryx/raft-server/internal/common/constants"
)

type EnvConfig struct {
	Host                  string
	Port                  string
	ServiceAccountKeyFile string
}

var once sync.Once

func LoadConfig() (*EnvConfig, error) {
	var env *EnvConfig
	once.Do(func() {
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
	})
	return env, nil
}

// var ConfigSet =	wire.NewSet(
// 	LoadConfig,
// )
