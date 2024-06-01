package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/api/article/internal/logic"
	"github.com/lius-new/blog-backend/api/article/internal/types"
	"github.com/lius-new/blog-backend/api/article/tests"
)

func TestViewImageWithBase64(t *testing.T) {
	ctx := context.Background()

	logicClient := logic.NewViewImageWithBase64Logic(ctx, tests.SVC_CONTEXT)
	logicResp, err := logicClient.ViewImageWithBase64(&types.ViewImageRequest{
		Hash: "74289a26",
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp: ", logicResp)
	}
}
