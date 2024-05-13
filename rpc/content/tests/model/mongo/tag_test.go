package mongo_test

import (
	"context"
	"fmt"
	"testing"

	tagModel "github.com/lius-new/blog-backend/rpc/content/model/mongo/tag"
	"github.com/lius-new/blog-backend/rpc/content/tests"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TestFindByName: 根据name查询tag
func TestFindByName(t *testing.T) {
	ctx := context.Background()
	tag, err := tests.SVC_CONTEXT.ModelWithTag.FindByName(ctx, "name")
	if err != nil {
		fmt.Println("错误信息: ", err)
	}
	fmt.Println("查询结果: ", tag)
}

// TestTagFindByPage: 分页查询
func TestTagFindByPage(t *testing.T) {
	ctx := context.Background()
	tags, total, err := tests.SVC_CONTEXT.ModelWithTag.FindByPage(ctx, 0, 2, true)
	if err != nil {
		fmt.Println("错误信息: ", err)
	}

	fmt.Println("total: ", total)
	fmt.Println("articles: ", tags)
}

// TestTagInsertReturnId: 添加tag, 返回id
func TestTagInsertReturnId(t *testing.T) {
	count := 5
	tags := make([]tagModel.Tag, count)

	names := []string{"java", "python", "golang", "python", "Java"}
	for i := 0; i < count; i++ {
		tags = append(tags, tagModel.Tag{
			ID:       primitive.NewObjectID(),
			Name:     names[i],
			Articles: []string{},
			Visiable: true,
		})
	}

	insert := func(t tagModel.Tag) {
		ctx := context.Background()

		id, err := tests.SVC_CONTEXT.ModelWithTag.InsertReturnId(ctx, &t)

		if err != nil {
			fmt.Println("添加失败, 异常tag为: ", t.Name, " 错误信息: ", err)
		}

		fmt.Println("insert return id: ", id)
	}
	for _, v := range tags {
		insert(v)
	}
}
