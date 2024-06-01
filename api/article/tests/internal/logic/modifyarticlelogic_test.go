package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/api/article/internal/logic"
	"github.com/lius-new/blog-backend/api/article/internal/types"
	"github.com/lius-new/blog-backend/api/article/tests"
)

func TestModifyArticle(t *testing.T) {
	ctx := context.Background()

	logicClient := logic.NewModifyArticleLogic(ctx, tests.SVC_CONTEXT)
	logicResp, err := logicClient.ModifyArticle(&types.ModifyArticleRequest{
		Id:          "664d6d1a04c15050fc092f72",
		Title:       "TEST3",
		Description: "this",
		Content:     "sdfs",
		Covers:      []string{"ab"},
		Tags:        []string{"abc"},
		Visiable:    true,
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp:", logicResp, "更新成功。")
	}
}
