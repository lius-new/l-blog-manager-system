package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

// **cover **
func TestCreateCover(t *testing.T) {
	ctx := context.Background()

	createCoverLogic := logic.NewCreateCoverLogic(ctx, tests.SVC_CONTEXT)

	resp, err := createCoverLogic.CreateCover(&content.CreateCoverRequest{})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
