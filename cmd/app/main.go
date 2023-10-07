package main

import (
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/prcryx/raft-server/cmd/routes"
	config "github.com/prcryx/raft-server/config"
	e "github.com/prcryx/raft-server/internal/common/err"
	"github.com/prcryx/raft-server/internal/db"
)

type RunAppConfig struct {
	FirebaseInstance *firebase.App
	Env              *config.EnvConfig
}

func main() {
	env, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	//initialize firebase

	firebaseInstance, err := db.InitFirebaseApp(env)
	if err != nil {
		log.Fatal(e.FirebaseLoadError)
	}

	//run the web server
	runAppConfig := &RunAppConfig{
		Env:              env,
		FirebaseInstance: firebaseInstance,
	}
	runApp(runAppConfig)
}

func runApp(runAppConfig *RunAppConfig) {
	rootRoute := routes.Root(runAppConfig.FirebaseInstance)
	routes.MountAll(rootRoute, routes.V1)

	//create server
	srv := http.Server{
		Handler: rootRoute,
		Addr:    ":" + runAppConfig.Env.Port,
	}

	fmt.Printf("server is starting on port: %v", runAppConfig.Env.Port)

	//listen
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
