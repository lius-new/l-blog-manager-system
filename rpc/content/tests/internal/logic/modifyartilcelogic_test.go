package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestModifyArtilce(t *testing.T) {

	ctx := context.Background()

	modifyArtilceDescLogic := logic.NewModifyArtilceLogic(ctx, tests.SVC_CONTEXT)

	resp, err := modifyArtilceDescLogic.ModifyArtilce(&content.ModifyArticleRequest{
		Id:          "664d6d1a04c15050fc092f72",
		Title:       "TEST1",
		Description: "this",
		Content:     "sdfs",
		Covers:      []string{"abc"},
		Tags:        []string{"abc"},
		Visiable:    false,
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp:", resp)
	}

}
