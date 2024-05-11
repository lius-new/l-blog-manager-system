package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ ArticleModel = (*customArticleModel)(nil)

type (
	// ArticleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleModel.
	ArticleModel interface {
		articleModel
		FindByTitle(ctx context.Context, title string) (*Article, error)
		FindByIds(ctx context.Context, ids []string) ([]Article, error)
		FindByPage(
			ctx context.Context,
			pageSize, pageNum int64,
			hideShow bool,
		) ([]Article, int64, error)
		Search(ctx context.Context, search string) ([]Article, error)
		InsertReturnId(ctx context.Context, data *Article) (id string, err error)
	}

	customArticleModel struct {
		*defaultArticleModel
	}
)

// NewArticleModel returns a model for the mongo.
func NewArticleModel(url, db, collection string, c cache.CacheConf) ArticleModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customArticleModel{
		defaultArticleModel: newDefaultArticleModel(conn),
	}
}

// FindByTitle; 根据title查询
func (m *customArticleModel) FindByTitle(ctx context.Context, title string) (*Article, error) {
	key := prefixArticleCacheKey + title
	var article Article
	err := m.conn.FindOne(ctx, key, &article, bson.M{"title": title})
	switch err {
	case nil:
		return &article, nil
	case monc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customArticleModel) FindByIds(ctx context.Context, ids []string) ([]Article, error) {
	return nil, nil
}

// FindByPage: 查看文章, 分页
// hideShow : 是否包含显示被设置为visiable = false的文章
// 写入到缓存中
func (m *customArticleModel) FindByPage(
	ctx context.Context,
	pageSize, pageNum int64,
	hideShow bool,
) ([]Article, int64, error) {
	articles := make([]Article, pageNum)

	findOptions := options.Find()
	// 如果传入参数小于等于0那么就设置为1
	if pageNum <= 0 {
		pageNum = 1
	}

	findOptions.SetLimit(pageSize)
	findOptions.SetSkip(pageSize * (pageNum - 1))
	findOptions.SetSort(bson.M{"updateAt": -1}) // 根据时间降序排序

	// 查询
	var err error
	if hideShow {
		err = m.conn.Find(ctx, &articles, bson.M{}, findOptions)
	} else {
		err = m.conn.Find(ctx, &articles, bson.M{"visiable": true}, findOptions)
	}

	// TODO: 应该将查询结果缓存起来, 但是因为如果缓存我就要去修改更新和删除的代码.

	switch err {
	case nil:
		total, _ := m.conn.CountDocuments(ctx, bson.M{})
		return articles, total, nil
	case monc.ErrNotFound:
		return nil, 0, ErrNotFound
	default:
		return nil, 0, err
	}
}

// Search: 搜索文章
func (m *customArticleModel) Search(ctx context.Context, search string) ([]Article, error) {
	articles := make([]Article, 0)

	regex := primitive.Regex{Pattern: search, Options: "i"}

	// 查询, 如果查询成功那么就使用title字段
	err := m.conn.Find(ctx, &articles, bson.M{"title": regex})

	// 如果title查询的数量为空那么就查询desc字段
	if len(articles) == 0 {
		err = m.conn.Find(ctx, &articles, bson.M{"desc": regex})
	}

	return articles, err
}

// InsertReturnId: 向数据库中添加Article数据，和生成的方法一样但是额外返回一个id
func (m *customArticleModel) InsertReturnId(
	ctx context.Context,
	data *Article,
) (id string, err error) {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	key := prefixArticleCacheKey + data.ID.Hex()
	_, err = m.conn.InsertOne(ctx, key, data)
	return data.ID.Hex(), err
}
