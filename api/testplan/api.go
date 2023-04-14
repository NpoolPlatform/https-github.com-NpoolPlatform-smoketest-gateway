package testplan

import (
	"github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	"google.golang.org/grpc"
)

type Server struct {
	testplan.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	testplan.RegisterMiddlewareServer(server, &Server{})
}
