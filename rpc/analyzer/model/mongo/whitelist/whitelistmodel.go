package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

var _ WhitelistModel = (*customWhitelistModel)(nil)

type (
	// WhitelistModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWhitelistModel.
	WhitelistModel interface {
		whitelistModel
	}

	customWhitelistModel struct {
		*defaultWhitelistModel
	}
)

// NewWhitelistModel returns a model for the mongo.
func NewWhitelistModel(url, db, collection string, c cache.CacheConf) WhitelistModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customWhitelistModel{
		defaultWhitelistModel: newDefaultWhitelistModel(conn),
	}
}
