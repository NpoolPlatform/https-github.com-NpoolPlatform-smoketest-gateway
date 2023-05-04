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

	handler, err := testcase1.NewHandler(
		ctx,
		testcase1.WithName(req.Name),
		testcase1.WithDescription(req.Description),
		testcase1.WithModuleName(req.ModuleName),
		testcase1.WithApiID(req.ApiID),
		testcase1.WithInput(req.Input),
		testcase1.WithInputDesc(req.InputDesc),
		testcase1.WithExpectation(req.Expectation),
		testcase1.WithOutputDesc(req.OutputDesc),
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
