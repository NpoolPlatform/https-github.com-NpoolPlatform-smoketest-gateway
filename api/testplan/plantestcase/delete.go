package plantestcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	plantestcase1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testplan/plantestcase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeletePlanTestCase(ctx context.Context, in *npool.DeletePlanTestCaseRequest) (*npool.DeletePlanTestCaseResponse, error) {
	id := in.GetInfo().GetID()
	handler, err := plantestcase1.NewHandler(
		ctx,
		plantestcase1.WithID(&id, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeletePlanTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePlanTestCaseResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeletePlanTestCase(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeletePlanTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePlanTestCaseResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeletePlanTestCaseResponse{
		Info: info,
	}, nil
}
