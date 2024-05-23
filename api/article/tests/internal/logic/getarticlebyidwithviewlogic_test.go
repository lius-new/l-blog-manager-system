package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/api/article/internal/logic"
	"github.com/lius-new/blog-backend/api/article/internal/types"
	"github.com/lius-new/blog-backend/api/article/tests"
)

func TestGetArticleByIdWithView(t *testing.T) {
	ctx := context.Background()

	logicClient := logic.NewGetArticleByIdWithViewLogic(ctx, tests.SVC_CONTEXT)
	logicResp, err := logicClient.GetArticleByIdWithView(&types.GetArticleByIdWithViewRequest{
		Id: "664d6d1a04c15050fc092f72",
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp: ", logicResp, " 查询成功。")
	}
}
