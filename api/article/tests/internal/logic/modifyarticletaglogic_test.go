package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/api/article/internal/logic"
	"github.com/lius-new/blog-backend/api/article/internal/types"
	"github.com/lius-new/blog-backend/api/article/tests"
)

func TestModifyArticleTag(t *testing.T) {

	ctx := context.Background()

	logicClient := logic.NewModifyArticleTagLogic(ctx, tests.SVC_CONTEXT)
	logicResp, err := logicClient.ModifyArticleTag(&types.ModifyArticleTagRequest{
		Id:   "664d6d1a04c15050fc092f72",
		Tags: []string{"666"},
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp: ", logicResp, "更新成功。")
	}
}
