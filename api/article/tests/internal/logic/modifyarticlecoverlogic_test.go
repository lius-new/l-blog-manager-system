package logic

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/api/article/internal/logic"
	"github.com/lius-new/blog-backend/api/article/internal/types"
	"github.com/lius-new/blog-backend/api/article/tests"
)

func TestModifyArticleCover(t *testing.T) {

	ctx := context.Background()

	logicClient := logic.NewModifyArticleCoverLogic(ctx, tests.SVC_CONTEXT)
	logicResp, err := logicClient.ModifyArticleCover(&types.ModifyArticleCoverRequest{
		Id:     "664d6d1a04c15050fc092f72",
		Covers: []string{"shjdfh"},
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp: ", logicResp, "更新成功。")
	}
}
