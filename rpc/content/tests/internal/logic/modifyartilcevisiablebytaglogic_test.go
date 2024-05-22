package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestModifyArtilceVisiableByTag(t *testing.T) {
	ctx := context.Background()

	modifyArtilceVisiableByTagLogic := logic.NewModifyArtilceVisiableByTagLogic(ctx, tests.SVC_CONTEXT)

	resp, err := modifyArtilceVisiableByTagLogic.ModifyArtilceVisiableByTag(&content.ModifyArticleVisiableByTagRequest{
		TagId:    "664d6385c2dea61294bd7141",
		Visiable: false,
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp: ", resp)
	}
}
