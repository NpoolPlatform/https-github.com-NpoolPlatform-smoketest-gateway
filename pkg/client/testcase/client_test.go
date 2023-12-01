package testcase

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	modulemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	module1 "github.com/NpoolPlatform/smoketest-middleware/pkg/client/module"
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
	ret = npool.TestCase{
		Name:             uuid.NewString(),
		Description:      uuid.NewString(),
		ModuleID:         uuid.NewString(),
		ModuleName:       uuid.NewString(),
		ApiID:            uuid.NewString(),
		Input:            "{}",
		InputDesc:        "{}",
		Expectation:      "{}",
		OutputDesc:       "{}",
		Deprecated:       false,
		TestCaseType:     npool.TestCaseType_Automatic,
		TestCaseTypeStr:  npool.TestCaseType_Automatic.String(),
		TestCaseClass:    npool.TestCaseClass_Interface,
		TestCaseClassStr: npool.TestCaseClass_Interface.String(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	info, err := module1.CreateModule(context.Background(), &modulemwpb.ModuleReq{
		EntID: &ret.ModuleID,
		Name:  &ret.ModuleName,
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = module1.DeleteModule(context.Background(), info.ID)
	}
}

func createTestCase(t *testing.T) {
	var (
		req = &npool.TestCaseReq{
			Name:          &ret.Name,
			Description:   &ret.Description,
			ModuleID:      &ret.ModuleID,
			ApiID:         &ret.ApiID,
			Input:         &ret.Input,
			InputDesc:     &ret.InputDesc,
			Expectation:   &ret.Expectation,
			OutputDesc:    &ret.OutputDesc,
			TestCaseType:  &ret.TestCaseType,
			TestCaseClass: &ret.TestCaseClass,
		}
	)

	info, err := CreateTestCase(context.Background(), req)
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.EntID = info.EntID
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
	info, err := GetTestCase(context.Background(), ret.EntID)
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
		EntID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
		ModuleID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.ModuleID,
		},
		ApiID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.ApiID,
		},
		Deprecated: &basetypes.BoolVal{
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

	info, err = GetTestCase(context.Background(), ret.EntID)
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

	teardown := setup(t)
	defer teardown(t)

	t.Run("createTestCase", createTestCase)
	t.Run("updateTestCase", updateTestCase)
	t.Run("getTestCase", getTestCase)
	t.Run("getTestCases", getTestCases)
	t.Run("getTestCaseConds", getTestCaseConds)
	t.Run("deleteTestCase", deleteTestCase)
}
