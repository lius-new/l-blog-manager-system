package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

var _ TagModel = (*customTagModel)(nil)

type (
	// TagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTagModel.
	TagModel interface {
		tagModel
	}

	customTagModel struct {
		*defaultTagModel
	}
)

// NewTagModel returns a model for the mongo.
func NewTagModel(url, db, collection string, c cache.CacheConf) TagModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customTagModel{
		defaultTagModel: newDefaultTagModel(conn),
	}
}
