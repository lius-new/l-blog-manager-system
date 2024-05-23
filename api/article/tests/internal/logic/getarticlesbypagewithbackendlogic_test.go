package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/api/article/internal/logic"
	"github.com/lius-new/blog-backend/api/article/internal/types"
	"github.com/lius-new/blog-backend/api/article/tests"
)

func TestGetArticlesByPageWithBackend(t *testing.T) {
	ctx := context.Background()

	logicClient := logic.NewGetArticlesByPageWithBackendLogic(ctx, tests.SVC_CONTEXT)
	logicResp, err := logicClient.GetArticlesByPageWithBackend(&types.GetArticlesByPageWithBackendRequest{
		PageNum:  1,
		PageSize: 2,
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp: ", logicResp, " 查询成功。")
	}
}
