package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestSelectTagByPage(t *testing.T) {
	ctx := context.Background()

	selectTagByPageLogic := logic.NewSelectTagByPageLogic(ctx, tests.SVC_CONTEXT)

	resp, err := selectTagByPageLogic.SelectTagByPage(&content.SelectTagByPageRequest{
		PageNum:  1,
		PageSize: -1,
		HideShow: true,
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(resp)
	}
}
