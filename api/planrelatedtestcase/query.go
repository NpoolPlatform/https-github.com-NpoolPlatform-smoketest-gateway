package planrelatedtestcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/planrelatedtestcase"
	planrelatedtestcase1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/planrelatedtestcase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint
func (s *Server) GetPlanRelatedTestCases(ctx context.Context, in *npool.GetPlanRelatedTestCasesRequest) (*npool.GetPlanRelatedTestCasesResponse, error) {
	handler, err := planrelatedtestcase1.NewHandler(
		ctx,
		planrelatedtestcase1.WithConds(
			in.GetConds(),
			in.GetOffset(),
			in.GetLimit(),
		),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPlanRelatedTestCases",
			"In", in,
			"Error", err,
		)
		return &npool.GetPlanRelatedTestCasesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetPlanRelatedTestCases(ctx)
	if err != nil {
		return &npool.GetPlanRelatedTestCasesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetPlanRelatedTestCasesResponse{
		Infos: infos,
		Total: total,
	}, nil
}

//nolint
func (s *Server) GetPlanRelatedTestCase(ctx context.Context, in *npool.GetPlanRelatedTestCaseRequest) (*npool.GetPlanRelatedTestCaseResponse, error) {
	handler, err := planrelatedtestcase1.NewHandler(
		ctx,
		planrelatedtestcase1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPlanRelatedTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.GetPlanRelatedTestCaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetPlanRelatedTestCase(ctx)
	if err != nil {
		return &npool.GetPlanRelatedTestCaseResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetPlanRelatedTestCaseResponse{
		Info: info,
	}, nil
}
