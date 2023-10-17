package plantestcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	plantestcase1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testplan/plantestcase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreatePlanTestCase(ctx context.Context, in *npool.CreatePlanTestCaseRequest) (*npool.CreatePlanTestCaseResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreatePlanTestCase",
			"Req", in,
		)
		return &npool.CreatePlanTestCaseResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := plantestcase1.NewHandler(
		ctx,
		plantestcase1.WithID(req.ID, false),
		plantestcase1.WithTestPlanID(req.TestPlanID, true),
		plantestcase1.WithTestCaseID(req.TestCaseID, true),
		plantestcase1.WithIndex(req.Index, false),
		plantestcase1.WithInput(req.Input, false),
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
