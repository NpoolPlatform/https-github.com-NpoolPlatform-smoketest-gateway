package testcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	testcase1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testcase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteTestCase(ctx context.Context, in *npool.DeleteTestCaseRequest) (*npool.DeleteTestCaseResponse, error) {
	id := in.GetInfo().GetID()
	handler, err := testcase1.NewHandler(
		ctx,
		testcase1.WithID(&id, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTestCase",
			"Req", in,
			"Error", err,
		)
		return &npool.DeleteTestCaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteTestCase(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTestCase",
			"Req", in,
			"Error", err,
		)
		return &npool.DeleteTestCaseResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteTestCaseResponse{
		Info: info,
	}, nil
}
