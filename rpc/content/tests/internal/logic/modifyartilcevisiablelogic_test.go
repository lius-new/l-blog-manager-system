package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestModifyArtilceVisiable(t *testing.T) {
	ctx := context.Background()

	modifyArtilceVisiableLogic := logic.NewModifyArtilceVisiableLogic(ctx, tests.SVC_CONTEXT)

	resp, err := modifyArtilceVisiableLogic.ModifyArtilceVisiable(&content.ModifyArticleVisiableRequest{})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
