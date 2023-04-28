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
		testplan1.WithID(req.ID),
		testplan1.WithName(req.Name),
		testplan1.WithExecutor(req.Executor),
		testplan1.WithState(req.State),
		testplan1.WithDeadline(req.Deadline),
		testplan1.WithFails(req.Fails),
		testplan1.WithSkips(req.Skips),
		testplan1.WithPasses(req.Passes),
		testplan1.WithResult(req.Result),
		testplan1.WithRunDuration(req.RunDuration),
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
