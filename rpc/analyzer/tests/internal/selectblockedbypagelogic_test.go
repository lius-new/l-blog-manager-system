package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/logic"
	"github.com/lius-new/blog-backend/rpc/analyzer/tests"
)

func TestSelectBlockedByPage(t *testing.T) {

	ctx := context.Background()

	selectBlockedByPageResp, err := logic.NewSelectBlockedByPageLogic(ctx, tests.SVC_CONTEXT).SelectBlockedByPage(&analyzer.SelectBlockedByPageRequest{
		PageNum:  1,
		PageSize: 2,
	})

	if err != nil {
		fmt.Println("error :", err)
	} else {
		fmt.Println(selectBlockedByPageResp)
	}
}
