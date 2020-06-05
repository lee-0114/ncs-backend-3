package main

import (
	api "backend/app/user/ban/api/grpc"
	"backend/app/user/ban/conf"
	"backend/app/user/ban/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
)

func main() {
	// Init
	config := conf.Init()
	log.Init(config.Log)
	srv := service.Init(config)

	// rpc
	server := api.InitServer("tcp", "0.0.0.0:2333", srv)

	log.Info("Ban app started!")

	// cmd
	cmd.Run("Ban", func() {
		server.Stop()
		srv.Close()
		log.Close()
	})
}