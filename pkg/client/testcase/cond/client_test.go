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
	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase/cond"
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
		ID:             uuid.NewString(),
		TestCaseID:     uuid.NewString(),
		CondTestCaseID: uuid.NewString(),
		CondType:       npool.CondType_DefaultCondType,
		CondTypeStr:    npool.CondType_DefaultCondType.String(),
		Index:          10,
		ArgumentMap:    "{}",
	}
)

func createCond(t *testing.T) {
	var (
		req = &npool.CondReq{
			ID:             &ret.ID,
			TestCaseID:     &ret.TestCaseID,
			CondTestCaseID: &ret.CondTestCaseID,
			CondType:       &ret.CondType,
			ArgumentMap:    &ret.ArgumentMap,
			Index:          &ret.Index,
		}
	)

	info, err := CreateCond(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
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

	t.Run("createCond", createCond)
	t.Run("getCond", getCond)
	t.Run("getConds", getConds)
	t.Run("deleteCond", deleteCond)
}
