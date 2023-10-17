package testplan

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	testplan1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testplan"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetTestPlans(ctx context.Context, in *npool.GetTestPlansRequest) (*npool.GetTestPlansResponse, error) {
	handler, err := testplan1.NewHandler(
		ctx,
		testplan1.WithConds(in.GetConds()),
		testplan1.WithOffset(in.GetOffset()),
		testplan1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTestPlans",
			"In", in,
			"Error", err,
		)
		return &npool.GetTestPlansResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetTestPlans(ctx)
	if err != nil {
		return &npool.GetTestPlansResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTestPlansResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetTestPlan(ctx context.Context, in *npool.GetTestPlanRequest) (*npool.GetTestPlanResponse, error) {
	handler, err := testplan1.NewHandler(
		ctx,
		testplan1.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTestPlan",
			"In", in,
			"Error", err,
		)
		return &npool.GetTestPlanResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetTestPlan(ctx)
	if err != nil {
		return &npool.GetTestPlanResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTestPlanResponse{
		Info: info,
	}, nil
}
