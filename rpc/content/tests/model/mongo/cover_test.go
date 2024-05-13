package mongo_test

import (
	"context"
	"fmt"
	"testing"

	coverModel "github.com/lius-new/blog-backend/rpc/content/model/mongo/cover"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

// TestFindOneByHash: 根据hash查询指定Cover图片
func TestFindOneByHash(t *testing.T) {
	ctx := context.Background()
	cover, err := tests.SVC_CONTEXT.ModelWithCover.FindOneByHash(ctx, "hash")
	if err != nil {
		fmt.Println("错误信息: ", err)
	}
	fmt.Println("查询结果: ", cover)
}

// TestFindOneByHash: 插入指定图片并返回指定id
func TestCoverInsertReturnId(t *testing.T) {
	ctx := context.Background()
	coverId, err := tests.SVC_CONTEXT.ModelWithCover.InsertReturnId(ctx, &coverModel.Cover{})
	if err != nil {
		fmt.Println("错误信息: ", err)
	}
	fmt.Println("查询结果: ", coverId)
}
