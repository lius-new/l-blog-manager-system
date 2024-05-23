package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/api/article/internal/logic"
	"github.com/lius-new/blog-backend/api/article/internal/types"
	"github.com/lius-new/blog-backend/api/article/tests"
)

func TestModifyArticleContent(t *testing.T) {

	ctx := context.Background()

	logicClient := logic.NewModifyArticleContentLogic(ctx, tests.SVC_CONTEXT)
	logicResp, err := logicClient.ModifyArticleContent(&types.ModifyArticleContentRequest{
		Id:      "664d6d1a04c15050fc092f72",
		Content: "# 标题0",
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp: ", logicResp, "更新成功。")
	}
}
