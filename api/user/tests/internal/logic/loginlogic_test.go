package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/api/user/internal/logic"
	"github.com/lius-new/blog-backend/api/user/internal/types"
	"github.com/lius-new/blog-backend/api/user/tests"
)

// 用户不存在的情况: OK
// 用户密码正确的情况: OK
// 用户密码错误的情况: OK
// 用户凭证不存在的情况: OK
// 用户凭证ID存在,但凭证被删除的情况: OK, 直接生成新的凭证.
// 用户凭证ID不存在,但凭证数据存在的情况: OK, 删除原本的凭证, 生成新的凭证.
func TestLogin(t *testing.T) {
	ctx := context.Background()

	logicClient := logic.NewLoginLogic(ctx, tests.SVC_CONTEXT)

	loginResp, err := logicClient.Login(&types.LoginRequest{
		Username: "lius6666",
		Password: "14569636547",
	})

	if err != nil {
		fmt.Println("match", err)
	}

	fmt.Println("result", loginResp, "err:", err)
}
