package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestCreateCovers(t *testing.T) {
	ctx := context.Background()

	createCoversLogic := logic.NewCreateCoversLogic(ctx, tests.SVC_CONTEXT)

	resp, err := createCoversLogic.CreateCovers(&content.CreateCoversRequest{
		Content: []string{tests.TEMP_IMAGE_BASE64, tests.TEMP_IMAGE_BASE64 + "sdfhj"},
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(resp.Ids)
	}

}
