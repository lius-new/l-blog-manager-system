package logic

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestDeleteTag(t *testing.T) {
	ctx := context.Background()

	deleteTagLogic := logic.NewDeleteTagLogic(ctx, tests.SVC_CONTEXT)

	ids := []string{"664d6f02d372f3a71449e0c5", "664d6f2adf9edb778cb9a890", "664d6f2adf9edb778cb9a892", "664d6f322744daa9c9b0b981", "664d6f322744daa9c9b0b983", "664d6f62be707cdddc39d2a6", "664d6f725b1babbd31598646", "664d6f725b1babbd31598648"}
	for _, v := range ids {
		resp, err := deleteTagLogic.DeleteTag(&content.DeleteTagRequest{
			Id: v,
		})
		if err != nil {
			fmt.Println("error: ", err)
		} else {
			fmt.Println("resp: ", resp)
		}
	}
}
