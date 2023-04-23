package plantestcase

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan/plantestcase"
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
		TestPlanID:     uuid.NewString(),
		TestCaseID:     uuid.NewString(),
		TestCaseOutput: "",
		RunDuration:    1,
		Index:          10,
		Description:    uuid.NewString(),
		TestUserID:     uuid.NewString(),
		Result:         npool.TestCaseResult_Passed,
	}
)

func createPlanTestCase(t *testing.T) {
	var (
		req = &npool.PlanTestCaseReq{
			TestPlanID:     &ret.TestPlanID,
			TestCaseID:     &ret.TestCaseID,
			TestCaseOutput: &ret.TestCaseOutput,
			Description:    &ret.Description,
			RunDuration:    &ret.RunDuration,
			Index:          &ret.Index,
			Result:         &ret.Result,
		}
	)

	fmt.Println("TestPlanID: ", *req.TestPlanID)
	fmt.Println("TestCaseID: ", *req.TestCaseID)
	info, err := CreatePlanTestCase(context.Background(), req)
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
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

	t.Run("createPlanTestCase", createPlanTestCase)
	t.Run("getPlanTestCase", getPlanTestCase)
	t.Run("getPlanTestCases", getPlanTestCases)
	t.Run("deletePlanTestCase", deletePlanTestCase)
}
