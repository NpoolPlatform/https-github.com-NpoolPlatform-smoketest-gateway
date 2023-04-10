package testcase

import (
	testcasemgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"
)

type Handler struct {
	ID                *string
	Name              *string
	Description       *string
	ModuleID          *string
	ModuleName        *string
	ApiID             *string
	Arguments         *string
	ExpectationResult *string
	TestCaseType      *testcasemgrpb.TestCaseType
	Deprecated        *bool
	Conds             *testcasemgrpb.Conds
	Offset            *int32
	Limit             *int32
}
