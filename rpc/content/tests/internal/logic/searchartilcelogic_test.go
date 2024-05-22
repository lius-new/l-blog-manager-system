package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestSearchArtilce(t *testing.T) {
	ctx := context.Background()

	searchArtilceLogic := logic.NewSearchArtilceLogic(ctx, tests.SVC_CONTEXT)

	resp, err := searchArtilceLogic.SearchArtilce(&content.SearchArtilceRequest{})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
