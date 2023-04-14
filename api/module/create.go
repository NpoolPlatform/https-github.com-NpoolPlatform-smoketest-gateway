package module

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	module1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/module"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateModule(ctx context.Context, in *npool.CreateModuleRequest) (*npool.CreateModuleResponse, error) {
	handler, err := module1.NewHandler(
		ctx,
		module1.WithName(&in.Name),
		module1.WithDescription(in.Description),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateModule",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateModuleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateModule(ctx)
	if err != nil {
		return &npool.CreateModuleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateModuleResponse{
		Info: info,
	}, nil
}
