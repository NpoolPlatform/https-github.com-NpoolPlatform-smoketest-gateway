package testcase

import (
	testplantestcase "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/testcase"
	"google.golang.org/grpc"
)

type Server struct {
	testplantestcase.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	testplantestcase.RegisterMiddlewareServer(server, &Server{})
}
