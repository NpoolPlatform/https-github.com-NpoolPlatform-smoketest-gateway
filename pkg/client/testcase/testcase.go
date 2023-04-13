//nolint:nolintlint,dupl
package user

import (
	"context"
	"time"

	mgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"

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

func CreateTestCase(ctx context.Context, in *npool.CreateTestCaseReq) (*npool.TestCase, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateTestCase(ctx, &npool.CreateTestCaseRequest{
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
	return info.(*npool.TestCase), nil
}

func GetTestCase(ctx context.Context, id string) (*npool.TestCase, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetTestCase(ctx, &npool.GetTestCaseRequest{
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
	return info.(*npool.TestCase), nil
}

func GetTestCases(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*npool.TestCase, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetTestCases(ctx, &npool.GetTestCasesRequest{
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
	return infos.([]*npool.TestCase), total, nil
}

func DeleteTestCase(ctx context.Context, id string) (*npool.TestCase, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteTestCase(ctx, &npool.DeleteTestCaseRequest{
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
	return info.(*npool.TestCase), nil
}

func UpdateUser(ctx context.Context, in *npool.UpdateTestCaseRequest) (*npool.TestCase, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateTestCase(ctx, &npool.UpdateTestCaseRequest{
			Info: in.Info,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.TestCase), nil
}
