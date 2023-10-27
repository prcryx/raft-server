package main

import (
	"log"

	// "net/http"
	"os"

	config "github.com/prcryx/raft-server/config"
	"github.com/prcryx/raft-server/di/wire"
	"github.com/prcryx/raft-server/internal/application/server"
	"github.com/prcryx/raft-server/internal/common/constants/routesconst"
)

func main() {
	var exitCode int = 0

	defer func() {
		if exitCode != 0 {
			//other clean up code
			os.Exit(exitCode)
		}
	}()

	//load the config
	config, configError := config.LoadConfig()
	if configError != nil {
		log.Printf("Error: %s", configError.Error())
		exitCode = 1
		return
	}

	//init twilio app
	twilioApp, twilioInitializeErr := wire.InitTwilioApp(config)
	if twilioInitializeErr != nil {
		log.Printf("Error: %s", twilioInitializeErr.Error())
		exitCode = 1
		return
	}

	//init database
	db, dbInitializeError := wire.InitDatabase(config)
	if dbInitializeError != nil {
		log.Printf("Error: %s", dbInitializeError.Error())
		exitCode = 1
		return
	}
	// init controllerRegistry
	controllerRegistry, controllerRegistryIntializationError := wire.InitializeControllerRegistry(db, twilioApp)
	if controllerRegistryIntializationError != nil {
		log.Printf("Error: %s", controllerRegistryIntializationError.Error())
		exitCode = 1
		return
	}

	// init Server
	srv, serverInitializationError := wire.InitServer(controllerRegistry, config, routesconst.V1)
	if serverInitializationError != nil {
		log.Printf("Error: %s", serverInitializationError.Error())
		exitCode = 1
		return
	}

	//start the server
	srvStartErr := server.StartServer(srv)
	if srvStartErr != nil {
		log.Printf("Error: %s", srvStartErr.Error())
		exitCode = 1
		return
	}

}
