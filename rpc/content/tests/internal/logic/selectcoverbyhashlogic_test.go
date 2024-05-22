package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestSelectCoverByHash(t *testing.T) {
	ctx := context.Background()

	selectCoverLogic := logic.NewSelectCoverByHashLogic(ctx, tests.SVC_CONTEXT)

	resp, err := selectCoverLogic.SelectCoverByHash(&content.SelectCoverByHashRequest{
		Hash: "6643e9f2",
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp: ", len(resp.Cover.Content))
	}

}
