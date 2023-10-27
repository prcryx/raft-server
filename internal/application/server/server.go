package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/prcryx/raft-server/internal/application/app"
	"github.com/prcryx/raft-server/internal/application/routes"
	"github.com/prcryx/raft-server/internal/domain/types"
)

func NewServer(app *app.App) (*types.Server, error) {
	port, err := strconv.Atoi(app.EnvConfig.Port)
	if err != nil {
		return nil, err
	}
	server := initServer(port)
	routes.SetupRoutes(app, server)

	return server, nil
}

func initServer(port int) *types.Server {
	return &types.Server{
		Port: port,
	}
}

// start a server
func StartServer(server *types.Server) error {
	fmt.Printf("server is listening on: %v\n", server.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", server.Port), server.Router); err != nil {
		fmt.Printf("Server failed: %v", err)
		return err
	}
	return nil
}
