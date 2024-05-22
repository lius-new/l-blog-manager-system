package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestModifyArtilceVisiable(t *testing.T) {
	ctx := context.Background()

	modifyArtilceVisiableLogic := logic.NewModifyArtilceVisiableLogic(ctx, tests.SVC_CONTEXT)

	resp, err := modifyArtilceVisiableLogic.ModifyArtilceVisiable(&content.ModifyArticleVisiableRequest{
		Id:       "664d6d1a04c15050fc092f72",
		Visiable: true,
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(resp)
	}
}
