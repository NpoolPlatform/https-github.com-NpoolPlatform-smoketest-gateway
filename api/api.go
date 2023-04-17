package api

import (
	"context"

	"github.com/NpoolPlatform/ledger-manager/api/detail"
	"github.com/NpoolPlatform/message/npool/servicetmpl"

	"github.com/NpoolPlatform/smoketest-middleware/api/module"
	"github.com/NpoolPlatform/smoketest-middleware/api/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/api/testplan"
	planrelatedtestcase "github.com/NpoolPlatform/smoketest-middleware/api/testplan/testcase"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	servicetmpl.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	servicetmpl.RegisterManagerServer(server, &Server{})
	detail.Register(server)
	testcase.Register(server)
	module.Register(server)
	testplan.Register(server)
	planrelatedtestcase.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := servicetmpl.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := detail.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
