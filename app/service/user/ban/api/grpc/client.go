package grpc

import (
	"backend/pkg/log"
	"backend/pkg/rpc"
	"context"
	"google.golang.org/grpc"
)

const (
	ServiceName = "ncs-service-user-ban"
	ServiceAddr = ServiceName + ":2333"
)

var conn *grpc.ClientConn

func InitClient(target string) BanClient {
	conn = rpc.Dial(context.Background(), target, nil)
	return NewBanClient(conn)
}

func Close() {
	if conn != nil {
		if err := conn.Close(); err != nil {
			log.Error(err)
		}
	}
}
