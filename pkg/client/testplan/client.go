//nolint:nolintlint,dupl
package testplan

import (
	"context"
	"time"

	mgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"

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

func CreateTestPlan(ctx context.Context, in *mgrpb.TestPlanReq) (*mgrpb.TestPlan, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateTestPlan(ctx, &npool.CreateTestPlanRequest{Info: in})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*mgrpb.TestPlan), nil
}

func GetTestPlan(ctx context.Context, id string) (*mgrpb.TestPlan, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetTestPlan(ctx, &npool.GetTestPlanRequest{
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
	return info.(*mgrpb.TestPlan), nil
}

func GetTestPlans(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*mgrpb.TestPlan, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetTestPlans(ctx, &npool.GetTestPlansRequest{
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
	return infos.([]*mgrpb.TestPlan), total, nil
}

func DeleteTestPlan(ctx context.Context, id string) (*mgrpb.TestPlan, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteTestPlan(ctx, &npool.DeleteTestPlanRequest{
			Info: &mgrpb.TestPlanReq{
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
	return info.(*mgrpb.TestPlan), nil
}

func UpdateTestPlan(ctx context.Context, in *mgrpb.TestPlanReq) (*mgrpb.TestPlan, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateTestPlan(ctx, &npool.UpdateTestPlanRequest{
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
	return info.(*mgrpb.TestPlan), nil
}

func ExistTestPlan(ctx context.Context, id string) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistTestPlan(ctx, &npool.ExistTestPlanRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}
