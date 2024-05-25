package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/logic"
	"github.com/lius-new/blog-backend/rpc/analyzer/tests"
)

func TestJudgeBlockedByIP(t *testing.T) {

	ctx := context.Background()
	judgeResp, err := logic.NewJudgeBlockedByIPLogic(ctx, tests.SVC_CONTEXT).JudgeBlockedByIP(&analyzer.JudgeBlockedByIPRequest{
		BlockIP: "127.0.0.2",
	})

	if err != nil {
		fmt.Println("error :", err)
	} else {
		fmt.Println("resp:", judgeResp.Block)
	}
}
