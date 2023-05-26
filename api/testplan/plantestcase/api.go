package plantestcase

import (
	plantestcase "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	"google.golang.org/grpc"
)

type Server struct {
	plantestcase.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	plantestcase.RegisterMiddlewareServer(server, &Server{})
}
