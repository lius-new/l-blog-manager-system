package logic

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestModifyArtilceTag(t *testing.T) {
	ctx := context.Background()

	modifyArtilceTagLogic := logic.NewModifyArtilceTagLogic(ctx, tests.SVC_CONTEXT)

	resp, err := modifyArtilceTagLogic.ModifyArtilceTag(&content.ModifyArticleTagRequest{})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
