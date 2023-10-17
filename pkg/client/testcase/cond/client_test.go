package cond

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	modulemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	testcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	module1 "github.com/NpoolPlatform/smoketest-middleware/pkg/client/module"
	testcase1 "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/testinit"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	ret = npool.Cond{
		TestCaseID:     uuid.NewString(),
		CondTestCaseID: uuid.NewString(),
		CondType:       npool.CondType_PreCondition,
		CondTypeStr:    npool.CondType_PreCondition.String(),
		Index:          100,
		ArgumentMap:    "{}",
	}
	testCase = testcasemwpb.TestCase{
		ApiID:         uuid.NewString(),
		Name:          uuid.NewString(),
		ModuleID:      uuid.NewString(),
		TestCaseType:  testcasemwpb.TestCaseType_Automatic,
		TestCaseClass: testcasemwpb.TestCaseClass_Interface,
	}
)

func setupTestCase(t *testing.T) func(*testing.T) {
	_, err := module1.CreateModule(context.Background(), &modulemwpb.ModuleReq{
		ID:   &testCase.ModuleID,
		Name: &testCase.ModuleID,
	})
	assert.Nil(t, err)

	_, err = testcase1.CreateTestCase(context.Background(), &testcasemwpb.TestCaseReq{
		ID:            &ret.TestCaseID,
		Name:          &testCase.Name,
		ApiID:         &testCase.ApiID,
		ModuleID:      &testCase.ModuleID,
		TestCaseType:  &testCase.TestCaseType,
		TestCaseClass: &testCase.TestCaseClass,
	})
	assert.Nil(t, err)

	_, err = testcase1.CreateTestCase(context.Background(), &testcasemwpb.TestCaseReq{
		ID:            &ret.CondTestCaseID,
		Name:          &testCase.Name,
		ApiID:         &testCase.ApiID,
		ModuleID:      &testCase.ModuleID,
		TestCaseType:  &testCase.TestCaseType,
		TestCaseClass: &testCase.TestCaseClass,
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = testcase1.DeleteTestCase(context.Background(), ret.CondTestCaseID)
		_, _ = testcase1.DeleteTestCase(context.Background(), ret.TestCaseID)
	}
}

func createCond(t *testing.T) {
	var (
		req = &npool.CondReq{
			TestCaseID:     &ret.TestCaseID,
			CondTestCaseID: &ret.CondTestCaseID,
			CondType:       &ret.CondType,
			ArgumentMap:    &ret.ArgumentMap,
			Index:          &ret.Index,
		}
	)

	info, err := CreateCond(context.Background(), req)
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateCond(t *testing.T) {
	ret.Index = uint32(100)
	ret.ArgumentMap = "{\"Username\": \"daiki\"}"
	ret.CondType = npool.CondType_Cleaner
	ret.CondTypeStr = npool.CondType_Cleaner.String()
	var (
		req = &npool.CondReq{
			ID:             &ret.ID,
			TestCaseID:     &ret.TestCaseID,
			CondTestCaseID: &ret.CondTestCaseID,
			CondType:       &ret.CondType,
			Index:          &ret.Index,
			ArgumentMap:    &ret.ArgumentMap,
		}
	)
	info, err := UpdateCond(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getCond(t *testing.T) {
	info, err := GetCond(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getConds(t *testing.T) {
	infos, _, err := GetConds(context.Background(), &npool.Conds{}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteCond(t *testing.T) {
	info, err := DeleteCond(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetCond(context.Background(), ret.ID)
	assert.NotNil(t, err)
	assert.Nil(t, info)
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)
	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	teardown := setupTestCase(t)
	defer teardown(t)

	t.Run("createCond", createCond)
	t.Run("updateCond", updateCond)
	t.Run("getCond", getCond)
	t.Run("getConds", getConds)
	t.Run("deleteCond", deleteCond)
}
