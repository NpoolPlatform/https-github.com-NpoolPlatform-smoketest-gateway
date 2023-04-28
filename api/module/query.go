package module

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	module1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/module"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetModules(ctx context.Context, in *npool.GetModulesRequest) (*npool.GetModulesResponse, error) {
	handler, err := module1.NewHandler(
		ctx,
		module1.WithConds(in.GetConds()),
		module1.WithOffset(in.Offset),
		module1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetModules",
			"In", in,
			"Error", err,
		)
		return &npool.GetModulesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetModules(ctx)
	if err != nil {
		return &npool.GetModulesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetModulesResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetModule(ctx context.Context, in *npool.GetModuleRequest) (*npool.GetModuleResponse, error) {
	handler, err := module1.NewHandler(
		ctx,
		module1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetModule",
			"In", in,
			"error", err,
		)
		return &npool.GetModuleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetModule(ctx)
	if err != nil {
		return &npool.GetModuleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetModuleResponse{
		Info: info,
	}, nil
}

func (s *Server) GetModuleConds(ctx context.Context, in *npool.GetModuleCondsRequest) (*npool.GetModuleCondsResponse, error) {
	handler, err := module1.NewHandler(
		ctx,
		module1.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetModule",
			"In", in,
			"error", err,
		)
		return &npool.GetModuleCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, total, err := handler.GetModuleConds(ctx)
	if err != nil {
		return &npool.GetModuleCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetModuleCondsResponse{
		Infos: info,
		Total: total,
	}, nil
}
