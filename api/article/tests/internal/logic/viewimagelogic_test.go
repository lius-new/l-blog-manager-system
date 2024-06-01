package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/api/article/internal/logic"
	"github.com/lius-new/blog-backend/api/article/internal/types"
	"github.com/lius-new/blog-backend/api/article/tests"
)

func TestViewImage(t *testing.T) {
	ctx := context.Background()

	logicClient := logic.NewViewImageLogic(ctx, tests.SVC_CONTEXT)
	logicResp, err := logicClient.ViewImage(&types.ViewImageRequest{
		Hash: "sdfsdf",
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp: ", logicResp)
	}
}
