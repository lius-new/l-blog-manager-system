package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestModifyTagRemoveArticle(t *testing.T) {
	ctx := context.Background()

	modifyTagRemoveArticleLogic := logic.NewModifyTagRemoveArticleLogic(ctx, tests.SVC_CONTEXT)

	resp, err := modifyTagRemoveArticleLogic.ModifyTagRemoveArticle(&content.ModifyTagRemoveArticleRequest{
		Id:      "664d725218b7418753fe7db8",
		Article: "664d725218b7418753fe7db8",
	})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
