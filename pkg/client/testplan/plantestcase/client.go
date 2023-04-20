//nolint:nolintlint,dupl
package plantestcase

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	mgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan/plantestcase"
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

func CreatePlanTestCase(ctx context.Context, in *mgrpb.PlanTestCaseReq) (*mgrpb.PlanTestCase, error) {
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
	return info.(*mgrpb.PlanTestCase), nil
}

func GetPlanTestCase(ctx context.Context, id string) (*mgrpb.PlanTestCase, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetPlanTestCase(ctx, &npool.GetPlanTestCaseRequest{
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
	return info.(*mgrpb.PlanTestCase), nil
}

func GetPlanTestCases(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*mgrpb.PlanTestCase, uint32, error) {
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
	return infos.([]*mgrpb.PlanTestCase), total, nil
}

func DeletePlanTestCase(ctx context.Context, id string) (*mgrpb.PlanTestCase, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeletePlanTestCase(ctx, &npool.DeletePlanTestCaseRequest{
			Info: &mgrpb.PlanTestCaseReq{
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
	return info.(*mgrpb.PlanTestCase), nil
}

func UpdatePlanTestCase(ctx context.Context, in *mgrpb.PlanTestCaseReq) (*mgrpb.PlanTestCase, error) {
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
	return info.(*mgrpb.PlanTestCase), nil
}
