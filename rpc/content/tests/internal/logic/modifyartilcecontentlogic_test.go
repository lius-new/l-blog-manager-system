package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestModifyArtilceContent(t *testing.T) {
	ctx := context.Background()

	modifyArtilceContentLogic := logic.NewModifyArtilceContentLogic(ctx, tests.SVC_CONTEXT)

	resp, err := modifyArtilceContentLogic.ModifyArtilceContent(&content.ModifyArticleContentRequest{})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
