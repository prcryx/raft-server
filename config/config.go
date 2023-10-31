package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type jwtConfig struct {
	AccessTokenSecret  string `yaml:"accessTokenSecreet"`
	RefreshTokenSecret string `yaml:"refreshTokenSecret"`
}

type serverConfig struct {
	Host string    `yaml:"host"`
	Port int       `yaml:"port"`
	JWT  jwtConfig `yaml:"jwt"`
}

type dbConfig struct {
	URL string `yaml:"url"`
}

type twilioConfig struct {
	AccountSid string `yaml:"accountSid"`
	AuthToken  string `yaml:"authToken"`
	VerifySid  string `yaml:"verifySid"`
}

type EnvConfig struct {
	Server serverConfig `yaml:"server"`
	DB     dbConfig     `yaml:"db"`
	Twilio twilioConfig `yaml:"twilio"`
}

// var once sync.Once

func LoadConfig() (*EnvConfig, error) {
	conf := new(EnvConfig)

	file, err := os.ReadFile(".env.yaml")
	if err != nil {
		return nil, err
	}

	decodeErr := yaml.Unmarshal(file, conf)
	if decodeErr != nil {
		return nil, decodeErr
	}

	// envMap, err := godotenv.Read()
	// if err != nil {
	// 	log.Fatal(1)
	// }
	// env = &EnvConfig{
	// 	Port:             envMap[constants.Port],
	// 	Host:             envMap[constants.Host],
	// 	DbUrl:            envMap[constants.DbUrl],
	// 	TwilioAccountSid: envMap[constants.TwilioAccountSid],
	// 	TwilioAuthToken:  envMap[constants.TwilioAuthToken],
	// 	TwilioVerifySid:  envMap[constants.TwilioVerifySid],
	// 	SecretKey:        envMap[constants.TwilioVerifySid],
	// }

	return conf, nil
}
