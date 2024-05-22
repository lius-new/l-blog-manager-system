package logic

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

// 设置tag
func TestModifyArtilceTag(t *testing.T) {
	ctx := context.Background()

	modifyArtilceTagLogic := logic.NewModifyArtilceTagLogic(ctx, tests.SVC_CONTEXT)

	resp, err := modifyArtilceTagLogic.ModifyArtilceTag(&content.ModifyArticleTagRequest{
		Id:   "664d6d1a04c15050fc092f72",
		Tags: []string{"test"},
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp:", resp)
	}
}

// 设置文章的tag为空
func TestModifyArtilceTagTOEmpty(t *testing.T) {
	ctx := context.Background()

	modifyArtilceTagLogic := logic.NewModifyArtilceTagLogic(ctx, tests.SVC_CONTEXT)

	resp, err := modifyArtilceTagLogic.ModifyArtilceTag(&content.ModifyArticleTagRequest{
		Id:   "664d6d1a04c15050fc092f72",
		Tags: []string{""},
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp:", resp)
	}
}
