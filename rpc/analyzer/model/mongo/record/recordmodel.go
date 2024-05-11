package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

var _ RecordModel = (*customRecordModel)(nil)

type (
	// RecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecordModel.
	RecordModel interface {
		recordModel
	}

	customRecordModel struct {
		*defaultRecordModel
	}
)

// NewRecordModel returns a model for the mongo.
func NewRecordModel(url, db, collection string, c cache.CacheConf) RecordModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customRecordModel{
		defaultRecordModel: newDefaultRecordModel(conn),
	}
}
