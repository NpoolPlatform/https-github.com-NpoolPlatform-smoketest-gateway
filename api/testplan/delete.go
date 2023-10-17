package testplan

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	testplan1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testplan"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteTestPlan(ctx context.Context, in *npool.DeleteTestPlanRequest) (*npool.DeleteTestPlanResponse, error) {
	id := in.GetInfo().GetID()
	handler, err := testplan1.NewHandler(
		ctx,
		testplan1.WithID(&id, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTestPlan",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTestPlanResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteTestPlan(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTestPlan",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTestPlanResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteTestPlanResponse{
		Info: info,
	}, nil
}
