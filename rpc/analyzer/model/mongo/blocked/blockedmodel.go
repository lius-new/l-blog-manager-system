package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

var _ BlockedModel = (*customBlockedModel)(nil)

type (
	// BlockedModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBlockedModel.
	BlockedModel interface {
		blockedModel
	}

	customBlockedModel struct {
		*defaultBlockedModel
	}
)

// NewBlockedModel returns a model for the mongo.
func NewBlockedModel(url, db, collection string, c cache.CacheConf) BlockedModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customBlockedModel{
		defaultBlockedModel: newDefaultBlockedModel(conn),
	}
}
