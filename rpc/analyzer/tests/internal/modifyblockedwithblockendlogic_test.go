package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/logic"
	"github.com/lius-new/blog-backend/rpc/analyzer/tests"
)

func TestModifyBlockedWithBlockEnd(t *testing.T) {
	ctx := context.Background()

	modifyBlockedWithBlockEndResp, err := logic.NewModifyBlockedWithBlockEndLogic(ctx, tests.SVC_CONTEXT).ModifyBlockedWithBlockEnd(&analyzer.ModifyBlockedWithBlockEndRequest{
		BlockIp:  "127.0.0.1",
		BlockEnd: 7200, // 以妙为单位
	})

	if err != nil {
		fmt.Println("error :", err)
	} else {
		fmt.Println(modifyBlockedWithBlockEndResp)
	}
}
