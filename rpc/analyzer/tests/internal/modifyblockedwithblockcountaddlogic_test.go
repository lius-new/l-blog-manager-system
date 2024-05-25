package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/logic"
	"github.com/lius-new/blog-backend/rpc/analyzer/tests"
)

func TestModifyBlockedWithBlockCountAdd(t *testing.T) {
	ctx := context.Background()
	modifyBlockedWithBlockCountAddResp, err := logic.NewModifyBlockedWithBlockCountAddLogic(ctx, tests.SVC_CONTEXT).ModifyBlockedWithBlockCountAdd(&analyzer.ModifyBlockedWithBlockCountAddRequest{
		BlockIp: "127.0.0.1",
	})

	if err != nil {
		fmt.Println("error :", err)
	} else {
		fmt.Println(modifyBlockedWithBlockCountAddResp)
	}
}
