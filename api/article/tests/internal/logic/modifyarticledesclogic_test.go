package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/api/article/internal/logic"
	"github.com/lius-new/blog-backend/api/article/internal/types"
	"github.com/lius-new/blog-backend/api/article/tests"
)

func TestModifyArticleDesc(t *testing.T) {
	ctx := context.Background()

	logicClient := logic.NewModifyArticleDescLogic(ctx, tests.SVC_CONTEXT)
	logicResp, err := logicClient.ModifyArticleDesc(&types.ModifyArticleDescRequest{
		Id:   "664d6d1a04c15050fc092f72",
		Desc: "desc...",
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp: ", logicResp, "更新成功。")
	}
}
