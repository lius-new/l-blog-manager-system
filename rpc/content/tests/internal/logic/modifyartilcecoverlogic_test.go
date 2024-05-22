package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestModifyArtilceCover(t *testing.T) {
	ctx := context.Background()

	modifyArtilceCoverLogic := logic.NewModifyArtilceCoverLogic(ctx, tests.SVC_CONTEXT)

	resp, err := modifyArtilceCoverLogic.ModifyArtilceCover(&content.ModifyArticleCoverRequest{
		Id:     "664d6d1a04c15050fc092f72",
		Covers: []string{tests.TEMP_IMAGE_BASE64 + "abc"},
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(resp)
	}
}
