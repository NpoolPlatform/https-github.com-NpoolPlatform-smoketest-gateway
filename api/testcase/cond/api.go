package cond

import (
	"github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	"google.golang.org/grpc"
)

type Server struct {
	cond.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	cond.RegisterMiddlewareServer(server, &Server{})
}
