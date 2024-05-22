package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestCreateTag(t *testing.T) {
	ctx := context.Background()

	createTagLogic := logic.NewCreateTagLogic(ctx, tests.SVC_CONTEXT)

	resp, err := createTagLogic.CreateTag(&content.CreateTagRequest{
		Name: "hello",
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(resp)
	}
}
