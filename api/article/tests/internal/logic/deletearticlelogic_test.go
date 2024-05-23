package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/api/article/internal/logic"
	"github.com/lius-new/blog-backend/api/article/internal/types"
	"github.com/lius-new/blog-backend/api/article/tests"
)

func TestDeleteArticle(t *testing.T) {
	ctx := context.Background()

	logicClient := logic.NewDeleteArticleLogic(ctx, tests.SVC_CONTEXT)
	logicResp, err := logicClient.DeleteArticle(&types.DeleteArticleRequest{
		Id: "664ea8fc4d5ef011e75516eb",
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp: ", logicResp, " 删除成功。")
	}

}
