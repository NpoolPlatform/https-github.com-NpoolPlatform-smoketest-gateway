package plantestcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	plantestcase1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testplan/plantestcase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdatePlanTestCase(ctx context.Context, in *npool.UpdatePlanTestCaseRequest) (*npool.UpdatePlanTestCaseResponse, error) {
	req := in.GetInfo()
	handler, err := plantestcase1.NewHandler(
		ctx,
		plantestcase1.WithID(req.ID),
		plantestcase1.WithTestUserID(req.TestUserID),
		plantestcase1.WithTestCaseOutput(req.TestCaseOutput),
		plantestcase1.WithRunDuration(req.RunDuration),
		plantestcase1.WithTestCaseResult(req.Result),
		plantestcase1.WithIndex(req.Index),
		plantestcase1.WithDescription(req.Description),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePlanTestCase",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdatePlanTestCaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdatePlanTestCase(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePlanTestCase",
			"Req", req,
			"Error", err,
		)
		return &npool.UpdatePlanTestCaseResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdatePlanTestCaseResponse{
		Info: info,
	}, nil
}
