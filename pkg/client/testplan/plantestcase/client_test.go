package plantestcase

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	modulemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	testcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	testplanmwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	module1 "github.com/NpoolPlatform/smoketest-middleware/pkg/client/module"
	testcase1 "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"
	testplan1 "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan"
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
	ret = npool.PlanTestCase{
		ID:               uuid.NewString(),
		TestPlanID:       uuid.NewString(),
		TestCaseID:       uuid.NewString(),
		TestCaseName:     uuid.NewString(),
		TestCaseType:     testcasemwpb.TestCaseType_Automatic,
		TestCaseTypeStr:  testcasemwpb.TestCaseType_Automatic.String(),
		TestCaseClass:    testcasemwpb.TestCaseClass_Interface,
		TestCaseClassStr: testcasemwpb.TestCaseClass_Interface.String(),
		ModuleID:         uuid.NewString(),
		ModuleName:       uuid.NewString(),
		Index:            10,
		Input:            "{}",
	}
	testCase = testcasemwpb.TestCase{
		ApiID:         uuid.NewString(),
		Name:          ret.TestCaseName,
		ModuleID:      ret.ModuleID,
		ModuleName:    ret.ModuleName,
		TestCaseType:  ret.TestCaseType,
		TestCaseClass: ret.TestCaseClass,
	}
	testPlan = testplanmwpb.TestPlan{
		Name:      uuid.NewString(),
		CreatedBy: uuid.NewString(),
		Executor:  uuid.NewString(),
		State:     testplanmwpb.TestPlanState_WaitStart,
		StateStr:  testplanmwpb.TestPlanState_WaitStart.String(),
		Deadline:  uint32(time.Now().Unix() + 60*60*24),
	}
)

func setupTestCase(t *testing.T) func(*testing.T) {
	_, err := module1.CreateModule(context.Background(), &modulemwpb.ModuleReq{
		ID:   &testCase.ModuleID,
		Name: &testCase.ModuleName,
	})
	assert.Nil(t, err)

	_, err = testcase1.CreateTestCase(context.Background(), &testcasemwpb.TestCaseReq{
		ID:            &ret.TestCaseID,
		ModuleID:      &testCase.ModuleID,
		ApiID:         &testCase.ApiID,
		Name:          &testCase.Name,
		TestCaseType:  &testCase.TestCaseType,
		TestCaseClass: &testCase.TestCaseClass,
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = testcase1.DeleteTestCase(context.Background(), ret.TestCaseID)
		_, _ = module1.DeleteModule(context.Background(), testCase.ModuleID)
	}
}

func setupTestPlan(t *testing.T) func(*testing.T) {
	_, err := testplan1.CreateTestPlan(context.Background(), &testplanmwpb.TestPlanReq{
		ID:        &ret.TestPlanID,
		Name:      &testPlan.Name,
		CreatedBy: &testPlan.CreatedBy,
		Executor:  &testPlan.Executor,
		State:     &testPlan.State,
		Deadline:  &testPlan.Deadline,
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = testplan1.DeleteTestPlan(context.Background(), ret.TestPlanID)
	}
}

func createPlanTestCase(t *testing.T) {
	var (
		req = &npool.PlanTestCaseReq{
			TestPlanID: &ret.TestPlanID,
			TestCaseID: &ret.TestCaseID,
			Index:      &ret.Index,
			Input:      &ret.Input,
		}
	)

	info, err := CreatePlanTestCase(context.Background(), req)
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.TestUserID = info.TestUserID
		ret.Output = info.Output
		ret.Description = info.Description
		ret.RunDuration = info.RunDuration
		ret.Result = info.Result
		ret.ResultStr = info.Result.String()
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updatePlanTestCase(t *testing.T) {
	ret.Index = 1000
	ret.Description = "{}"
	ret.Input = "{\"Username\": \"hello\"}"
	ret.Output = "{}"
	ret.RunDuration = 100
	ret.Result = npool.TestCaseResult_Passed
	ret.ResultStr = npool.TestCaseResult_Passed.String()
	req := &npool.PlanTestCaseReq{
		ID:          &ret.ID,
		Index:       &ret.Index,
		Description: &ret.Description,
		Input:       &ret.Input,
		Output:      &ret.Output,
		Result:      &ret.Result,
		RunDuration: &ret.RunDuration,
	}

	info, err := UpdatePlanTestCase(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getPlanTestCase(t *testing.T) {
	info, err := GetPlanTestCase(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getPlanTestCases(t *testing.T) {
	infos, _, err := GetPlanTestCases(context.Background(), &npool.Conds{}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deletePlanTestCase(t *testing.T) {
	info, err := DeletePlanTestCase(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetPlanTestCase(context.Background(), ret.ID)
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

	tcTeardown := setupTestCase(t)
	defer tcTeardown(t)

	tpTeardown := setupTestPlan(t)
	defer tpTeardown(t)

	t.Run("createPlanTestCase", createPlanTestCase)
	t.Run("updatePlanTestCase", updatePlanTestCase)
	t.Run("getPlanTestCase", getPlanTestCase)
	t.Run("getPlanTestCases", getPlanTestCases)
	t.Run("deletePlanTestCase", deletePlanTestCase)
}
