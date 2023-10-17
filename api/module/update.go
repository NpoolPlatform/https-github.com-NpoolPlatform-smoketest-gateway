package module

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	module1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/module"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateModule(ctx context.Context, in *npool.UpdateModuleRequest) (*npool.UpdateModuleResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateModule",
			"Req", req,
		)
		return &npool.UpdateModuleResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := module1.NewHandler(
		ctx,
		module1.WithID(req.ID, true),
		module1.WithName(req.Name, false),
		module1.WithDescription(req.Description, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateModule",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateModuleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateModule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateModule",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateModuleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateModuleResponse{
		Info: info,
	}, nil
}
