package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestExistArtilce(t *testing.T) {
	ctx := context.Background()

	existArtilceLogic := logic.NewExistArtilceLogic(ctx, tests.SVC_CONTEXT)

	resp, err := existArtilceLogic.ExistArtilce(&content.ExistArtilceRequest{})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
