package logic_test

import (
	"context"
	"log"
	"testing"

	"github.com/lius-new/blog-backend/rpc/user/internal/logic"
	"github.com/lius-new/blog-backend/rpc/user/tests"
	"github.com/lius-new/blog-backend/rpc/user/user"
	"github.com/lius-new/blog-backend/rpc/utils/utils"
)

// TestInsert: 测试添加用户, 会调用utils rpc
func TestInsert(t *testing.T) {
	ctx := context.Background()
	insertLogic := logic.NewInsertLogic(ctx, tests.SVC_CONTEXT)

	md5Resp, err := tests.SVC_CONTEXT.Utiler.MD5(ctx, &utils.MD5Reqeust{
		Text: "14569636547",
	})
	if err != nil {
		panic(err)
	}

	resp, err := insertLogic.Insert(&user.InsertUserRequest{
		Username: "lius6666",
		Password: md5Resp.Text,
	})

	if err != nil {
		panic(err)
	}
	log.Println(resp)
}
