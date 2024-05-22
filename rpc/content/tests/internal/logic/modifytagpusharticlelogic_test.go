package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestModifyTagPushArticle(t *testing.T) {
	ctx := context.Background()

	modifyTagPushArticleLogic := logic.NewModifyTagPushArticleLogic(ctx, tests.SVC_CONTEXT)

	resp, err := modifyTagPushArticleLogic.ModifyTagPushArticle(&content.ModifyTagPushArticleRequest{
		Id:      "664d725218b7418753fe7db8",
		Article: "664d725218b7418753fe7db8",
	})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
