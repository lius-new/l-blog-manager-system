package mongo_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/user/tests"
	"github.com/zeromicro/go-zero/core/logc"
)

func TestLogin(t *testing.T) {
	ctx := context.Background()

	findRes, err := tests.SVC_CONTEXT.Model.FindByUserName(ctx, "lius6666")
	fmt.Println(findRes)

	if err != nil {
		logc.Error(ctx, "test failed: ", err)
	}
	logc.Info(ctx, findRes)
}
