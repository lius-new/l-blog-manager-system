package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestSelectArtilceByTag(t *testing.T) {
	ctx := context.Background()

	selectArtilceByTagLogic := logic.NewSelectArtilceByTagLogic(ctx, tests.SVC_CONTEXT)

	resp, err := selectArtilceByTagLogic.SelectArtilceByTag(&content.SelectArticleByTagRequest{
		Tag:      "test",
		PageNum:  1,
		PageSize: 3,
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp: ", resp)
	}
}
