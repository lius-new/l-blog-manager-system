package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestModifyArtilceTitle(t *testing.T) {
	ctx := context.Background()

	modifyArtilceTitleLogic := logic.NewModifyArtilceTitleLogic(ctx, tests.SVC_CONTEXT)

	resp, err := modifyArtilceTitleLogic.ModifyArtilceTitle(&content.ModifyArticleTitleRequest{})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}