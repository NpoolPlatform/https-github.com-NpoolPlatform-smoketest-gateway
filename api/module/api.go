package module

import (
	"github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	"google.golang.org/grpc"
)

type Server struct {
	module.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	module.RegisterMiddlewareServer(server, &Server{})
}
