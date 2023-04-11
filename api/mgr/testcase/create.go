package testcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	testcase1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testcase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateTestCase(ctx context.Context, in *npool.CreateTestCaseRequest) (*npool.CreateTestCaseResponse, error) {
	req := in.GetInfo()
	if req.ModuleID == nil && req.ModuleName == nil {
		logger.Sugar().Errorw("CreateTestCase", "Need ModuleID or ModuleName")
	}

	handler, err := testcase1.NewHandler(
		ctx,
		testcase1.WithName(&req.Name),
		testcase1.WithAPIID(&req.ApiID),
		testcase1.WithModuleID(req.ModuleID),
		testcase1.WithModuleName(req.ModuleName),
		testcase1.WithExpectationResult(&req.ExpectationResult),
		testcase1.WithExpectationResult(&req.ExpectationResult),
		testcase1.WithArguments(&req.Arguments),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTestCase",
			"Req", req,
			"error", err,
		)
		return &npool.CreateTestCaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.Create(ctx)
	if err != nil {
		return &npool.CreateTestCaseResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateTestCaseResponse{
		Info: info,
	}, nil
}
