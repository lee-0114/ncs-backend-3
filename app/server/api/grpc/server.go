package grpc

import (
	"backend/pkg/rpc"
	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
}

func InitServer(network string, address string, srv ServerServer) (s *Server) {
	s = &Server{}
	s.server = rpc.NewServer(&rpc.ServerConfig{
		Network: network,
		Addr:    address,
		RegFunc: func(s *grpc.Server) {
			RegisterServerServer(s, srv)
		},
	})
	return
}

func (s *Server) Stop() {
	if s.server != nil {
		s.server.Stop()
	}
}
