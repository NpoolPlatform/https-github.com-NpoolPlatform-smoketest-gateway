package testplan

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan"
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
	ret = npool.TestPlan{
		Name:        uuid.NewString(),
		CreatedBy:   uuid.NewString(),
		Executor:    uuid.NewString(),
		State:       npool.TestPlanState_WaitStart,
		Fails:       0,
		Passes:      0,
		Skips:       0,
		RunDuration: 0,
		Result:      npool.TestResultState_DefaultTestResultState,
		Deadline:    0,
		CreatedAt:   0,
		UpdatedAt:   0,
	}
)

func createTestPlan(t *testing.T) {
	var (
		req = &npool.TestPlanReq{
			Name:      &ret.Name,
			CreatedBy: &ret.CreatedBy,
			Executor:  &ret.Executor,
			Deadline:  &ret.Deadline,
		}
	)

	info, err := CreateTestPlan(context.Background(), req)
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getTestPlan(t *testing.T) {
	info, err := GetTestPlan(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getTestPlans(t *testing.T) {
	infos, _, err := GetTestPlans(context.Background(), &npool.Conds{}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteTestPlan(t *testing.T) {
	info, err := DeleteTestPlan(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetTestPlan(context.Background(), ret.ID)
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

	t.Run("createTestPlan", createTestPlan)
	t.Run("getTestPlan", getTestPlan)
	t.Run("getTestPlans", getTestPlans)
	t.Run("deleteTestPlan", deleteTestPlan)
}
