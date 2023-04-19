package testcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	testcasecrud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func (h *Handler) UpdateTestCase(ctx context.Context) (info *npool.TestCase, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := testcasecrud.UpdateSet(
			cli.TestCase.UpdateOneID(*h.ID),
			&testcasecrud.Req{
				ID:           h.ID,
				Name:         h.Name,
				Description:  h.Description,
				Input:        h.Input,
				InputDesc:    h.InputDesc,
				Expectation:  h.Expectation,
				TestCaseType: h.TestCaseType,
				Deprecated:   h.Deprecated,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	info, err = h.GetTestCase(ctx)
	if err != nil {
		return nil, err
	}

	return info, nil
}
