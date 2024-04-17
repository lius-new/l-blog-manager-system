package svc

import (
	"github.com/lius-new/blog-backend/rpc/content/internal/config"
	articleModel "github.com/lius-new/blog-backend/rpc/content/model/mongo/article"
	tagModel "github.com/lius-new/blog-backend/rpc/content/model/mongo/tag"
)

type ServiceContext struct {
	Config           config.Config
	ModelWithArticle articleModel.ArticleModel
	ModelWithTag     tagModel.TagModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		ModelWithArticle: articleModel.NewArticleModel(c.MongoURL, c.DBName, "article", c.Cache),
		ModelWithTag:     tagModel.NewTagModel(c.MongoURL, c.DBName, "tag", c.Cache),
	}
}
