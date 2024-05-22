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

	resp, err := deleteTagLogic.DeleteTag(&content.DeleteTagRequest{})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
