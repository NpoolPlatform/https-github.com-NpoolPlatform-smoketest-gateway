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
	if req == nil {
		logger.Sugar().Errorw(
			"UpdatePlanTestCase",
			"Req", req,
		)
		return &npool.UpdatePlanTestCaseResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := plantestcase1.NewHandler(
		ctx,
		plantestcase1.WithID(req.ID, true),
		plantestcase1.WithTestUserID(req.TestUserID, false),
		plantestcase1.WithOutput(req.Output, false),
		plantestcase1.WithInput(req.Input, false),
		plantestcase1.WithRunDuration(req.RunDuration, false),
		plantestcase1.WithResult(req.Result, false),
		plantestcase1.WithIndex(req.Index, false),
		plantestcase1.WithDescription(req.Description, false),
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
