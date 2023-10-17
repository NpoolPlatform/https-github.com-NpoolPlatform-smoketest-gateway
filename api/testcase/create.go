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
	if req == nil {
		logger.Sugar().Errorw(
			"CreateTestCase",
			"Req", req,
		)
		return &npool.CreateTestCaseResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}

	handler, err := testcase1.NewHandler(
		ctx,
		testcase1.WithID(req.ID, false),
		testcase1.WithName(req.Name, true),
		testcase1.WithDescription(req.Description, false),
		testcase1.WithModuleID(req.ModuleID, true),
		testcase1.WithApiID(req.ApiID, false),
		testcase1.WithInput(req.Input, false),
		testcase1.WithInputDesc(req.InputDesc, false),
		testcase1.WithExpectation(req.Expectation, false),
		testcase1.WithOutputDesc(req.OutputDesc, false),
		testcase1.WithTestCaseType(req.TestCaseType, true),
		testcase1.WithTestCaseClass(req.TestCaseClass, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTestCase",
			"Req", req,
			"error", err,
		)
		return &npool.CreateTestCaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateTestCase(ctx)
	if err != nil {
		return &npool.CreateTestCaseResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateTestCaseResponse{
		Info: info,
	}, nil
}
