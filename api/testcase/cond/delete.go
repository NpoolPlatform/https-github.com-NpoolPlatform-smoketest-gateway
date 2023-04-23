package cond

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	cond1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testcase/cond"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteCond(ctx context.Context, in *npool.DeleteCondRequest) (*npool.DeleteCondResponse, error) {
	id := in.GetInfo().GetID()
	handler, err := cond1.NewHandler(
		ctx,
		cond1.WithID(&id),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCond",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCondResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteCond(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCond",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCondResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteCondResponse{
		Info: info,
	}, nil
}
