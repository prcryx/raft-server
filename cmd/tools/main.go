package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

// configuration data

type Config struct {
	clientHost        string `env:"CLIENT_HOST"`
	serverPort        string `env:"SERVER_PORT"`
	serviceAccountKey string `env:"SERVICE_ACCOUNT_KEY_FILE_PATH"`
	osName            string `env:"OS_NAME"`
	archName          string `env:"ARCH_NAME"`
}

// GenerateEnvFile generates a .env file based on the values in the EnvConfig struct.
func GenerateEnvFile(config Config, filePath string) error {
	// Open the .env file for writing
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get the type of the struct
	t := reflect.TypeOf(config)

	// Iterate through the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Tag.Get("env")
		fieldValue := reflect.ValueOf(config).Field(i).Interface()

		// Write key-value pairs to the .env file
		_, err := fmt.Fprintf(file, "%s=%v\n", fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	var config Config
	scanner := bufio.NewScanner(os.Stdin)

	//client host
	fmt.Print("CLIENT_HOST: ")
	scanner.Scan()
	config.clientHost = scanner.Text()

	//server port
	fmt.Print("SERVER_PORT: ")
	scanner.Scan()
	config.serverPort = scanner.Text()

	//serviceAccountKeyFilePath
	fmt.Print("SERVICE_ACCOUNT_KEY_FILE_PATH: ")
	scanner.Scan()
	config.serviceAccountKey = scanner.Text()

	//OS name
	fmt.Print("OS_NAME: ")
	scanner.Scan()
	config.osName = scanner.Text()

	//ARCH name
	fmt.Print("ARCH_NAME: ")
	scanner.Scan()
	config.archName = scanner.Text()

	//generate files
	if err := GenerateEnvFile(config, ".env"); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(".env file generated successfully!")
}
