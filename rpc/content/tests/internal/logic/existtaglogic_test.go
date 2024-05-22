package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestExistTag(t *testing.T) {
	ctx := context.Background()

	existTagLogic := logic.NewExistTagLogic(ctx, tests.SVC_CONTEXT)

	resp, err := existTagLogic.ExistTag(&content.ExistTagRequest{})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
