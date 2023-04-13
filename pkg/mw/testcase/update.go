package testcase

import (
	"context"

	testmgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	testcasecrud "github.com/NpoolPlatform/smoketest-middleware/pkg/mgr/testcase/crud"
	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateTestCase(ctx context.Context, tx *ent.Tx) error {
	if _, err := testcasecrud.UpdateSet(
		tx.TestCase.UpdateOneID(uuid.MustParse(*h.ID)),
		&testmgrpb.TestCaseReq{
			ID:                 h.ID,
			Name:               h.Name,
			Description:        h.Description,
			Arguments:          h.Arguments,
			ArgTypeDescription: h.ArgTypeDescription,
			ExpectationResult:  h.ExpectationResult,
			TestCaseType:       h.TestCaseType,
			Deprecated:         h.Deprecated,
		}).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateTestCase(ctx context.Context) (info *npool.TestCase, err error) {
	handler := &updateHandler{
		Handler: h,
	}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateTestCase(_ctx, tx); err != nil {
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
