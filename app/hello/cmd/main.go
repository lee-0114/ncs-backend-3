package main

import (
	api "backend/app/hello/api/grpc/v1"
	"backend/app/hello/conf"
	"backend/app/hello/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
	"backend/pkg/tracer"
)

// 这里修改服务名字
const serviceName = "hello"

func main() {
	// Init 初始化模块
	config := conf.Init()
	log.Init(config.Log)
	tracer.Init(serviceName)
	srv := service.Init(config)

	// rpc 服务注册
	server := api.InitServer(srv, func() bool {
		return srv.Healthy()
	})

	log.Info(serviceName, "service started!")

	// cmd 控制服务优雅关闭
	cmd.Run(serviceName, func() {
		server.Stop()
		srv.Close()
		tracer.Close()
		log.Close()
	})
}
