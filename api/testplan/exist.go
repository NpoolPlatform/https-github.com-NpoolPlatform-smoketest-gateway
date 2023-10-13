package testplan

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	testplan1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testplan"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistTestPlan(ctx context.Context, in *npool.ExistTestPlanRequest) (*npool.ExistTestPlanResponse, error) {
	handler, err := testplan1.NewHandler(
		ctx,
		testplan1.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTestPlan",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTestPlanResponse{}, status.Error(codes.Aborted, err.Error())
	}
	exist, err := handler.ExistTestPlan(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTestPlan",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTestPlanResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistTestPlanResponse{
		Info: exist,
	}, nil
}
