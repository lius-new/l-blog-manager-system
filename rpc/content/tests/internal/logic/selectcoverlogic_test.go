package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestSelectCover(t *testing.T) {
	ctx := context.Background()

	selectCoverLogic := logic.NewSelectCoverLogic(ctx, tests.SVC_CONTEXT)

	resp, err := selectCoverLogic.SelectCover(&content.SelectCoverRequest{})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}