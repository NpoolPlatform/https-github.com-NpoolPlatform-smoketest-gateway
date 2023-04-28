package testcase

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	apicli "github.com/NpoolPlatform/basal-manager/pkg/client/api"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
	apimgrpb "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
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
	})

	assert.Nil(t, err)
	assert.NotNil(t, info)

	ret.ApiID = info.ID
	return func(*testing.T) {
		_, _ = apicli.DeleteAPI(context.Background(), info.ID)
	}
}

var (
	ret = npool.TestCase{
		Name:            uuid.NewString(),
		Description:     uuid.NewString(),
		ModuleName:      uuid.NewString(),
		Input:           "{}",
		InputDesc:       "{}",
		Expectation:     "{}",
		OutputDesc:      "{}",
		Deprecated:      false,
		TestCaseType:    npool.TestCaseType_DefaultTestCaseType,
		TestCaseTypeStr: npool.TestCaseType_DefaultTestCaseType.String(),
	}
)

func createTestCase(t *testing.T) {
	var (
		req = &npool.TestCaseReq{
			Name:        &ret.Name,
			Description: &ret.Description,
			ModuleName:  &ret.ModuleName,
			ApiID:       &ret.ApiID,
			Input:       &ret.Input,
			InputDesc:   &ret.InputDesc,
			Expectation: &ret.Expectation,
			OutputDesc:  &ret.OutputDesc,
		}
	)

	info, err := CreateTestCase(context.Background(), req)
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.ModuleID = info.ModuleID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateTestCase(t *testing.T) {
	ret.Name = "用例名称111"
	ret.Description = "用例描述111"
	ret.Input = "{\"Name\": \"HelloWorld\"}"
	ret.InputDesc = "{\"Name\": \"string\"}"
	ret.Expectation = "{\"ID\": \"xxxxx\", \"Name\": \"HelloWorld\"}"
	ret.OutputDesc = "{\"ID\": \"string\", \"Name\": \"string\"}"
	ret.TestCaseType = npool.TestCaseType_Automatic
	ret.TestCaseTypeStr = npool.TestCaseType_Automatic.String()
	var (
		req = &npool.TestCaseReq{
			ID:           &ret.ID,
			Name:         &ret.Name,
			Description:  &ret.Description,
			Input:        &ret.Input,
			InputDesc:    &ret.InputDesc,
			Expectation:  &ret.Expectation,
			OutputDesc:   &ret.OutputDesc,
			TestCaseType: &ret.TestCaseType,
		}
	)

	info, err := UpdateTestCase(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getTestCase(t *testing.T) {
	info, err := GetTestCase(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getTestCases(t *testing.T) {
	infos, _, err := GetTestCases(context.Background(), &npool.Conds{}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func getTestCaseConds(t *testing.T) {
	infos, _, err := GetTestCases(context.Background(), &npool.Conds{
		ID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.ID,
		},
		ModuleID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.ModuleID,
		},
		ApiID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.ApiID,
		},
		TestCaseType: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.TestCaseType.String(),
		},
		Deprecated: &commonpb.BoolVal{
			Op:    cruder.EQ,
			Value: ret.Deprecated,
		},
	}, 0, 1)

	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteTestCase(t *testing.T) {
	info, err := DeleteTestCase(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetTestCase(context.Background(), ret.ID)
	assert.NotNil(t, err)
	assert.Nil(t, info)
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	teardown := setupAPI(t)
	defer teardown(t)

	patch := monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createTestCase", createTestCase)
	t.Run("updateTestCase", updateTestCase)
	t.Run("getTestCase", getTestCase)
	t.Run("getTestCases", getTestCases)
	t.Run("getTestCaseConds", getTestCaseConds)
	t.Run("deleteTestCase", deleteTestCase)

	patch.Unpatch()
}
