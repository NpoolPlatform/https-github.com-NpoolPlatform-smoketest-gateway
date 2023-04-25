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
	testcasepb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	testcase1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testcase"
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
	tt = testcasepb.TestCase{
		Name:       uuid.NewString(),
		ModuleName: uuid.NewString(),
		ApiID:      uuid.NewString(),
	}
)

func setupTestCase(t *testing.T) func(*testing.T) {
	handler, err := testcase1.NewHandler(
		context.Background(),
		testcase1.WithName(&tt.Name),
		testcase1.WithModuleName(&tt.ModuleName),
		testcase1.WithApiID(&tt.ApiID),
	)
	assert.Nil(t, err)
	assert.NotNil(t, handler)

	testcase, err := handler.CreateTestCase(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, testcase)

	tt.ID = testcase.ID
	return func(*testing.T) {
		_, _ = handler.DeleteTestCase(context.Background())
	}
}

var (
	ret = npool.Cond{
		TestCaseID:     tt.ID,
		CondTestCaseID: tt.ID,
		CondType:       npool.CondType_PreCondition,
		CondTypeStr:    npool.CondType_PreCondition.String(),
		Index:          100,
		ArgumentMap:    "{}",
	}
)

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
	var (
		index       = uint32(100)
		argumentMap = "{'Username': ''}"
		condType    = npool.CondType_PreCondition
		condTypeStr = npool.CondType_PreCondition.String()
		req         = &npool.CondReq{
			ID:             &ret.ID,
			TestCaseID:     &ret.TestCaseID,
			CondTestCaseID: &ret.CondTestCaseID,
			CondType:       &condType,
			Index:          &index,
			ArgumentMap:    &argumentMap,
		}
	)
	info, err := UpdateCond(context.Background(), req)
	if assert.Nil(t, err) {
		ret.Index = index
		ret.ArgumentMap = argumentMap
		ret.CondTypeStr = condTypeStr
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
	assert.Nil(t, err)
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

	teardown := setupTestCase(t)
	defer teardown(t)

	t.Run("createCond", createCond)
	t.Run("updateCond", updateCond)
	t.Run("getCond", getCond)
	t.Run("getConds", getConds)
	t.Run("deleteCond", deleteCond)
}
