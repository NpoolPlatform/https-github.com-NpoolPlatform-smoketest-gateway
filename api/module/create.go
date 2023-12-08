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
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateModule",
			"Req", in,
		)
		return &npool.CreateModuleResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := module1.NewHandler(
		ctx,
		module1.WithEntID(req.EntID, false),
		module1.WithName(req.Name, true),
		module1.WithDescription(req.Description, false),
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
