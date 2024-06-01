package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/api/article/internal/logic"
	"github.com/lius-new/blog-backend/api/article/internal/types"
	"github.com/lius-new/blog-backend/api/article/tests"
)

func TestCreateArticle(t *testing.T) {
	ctx := context.Background()

	logicClient := logic.NewCreateArticleLogic(ctx, tests.SVC_CONTEXT)
	logicResp, err := logicClient.CreateArticle(&types.CreateArticleRequest{
		Title:       "测试文章",
		Description: "Desc",
		Content:     "this is cotnent",
		Tags:        []string{"hello"},
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp: ", logicResp, " 创建成功。")
	}
}
