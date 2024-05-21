package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/lius-new/blog-backend/rpc/content/internal/config"
	articleModel "github.com/lius-new/blog-backend/rpc/content/model/mongo/article"
	coverModel "github.com/lius-new/blog-backend/rpc/content/model/mongo/cover"
	tagModel "github.com/lius-new/blog-backend/rpc/content/model/mongo/tag"
	"github.com/lius-new/blog-backend/rpc/utils/utiler"
)

type ServiceContext struct {
	Config           config.Config
	ModelWithArticle articleModel.ArticleModel
	ModelWithTag     tagModel.TagModel
	ModelWithCover   coverModel.CoverModel
	Utiler           utiler.Utiler
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		ModelWithArticle: articleModel.NewArticleModel(c.MongoURL, c.DBName, "articles", c.Cache),
		ModelWithTag:     tagModel.NewTagModel(c.MongoURL, c.DBName, "tags", c.Cache),
		ModelWithCover:   coverModel.NewCoverModel(c.MongoURL, c.DBName, "covers", c.Cache),
		Utiler:           utiler.NewUtiler(zrpc.MustNewClient(c.Utils)),
	}
}
