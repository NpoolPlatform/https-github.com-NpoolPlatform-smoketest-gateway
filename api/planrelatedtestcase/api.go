package planrelatedtestcase

import (
	"github.com/NpoolPlatform/message/npool/smoketest/mw/v1/planrelatedtestcase"
	"google.golang.org/grpc"
)

type Server struct {
	planrelatedtestcase.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	planrelatedtestcase.RegisterMiddlewareServer(server, &Server{})
}
