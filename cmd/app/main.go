package main

import (
	"log"
	// "net/http"
	"os"

	config "github.com/prcryx/raft-server/config"
	"github.com/prcryx/raft-server/di/wire"
	"github.com/prcryx/raft-server/internal/application/server"
	"github.com/prcryx/raft-server/internal/common/constants/routesconst"
	// di "github.com/prcryx/raft-server/cmd/di"
)

func main() {
	var exitCode int = 0

	defer func() {
		if exitCode != 0 {
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

	//init firebase App
	firebaseApp, appInitializationError := wire.InitFirebaseApp(config)
	if appInitializationError != nil {
		log.Printf("Error: %s", appInitializationError.Error())
		exitCode = 1
		return
	}

	// init authClient
	authClient, authClientInitializationError := wire.InitFirebaseAuthClient(firebaseApp)
	if authClientInitializationError != nil {
		log.Printf("Error: %s", authClientInitializationError.Error())
		exitCode = 1
		return
	}

	// init controllerRegistry
	controllerRegistry, controllerRegistryIntializationError := wire.InitializeControllerRegistry(authClient)
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
		exitCode = 1
		return
	}

}
