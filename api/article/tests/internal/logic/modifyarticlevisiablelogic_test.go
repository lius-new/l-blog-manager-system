package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/api/article/internal/logic"
	"github.com/lius-new/blog-backend/api/article/internal/types"
	"github.com/lius-new/blog-backend/api/article/tests"
)

func TestModifyArticleVisiable(t *testing.T) {
	ctx := context.Background()

	logicClient := logic.NewModifyArticleVisiableLogic(ctx, tests.SVC_CONTEXT)
	logicResp, err := logicClient.ModifyArticleVisiable(&types.ModifyArticleVisiableRequest{
		Id:       "664d6d1a04c15050fc092f72",
		Visiable: true,
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp: ", logicResp, "更新成功。")
	}
}
