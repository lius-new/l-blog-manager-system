package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/logic"
	"github.com/lius-new/blog-backend/rpc/analyzer/tests"
)

// 数据库中存在封禁信息的
func TestCreateBlocked(t *testing.T) {
	ctx := context.Background()
	createBlockedLogic := logic.NewCreateBlockedLogic(ctx, tests.SVC_CONTEXT)

	blockFunc := func(ip string) {
		createBlockResp, err := createBlockedLogic.CreateBlocked(&analyzer.CreateBlockedRequest{
			BlockIp: ip,
		})

		if err != nil {
			fmt.Println("error : ", err)
		} else {
			fmt.Println(createBlockResp)
		}
	}

	// 对同一个人执行二次封禁
	blockFunc("127.0.0.1")
	blockFunc("127.0.0.1")
}

// 测试第一次封禁(数据库中不存在封禁信息)
func TestCreateBlocked2(t *testing.T) {
	ctx := context.Background()
	createBlockResp, err := logic.NewCreateBlockedLogic(ctx, tests.SVC_CONTEXT).CreateBlocked(&analyzer.CreateBlockedRequest{
		BlockIp: "127.0.0.2",
	})

	if err != nil {
		fmt.Println("error :", err)
	} else {
		fmt.Println(createBlockResp)
	}
}
