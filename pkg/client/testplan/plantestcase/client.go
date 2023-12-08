//nolint:nolintlint,dupl
package plantestcase

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"

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

func CreatePlanTestCase(ctx context.Context, in *npool.PlanTestCaseReq) (*npool.PlanTestCase, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreatePlanTestCase(ctx, &npool.CreatePlanTestCaseRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.PlanTestCase), nil
}

func GetPlanTestCase(ctx context.Context, id string) (*npool.PlanTestCase, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetPlanTestCase(ctx, &npool.GetPlanTestCaseRequest{
			EntID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.PlanTestCase), nil
}

func GetPlanTestCases(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.PlanTestCase, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetPlanTestCases(ctx, &npool.GetPlanTestCasesRequest{
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
	return infos.([]*npool.PlanTestCase), total, nil
}

func DeletePlanTestCase(ctx context.Context, id uint32) (*npool.PlanTestCase, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeletePlanTestCase(ctx, &npool.DeletePlanTestCaseRequest{
			Info: &npool.PlanTestCaseReq{
				ID: &id,
			},
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.PlanTestCase), nil
}

func UpdatePlanTestCase(ctx context.Context, in *npool.PlanTestCaseReq) (*npool.PlanTestCase, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdatePlanTestCase(ctx, &npool.UpdatePlanTestCaseRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.PlanTestCase), nil
}

func GetPlanTestCaseOnly(ctx context.Context, conds *npool.Conds) (*npool.PlanTestCase, error) {
	const limit = 2
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetPlanTestCases(ctx, &npool.GetPlanTestCasesRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.PlanTestCase)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.PlanTestCase)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.PlanTestCase)[0], nil
}
