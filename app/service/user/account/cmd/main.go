package main

import (
	api "backend/app/service/user/account/api/grpc"
	"backend/app/service/user/account/conf"
	"backend/app/service/user/account/service"
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

	log.Info("Account app started!")

	// cmd
	cmd.Run("Account", func() {
		server.Stop()
		srv.Close()
		log.Close()
	})
}