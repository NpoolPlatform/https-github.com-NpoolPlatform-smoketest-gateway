package plantestcase

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
	apimgrpb "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"
	testcasepb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	testplanpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	testcase1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testcase"
	testplan1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testplan"
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
	return func(*testing.T) {
		_, _ = handler.DeleteTestCase(context.Background())
	}
}

var (
	tp = testplanpb.TestPlan{
		Name:      uuid.NewString(),
		CreatedBy: uuid.NewString(),
		Executor:  uuid.NewString(),
		State:     testplanpb.TestPlanState_WaitStart,
		StateStr:  testplanpb.TestPlanState_WaitStart.String(),
	}
)

func setupTestPlan(t *testing.T) func(*testing.T) {
	handler, err := testplan1.NewHandler(
		context.Background(),
		testplan1.WithName(&tp.Name),
		testplan1.WithCreatedBy(&tp.CreatedBy),
		testplan1.WithExecutor(&tp.Executor),
		testplan1.WithState(&tp.State),
	)
	assert.Nil(t, err)
	assert.NotNil(t, handler)

	testplan, err := handler.CreateTestPlan(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, testplan)

	ret.TestPlanID = testplan.ID
	return func(*testing.T) {
		_, _ = handler.DeleteTestPlan(context.Background())
	}
}

var (
	ret = npool.PlanTestCase{
		Index: 10,
		Input: "{}",
	}
)

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

	apiTearDown := setupAPI(t)
	defer apiTearDown(t)

	tcTeardown := setupTestCase(t)
	defer tcTeardown(t)

	tpTeardown := setupTestPlan(t)
	defer tpTeardown(t)

	patch := monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createPlanTestCase", createPlanTestCase)
	t.Run("updatePlanTestCase", updatePlanTestCase)
	t.Run("getPlanTestCase", getPlanTestCase)
	t.Run("getPlanTestCases", getPlanTestCases)
	t.Run("deletePlanTestCase", deletePlanTestCase)

	patch.Unpatch()
}
