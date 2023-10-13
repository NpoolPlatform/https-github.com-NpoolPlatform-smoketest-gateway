package testcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	testcase1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testcase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateTestCase(ctx context.Context, in *npool.UpdateTestCaseRequest) (*npool.UpdateTestCaseResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateTestCase",
			"Req", req,
		)
		return &npool.UpdateTestCaseResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := testcase1.NewHandler(
		ctx,
		testcase1.WithID(req.ID, true),
		testcase1.WithName(req.Name, false),
		testcase1.WithDescription(req.Description, false),
		testcase1.WithInput(req.Input, false),
		testcase1.WithInputDesc(req.InputDesc, false),
		testcase1.WithExpectation(req.Expectation, false),
		testcase1.WithOutputDesc(req.OutputDesc, false),
		testcase1.WithDeprecated(req.Deprecated, false),
		testcase1.WithTestCaseType(req.TestCaseType, false),
		testcase1.WithTestCaseClass(req.TestCaseClass, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTestCase",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateTestCaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateTestCase(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTestCase",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdateTestCaseResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTestCaseResponse{
		Info: info,
	}, nil
}
