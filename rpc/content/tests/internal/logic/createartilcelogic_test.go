package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestCreateArtilce(t *testing.T) {
	ctx := context.Background()

	createArtilceLogic := logic.NewCreateArtilceLogic(ctx, tests.SVC_CONTEXT)

	resp, err := createArtilceLogic.CreateArtilce(&content.CreateArticleRequest{
		Title:   "测试文章4",
		Desc:    "这是一片测试的文章",
		Content: "# 标题1",
		Tags:    []string{"test"},
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp: ", resp)
	}
}

// 该测试方法只是在提交数据的时候多添加了covers参数
func TestCreateArtilce2(t *testing.T) {
	ctx := context.Background()

	createArtilceLogic := logic.NewCreateArtilceLogic(ctx, tests.SVC_CONTEXT)

	resp, err := createArtilceLogic.CreateArtilce(&content.CreateArticleRequest{
		Title:   "测试文章2",
		Desc:    "这是一片测试的文章",
		Content: "# 标题1",
		Tags:    []string{"test"},
		Covers:  []string{tests.TEMP_IMAGE_BASE64},
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("resp: ", resp)
	}
}
