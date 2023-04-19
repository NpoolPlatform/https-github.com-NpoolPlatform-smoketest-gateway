package plantestcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	planrelatedtestcase1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testplan/plantestcase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint
func (s *Server) CreatePlanTestCase(ctx context.Context, in *npool.CreatePlanTestCaseRequest) (*npool.CreatePlanTestCaseResponse, error) {
	req := in.GetInfo()
	handler, err := planrelatedtestcase1.NewHandler(
		ctx,
		planrelatedtestcase1.WithTestPlanID(req.TestPlanID),
		planrelatedtestcase1.WithTestCaseID(req.TestCaseID),
		planrelatedtestcase1.WithTestUserID(req.TestUserID),
		planrelatedtestcase1.WithTestCaseOutput(req.TestCaseOutput),
		planrelatedtestcase1.WithTestCaseResult(req.TestCaseResult),
		planrelatedtestcase1.WithIndex(req.Index),
		planrelatedtestcase1.WithRunDuration(req.RunDuration),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePlanTestCase",
			"Req", in,
			"Error", err,
		)
		return &npool.CreatePlanTestCaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreatePlanTestCase(ctx)
	if err != nil {
		return &npool.CreatePlanTestCaseResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreatePlanTestCaseResponse{
		Info: info,
	}, nil
}