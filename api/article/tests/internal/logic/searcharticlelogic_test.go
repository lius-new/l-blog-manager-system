package logic

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/api/article/internal/logic"
	"github.com/lius-new/blog-backend/api/article/internal/types"
	"github.com/lius-new/blog-backend/api/article/tests"
)

func TestSearchArticle(t *testing.T) {
	ctx := context.Background()

	logicClient := logic.NewSearchArticleLogic(ctx, tests.SVC_CONTEXT)
	logicResp, err := logicClient.SearchArticle(&types.SearchArticleRequest{
		Search: "这是一片测试的文章",
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp: ", logicResp, "搜索成功。")
	}
}
