package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/logic"
	"github.com/lius-new/blog-backend/rpc/analyzer/tests"
)

func TestMergeRecord(t *testing.T) {

	ctx := context.Background()
	mergeResp, err := logic.NewMergeRecordLogic(ctx, tests.SVC_CONTEXT).MergeRecord(&analyzer.MergeRecordRequest{
		RequestIp: "localhost",
	})

	if err != nil {
		fmt.Println("error :", err)
	} else {
		fmt.Println("resp: : ", mergeResp)
	}
}
