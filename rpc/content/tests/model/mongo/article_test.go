package mongo_test

import (
	"context"
	"fmt"
	"testing"

	articleModel "github.com/lius-new/blog-backend/rpc/content/model/mongo/article"
	"github.com/lius-new/blog-backend/rpc/content/tests"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TestFindByTitle: 测试根据title查询article
func TestFindByTitle(t *testing.T) {
	ctx := context.Background()

	findRes, err := tests.SVC_CONTEXT.ModelWithArticle.FindByTitle(ctx, "测试文章1")

	if err != nil {
		panic(err)
	}

	fmt.Println(findRes)
}

// TestFindByIds: 批量搜索指定id的文章
func TestFindByIds(t *testing.T) {
	ctx := context.Background()

	articles, err := tests.SVC_CONTEXT.ModelWithArticle.FindByIds(ctx, []string{"id"})

	if err != nil {
		panic(err)
	}

	fmt.Println(articles)
}

// TestFindByPage: 根据分页查询文章
func TestFindByPage(t *testing.T) {
	ctx := context.Background()

	articles, total, err := tests.SVC_CONTEXT.ModelWithArticle.FindByPage(ctx, 0, 2, true)

	if err != nil {
		panic(err)
	}

	fmt.Println("total: ", total)
	fmt.Println("articles: ", articles)
}

// Search: 搜索文章
// 会先查询title字段，如果查询成功就显示指定title的article
// 如果查询title字段对应的article为空，那么就查询desc字段，显示指定desc的article
func TestSearch(t *testing.T) {
	ctx := context.Background()

	search := func(content string) {
		articles, err := tests.SVC_CONTEXT.ModelWithArticle.Search(ctx, content)

		if err != nil {
			panic(err)
		}

		fmt.Println("articles: ", articles)
	}

	// search title
	search("测试文章")

	// search desc
	search("Desc")
}

// TestInsertReturnId: 添加article, 返回id
func TestInsertReturnId(t *testing.T) {
	count := 5
	articles := make([]articleModel.Article, count)

	titles := []string{"测试文章1", "测试文章1", "测试文章2", "测试文章3", "测试文章3"}
	for i := 0; i < count; i++ {
		articles = append(articles,
			articleModel.Article{
				ID:       primitive.NewObjectID(),
				Title:    titles[i],
				Desc:     "Desc",
				Content:  "",
				Tags:     []string{"java", "golang"},
				Covers:   []string{},
				Visiable: true,
			},
		)
	}

	insert := func(a articleModel.Article) {
		ctx := context.Background()

		id, err := tests.SVC_CONTEXT.ModelWithArticle.InsertReturnId(ctx, &a)

		if err != nil {
			fmt.Println("添加失败, 异常artilce为: ", a.Title, " 错误信息: ", err)
		}

		fmt.Println("insert return id: ", id)
	}
	for _, v := range articles {
		insert(v)
	}
}
