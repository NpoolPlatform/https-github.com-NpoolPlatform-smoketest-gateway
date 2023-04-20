package module

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	module1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/module"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteModule(ctx context.Context, in *npool.DeleteModuleRequest) (*npool.DeleteModuleResponse, error) {
	id := in.GetInfo().GetID()
	handler, err := module1.NewHandler(
		ctx,
		module1.WithID(&id),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteModule",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteModuleResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteModule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteModule",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteModuleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteModuleResponse{
		Info: info,
	}, nil
}
