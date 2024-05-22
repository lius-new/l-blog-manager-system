package logic

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestCreateTags(t *testing.T) {
	ctx := context.Background()

	createTagsLogic := logic.NewCreateTagsLogic(ctx, tests.SVC_CONTEXT)

	resp, err := createTagsLogic.CreateTags(&content.CreateTagsRequest{})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
