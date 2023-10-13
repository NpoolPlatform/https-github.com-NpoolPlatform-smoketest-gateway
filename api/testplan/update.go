package testplan

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	testplan1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testplan"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateTestPlan(ctx context.Context, in *npool.UpdateTestPlanRequest) (*npool.UpdateTestPlanResponse, error) {
	req := in.GetInfo()
	handler, err := testplan1.NewHandler(
		ctx,
		testplan1.WithID(req.ID, true),
		testplan1.WithName(req.Name, false),
		testplan1.WithExecutor(req.Executor, false),
		testplan1.WithState(req.State, false),
		testplan1.WithDeadline(req.Deadline, false),
		testplan1.WithFails(req.Fails, false),
		testplan1.WithSkips(req.Skips, false),
		testplan1.WithPasses(req.Passes, false),
		testplan1.WithResult(req.Result, false),
		testplan1.WithRunDuration(req.RunDuration, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTestPlan",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateTestPlanResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateTestPlan(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTestPlan",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateTestPlanResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTestPlanResponse{
		Info: info,
	}, nil
}
