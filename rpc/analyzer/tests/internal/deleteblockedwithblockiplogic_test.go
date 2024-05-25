package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/logic"
	"github.com/lius-new/blog-backend/rpc/analyzer/tests"
)

func TestDeleteBlocked(t *testing.T) {
	ctx := context.Background()
	deleteRecordResp, err := logic.NewDeleteBlockedWithBlockIPLogic(ctx, tests.SVC_CONTEXT).DeleteBlockedWithBlockIP(&analyzer.DeleteBlockedWithBlockIPRequest{
		BlockIp: "127.0.0.1",
	})
	if err != nil {
		fmt.Println("error :", err)
	} else {
		fmt.Println(deleteRecordResp)
	}
}
