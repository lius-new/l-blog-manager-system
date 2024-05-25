package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/logic"
	"github.com/lius-new/blog-backend/rpc/analyzer/tests"
)

func TestSelectRecordByPage(t *testing.T) {

	ctx := context.Background()

	selectRecordByPageResp, err := logic.NewSelectRecordByPageLogic(ctx, tests.SVC_CONTEXT).SelectRecordByPage(&analyzer.SelectRecordByPageRequest{
		PageNum:  1,
		PageSize: 2,
	})

	if err != nil {
		fmt.Println("error :", err)
	} else {
		fmt.Println(selectRecordByPageResp)
	}
}
