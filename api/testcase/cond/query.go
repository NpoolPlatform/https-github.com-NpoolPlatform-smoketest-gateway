package cond

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	cond1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testcase/cond"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetConds(ctx context.Context, in *npool.GetCondsRequest) (*npool.GetCondsResponse, error) {
	handler, err := cond1.NewHandler(
		ctx,
		cond1.WithConds(in.GetConds()),
		cond1.WithOffset(in.Offset),
		cond1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetConds",
			"In", in,
			"Error", err,
		)
		return &npool.GetCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetConds(ctx)
	if err != nil {
		return &npool.GetCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCondsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetCond(ctx context.Context, in *npool.GetCondRequest) (*npool.GetCondResponse, error) {
	handler, err := cond1.NewHandler(
		ctx,
		cond1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCond",
			"In", in,
			"error", err,
		)
		return &npool.GetCondResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetCond(ctx)
	if err != nil {
		return &npool.GetCondResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCondResponse{
		Info: info,
	}, nil
}
