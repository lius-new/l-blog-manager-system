package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestModifyTagName(t *testing.T) {
	ctx := context.Background()

	modifyTagNameLogic := logic.NewModifyTagNameLogic(ctx, tests.SVC_CONTEXT)

	resp, err := modifyTagNameLogic.ModifyTagName(&content.ModifyTagNameRequest{
		Id:   "664d6385c2dea61294bd7141",
		Name: "Test",
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(resp)
	}

}
