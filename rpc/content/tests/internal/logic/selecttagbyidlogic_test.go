package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestSelectTagById(t *testing.T) {
	ctx := context.Background()

	selectTagByIdLogic := logic.NewSelectTagByIdLogic(ctx, tests.SVC_CONTEXT)

	resp, err := selectTagByIdLogic.SelectTagById(&content.SelectTagByIdRequest{
		Id: "664d6385c2dea61294bd7141",
	})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
