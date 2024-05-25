package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/logic"
	"github.com/lius-new/blog-backend/rpc/analyzer/tests"
)

func TestSelectBlockedByBlockIP(t *testing.T) {

	ctx := context.Background()

	selectBlockedByBlockIP, err := logic.NewSelectBlockedByBlockIPLogic(ctx, tests.SVC_CONTEXT).SelectBlockedByBlockIP(&analyzer.SelectBlockedByBlockIPRequest{
		BlockIp: "127.0.0.1",
	})

	if err != nil {
		fmt.Println("error :", err)
	} else {
		fmt.Println(selectBlockedByBlockIP)
	}
}
