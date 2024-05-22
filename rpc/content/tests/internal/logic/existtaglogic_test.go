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

	resp, err := existTagLogic.ExistTag(&content.ExistTagRequest{
		Id: "664d6385c2dea61294bd7141",
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(resp)
	}
}
