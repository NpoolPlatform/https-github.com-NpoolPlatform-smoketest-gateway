//nolint:nolintlint,dupl
package planrelatedtestcase

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/testcase"

	servicename "github.com/NpoolPlatform/smoketest-middleware/pkg/servicename"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second) //nolint
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return fn(_ctx, cli)
}

func CreatePlanRelatedTestCase(ctx context.Context, in *npool.CreatePlanRelatedTestCaseRequest) (*npool.PlanRelatedTestCase, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreatePlanRelatedTestCase(ctx, &npool.CreatePlanRelatedTestCaseRequest{Info: in.GetInfo()})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.PlanRelatedTestCase), nil
}

func GetPlanRelatedTestCase(ctx context.Context, id string) (*npool.PlanRelatedTestCase, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetPlanRelatedTestCase(ctx, &npool.GetPlanRelatedTestCaseRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.PlanRelatedTestCase), nil
}

func GetPlanRelatedTestCases(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.PlanRelatedTestCase, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetPlanRelatedTestCases(ctx, &npool.GetPlanRelatedTestCasesRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}

		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.PlanRelatedTestCase), total, nil
}
