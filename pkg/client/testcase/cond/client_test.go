package cond

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	apicli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	apimgrpb "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
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
	_api = apimgrpb.API{
		ServiceName: uuid.NewString(),
		Protocol:    apimgrpb.Protocol_HTTP,
		Method:      apimgrpb.Method_POST,
		Path:        uuid.NewString(),
		PathPrefix:  uuid.NewString(),
	}
)

func setupAPI(t *testing.T) func(*testing.T) {
	info, err := apicli.CreateAPI(context.Background(), &apimgrpb.APIReq{
		ServiceName: &_api.ServiceName,
		Protocol:    &_api.Protocol,
		Method:      &_api.Method,
		Path:        &_api.Path,
		PathPrefix:  &_api.PathPrefix,
		Domains:     _api.Domains,
	})

	assert.Nil(t, err)
	assert.NotNil(t, info)

	tc.ApiID = info.ID
	return func(*testing.T) {
		_, _ = apicli.DeleteAPI(context.Background(), info.ID)
	}
}

var (
	tc = testcasepb.TestCase{
		Name:       uuid.NewString(),
		ModuleName: uuid.NewString(),
	}
)

func setupTestCase(t *testing.T) func(*testing.T) {
	handler, err := testcase1.NewHandler(
		context.Background(),
		testcase1.WithName(&tc.Name),
		testcase1.WithModuleName(&tc.ModuleName),
		testcase1.WithApiID(&tc.ApiID),
	)
	assert.Nil(t, err)
	assert.NotNil(t, handler)

	info, err := handler.CreateTestCase(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info)

	ret.TestCaseID = info.ID
	ret.CondTestCaseID = info.ID
	return func(*testing.T) {
		_, _ = handler.DeleteTestCase(context.Background())
	}
}

var (
	ret = npool.Cond{
		TestCaseID:     tc.ID,
		CondTestCaseID: tc.ID,
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

	apiTeardown := setupAPI(t)
	defer apiTeardown(t)

	teardown := setupTestCase(t)
	defer teardown(t)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createCond", createCond)
	t.Run("updateCond", updateCond)
	t.Run("getCond", getCond)
	t.Run("getConds", getConds)
	t.Run("deleteCond", deleteCond)
}
