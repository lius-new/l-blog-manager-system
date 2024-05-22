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

	ids := []string{"664d6d0ad14d2fff3dee633a", "664d658c971fe3f0e74f5d15", "664d65e0901054b80054b089"}

	for _, id := range ids {
		resp, err := deleteArtilceByIdLogic.DeleteArtilceById(&content.DeleteArticleRequest{
			Id: id,
		})

		if err != nil {
			fmt.Println("error: ", err)
		} else {
			fmt.Println("resp: ", resp)
		}
	}
}
