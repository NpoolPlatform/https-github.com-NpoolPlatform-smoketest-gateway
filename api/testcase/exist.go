package testcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	testcase1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testcase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistTestCase(ctx context.Context, in *npool.ExistTestCaseRequest) (*npool.ExistTestCaseResponse, error) {
	handler, err := testcase1.NewHandler(
		ctx,
		testcase1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTestCaseResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistTestCase(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTestCaseResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistTestCaseResponse{
		Info: exist,
	}, nil
}
