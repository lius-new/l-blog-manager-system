package logic

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestDeleteCover(t *testing.T) {
	ctx := context.Background()

	deleteCoverLogic := logic.NewDeleteCoverLogic(ctx, tests.SVC_CONTEXT)

	resp, err := deleteCoverLogic.DeleteCover(&content.DeleteCoverRequest{Id: "664d72e3a6baa1344d7bb196"})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
