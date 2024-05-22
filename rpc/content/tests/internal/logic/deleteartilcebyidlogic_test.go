package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestDeleteArtilceById(t *testing.T) {
	ctx := context.Background()

	deleteArtilceByIdLogic := logic.NewDeleteArtilceByIdLogic(ctx, tests.SVC_CONTEXT)

	resp, err := deleteArtilceByIdLogic.DeleteArtilceById(&content.DeleteArticleRequest{})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
