package logic

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestSelectArtilceById(t *testing.T) {
	ctx := context.Background()

	selectArtilceByIdLogic := logic.NewSelectArtilceByIdLogic(ctx, tests.SVC_CONTEXT)

	resp, err := selectArtilceByIdLogic.SelectArtilceById(&content.SelectArticleByIdRequest{
		Id: "664d6d1a04c15050fc092f72",
	})

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(resp)
}
