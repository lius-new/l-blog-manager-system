package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestExistArtilce(t *testing.T) {
	ctx := context.Background()

	existArtilceLogic := logic.NewExistArtilceLogic(ctx, tests.SVC_CONTEXT)

	resp, err := existArtilceLogic.ExistArtilce(&content.ExistArtilceRequest{
		Id: "664d6d1a04c15050fc092f72",
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(resp)
	}
}
