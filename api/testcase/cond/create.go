package cond

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	cond1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testcase/cond"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCond(ctx context.Context, in *npool.CreateCondRequest) (*npool.CreateCondResponse, error) {
	req := in.GetInfo()
	handler, err := cond1.NewHandler(
		ctx,
		cond1.WithID(req.ID),
		cond1.WithTestCaseID(req.TestCaseID),
		cond1.WithCondTestCaseID(req.CondTestCaseID),
		cond1.WithCondType(req.CondType),
		cond1.WithArgumentMap(req.ArgumentMap),
		cond1.WithIndex(req.Index),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCond",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateCondResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateCond(ctx)
	if err != nil {
		return &npool.CreateCondResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateCondResponse{
		Info: info,
	}, nil
}
