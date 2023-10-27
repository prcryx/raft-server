package config

import (
	"log"
	"sync"

	godotenv "github.com/joho/godotenv"
	"github.com/prcryx/raft-server/internal/common/constants"
)

type EnvConfig struct {
	Host             string
	Port             string
	DbUrl            string
	TwilioAccountSid string
	TwilioAuthToken  string
	TwilioVerifySid  string
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
			Port:             envMap[constants.Port],
			Host:             envMap[constants.Host],
			DbUrl:            envMap[constants.DbUrl],
			TwilioAccountSid: envMap[constants.TwilioAccountSid],
			TwilioAuthToken:  envMap[constants.TwilioAuthToken],
			TwilioVerifySid:  envMap[constants.TwilioVerifySid],
		}
	})
	return env, nil
}

// var ConfigSet =	wire.NewSet(
// 	LoadConfig,
// )
