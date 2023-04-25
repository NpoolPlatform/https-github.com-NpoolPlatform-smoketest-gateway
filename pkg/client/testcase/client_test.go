package testcase

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	apimwcli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"
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
		Protocol:    api.Protocol_HTTP,
		Method:      api.Method_POST,
		Path:        uuid.NewString(),
		PathPrefix:  uuid.NewString(),
	}
)

func setupAPI(t *testing.T) func(*testing.T) {
	info, err := apimwcli.CreateAPI(context.Background(), &api.APIReq{
		ServiceName: &_api.ServiceName,
		Protocol:    &_api.Protocol,
		Method:      &_api.Method,
		Path:        &_api.Path,
		PathPrefix:  &_api.PathPrefix,
	})

	assert.Nil(t, err)
	assert.NotNil(t, info)

	_api.ID = info.ID
	return func(*testing.T) {
		_, _ = apimwcli.DeleteAPI(context.Background(), info.ID)
	}
}

var (
	ret = npool.TestCase{
		Name:            uuid.NewString(),
		Description:     uuid.NewString(),
		ModuleName:      uuid.NewString(),
		ApiID:           _api.ID,
		Input:           "{}",
		InputDesc:       "{}",
		Expectation:     "{}",
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

func updateTestCase(t *testing.T) {
	var (
		name            = "用例名称111"
		description     = "用例描述111"
		input           = "{\"Name\": \"HelloWorld\"}"
		inputDesc       = "{\"Name\": \"string\"}"
		expectation     = "{\"ID\": \"xxxxx\", \"Name\": \"HelloWorld\"}"
		testCaseType    = npool.TestCaseType_Automatic
		testCaseTypeStr = npool.TestCaseType_Automatic.String()
		req             = &npool.TestCaseReq{
			ID:           &ret.ID,
			Name:         &name,
			Description:  &description,
			Input:        &input,
			InputDesc:    &inputDesc,
			Expectation:  &expectation,
			TestCaseType: &testCaseType,
		}
	)

	info, err := UpdateTestCase(context.Background(), req)
	if assert.Nil(t, err) {
		ret.Name = name
		ret.Input = input
		ret.InputDesc = inputDesc
		ret.Description = description
		ret.Expectation = expectation
		ret.TestCaseType = testCaseType
		ret.TestCaseTypeStr = testCaseTypeStr
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
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

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	teardown := setupAPI(t)
	defer teardown(t)

	t.Run("createTestCase", createTestCase)
	t.Run("getTestCase", getTestCase)
	t.Run("getTestCases", getTestCases)
	t.Run("updateTestCase", updateTestCase)
	t.Run("deleteTestCase", deleteTestCase)
}
