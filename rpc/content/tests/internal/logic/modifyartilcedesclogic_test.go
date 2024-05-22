package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestModifyArtilceDesc(t *testing.T) {
	ctx := context.Background()

	modifyArtilceDescLogic := logic.NewModifyArtilceDescLogic(ctx, tests.SVC_CONTEXT)

	resp, err := modifyArtilceDescLogic.ModifyArtilceDesc(&content.ModifyArticleDescRequest{})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
