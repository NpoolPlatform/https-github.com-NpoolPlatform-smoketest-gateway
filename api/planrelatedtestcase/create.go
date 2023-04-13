package planrelatedtestcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/planrelatedtestcase"
	planrelatedtestcase1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/planrelatedtestcase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreatePlanRelatedTestCase(ctx context.Context, in *npool.CreatePlanRelatedTestCaseRequest) (*npool.CreatePlanRelatedTestCaseResponse, error) {
	req := in.GetInfo()
	result := req.TestCaseResult.String()
	handler, err := planrelatedtestcase1.NewHandler(
		ctx,
		planrelatedtestcase1.WithTestPlanID(req.TestPlanID),
		planrelatedtestcase1.WithTestCaseID(req.TestCaseID),
		planrelatedtestcase1.WithTestUserID(req.TestUserID),
		planrelatedtestcase1.WithTestCaseOutput(req.TestCaseOutput),
		planrelatedtestcase1.WithTestCaseResult(&result),
		planrelatedtestcase1.WithIndex(req.Index),
		planrelatedtestcase1.WithRunDuration(req.Index),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePlanRelatedTestCase",
			"Req", in,
			"Error", err,
		)
		return &npool.CreatePlanRelatedTestCaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreatePlanRelatedTestCase(ctx)
	if err != nil {
		return &npool.CreatePlanRelatedTestCaseResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreatePlanRelatedTestCaseResponse{
		Info: info,
	}, nil
}
