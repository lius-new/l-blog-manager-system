package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestSelectArtilceByPage(t *testing.T) {
	ctx := context.Background()

	selectArtilceByPageLogic := logic.NewSelectArtilceByPageLogic(ctx, tests.SVC_CONTEXT)

	resp, err := selectArtilceByPageLogic.SelectArtilceByPage(&content.SelectArticleByPageRequest{
		PageNum:  1,
		PageSize: 3,
		HideShow: true,
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(resp)
	}
}
