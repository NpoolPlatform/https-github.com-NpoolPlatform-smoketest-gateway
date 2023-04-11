package testcase

import (
	"github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	"google.golang.org/grpc"
)

type Server struct {
	testcase.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	testcase.RegisterMiddlewareServer(server, &Server{})
}
