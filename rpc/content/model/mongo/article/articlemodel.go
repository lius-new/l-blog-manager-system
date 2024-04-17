package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

var _ ArticleModel = (*customArticleModel)(nil)

type (
	// ArticleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleModel.
	ArticleModel interface {
		articleModel
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
