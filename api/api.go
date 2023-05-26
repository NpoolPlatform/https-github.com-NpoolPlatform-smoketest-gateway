package api

import (
	"context"

	smoketest "github.com/NpoolPlatform/message/npool/smoketest/mw/v1"

	"github.com/NpoolPlatform/smoketest-middleware/api/module"
	"github.com/NpoolPlatform/smoketest-middleware/api/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/api/testcase/cond"
	"github.com/NpoolPlatform/smoketest-middleware/api/testplan"
	plantestcase "github.com/NpoolPlatform/smoketest-middleware/api/testplan/plantestcase"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	smoketest.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	smoketest.RegisterMiddlewareServer(server, &Server{})
	testcase.Register(server)
	cond.Register(server)
	module.Register(server)
	testplan.Register(server)
	plantestcase.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := smoketest.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
