package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestModifyTagVisiable(t *testing.T) {
	ctx := context.Background()

	modifyTagVisiableLogic := logic.NewModifyTagVisiableLogic(ctx, tests.SVC_CONTEXT)

	resp, err := modifyTagVisiableLogic.ModifyTagVisiable(&content.ModifyTagVisiableRequest{})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
