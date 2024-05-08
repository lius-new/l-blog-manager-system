package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/user/internal/logic"
	"github.com/lius-new/blog-backend/rpc/user/tests"
	"github.com/lius-new/blog-backend/rpc/user/user"
)

func TestSelectByPage(t *testing.T) {
	ctx := context.Background()
	logicClient := logic.NewSelectByPageLogic(ctx, tests.SVC_CONTEXT)

	userResp, err := logicClient.SelectByPage(&user.SelectUserByPageRequest{
		PageNum:  0,
		PageSize: 10,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(userResp)
}
