package plantestcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	plantestcase1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testplan/plantestcase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetPlanTestCases(ctx context.Context, in *npool.GetPlanTestCasesRequest) (*npool.GetPlanTestCasesResponse, error) {
	handler, err := plantestcase1.NewHandler(
		ctx,
		plantestcase1.WithConds(in.GetConds()),
		plantestcase1.WithOffset(in.GetOffset()),
		plantestcase1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPlanTestCases",
			"In", in,
			"Error", err,
		)
		return &npool.GetPlanTestCasesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetPlanTestCases(ctx)
	if err != nil {
		return &npool.GetPlanTestCasesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetPlanTestCasesResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetPlanTestCase(ctx context.Context, in *npool.GetPlanTestCaseRequest) (*npool.GetPlanTestCaseResponse, error) {
	handler, err := plantestcase1.NewHandler(
		ctx,
		plantestcase1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPlanTestCase",
			"In", in,
			"Error", err,
		)
		return &npool.GetPlanTestCaseResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetPlanTestCase(ctx)
	if err != nil {
		return &npool.GetPlanTestCaseResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetPlanTestCaseResponse{
		Info: info,
	}, nil
}
