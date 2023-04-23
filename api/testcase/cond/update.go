package cond

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	cond1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testcase/cond"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateCond(ctx context.Context, in *npool.UpdateCondRequest) (*npool.UpdateCondResponse, error) {
	req := in.GetInfo()
	handler, err := cond1.NewHandler(
		ctx,
		cond1.WithID(req.ID),
		cond1.WithIndex(req.Index),
		cond1.WithArgumentMap(req.ArgumentMap),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCond",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateCondResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateCond(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCond",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateCondResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCondResponse{
		Info: info,
	}, nil
}
