package testcase

import (
	mgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"
)

type Handler struct {
	ID         *string
	ModuleID   *string
	Deprecated *string
	Conds      *mgrpb.Conds
	Offset     *int32
	Limit      *int32
}
