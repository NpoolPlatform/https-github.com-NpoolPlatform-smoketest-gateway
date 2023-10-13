package testplan

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	testplan1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testplan"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateTestPlan(ctx context.Context, in *npool.CreateTestPlanRequest) (*npool.CreateTestPlanResponse, error) {
	req := in.GetInfo()

	handler, err := testplan1.NewHandler(
		ctx,
		testplan1.WithID(req.ID, false),
		testplan1.WithName(req.Name, true),
		testplan1.WithCreatedBy(req.CreatedBy, true),
		testplan1.WithExecutor(req.Executor, true),
		testplan1.WithDeadline(req.Deadline, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTestPlan",
			"Req", req,
			"Error", err,
		)
		return &npool.CreateTestPlanResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateTestPlan(ctx)
	if err != nil {
		return &npool.CreateTestPlanResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateTestPlanResponse{
		Info: info,
	}, nil
}
