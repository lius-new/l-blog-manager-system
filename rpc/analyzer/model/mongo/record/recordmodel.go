package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

var _ RecordModel = (*customRecordModel)(nil)

type (
	// RecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecordModel.
	RecordModel interface {
		recordModel
		CountDayRecordNumber(ctx context.Context) (int64, error)
		FindByPage(ctx context.Context, pageNum, pageSize int64) ([]Record, int64, error)
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

func (m *customRecordModel) CountDayRecordNumber(ctx context.Context) (int64, error) {
	return 0, nil
}

func (m *customRecordModel) FindByPage(ctx context.Context, pageNum, pageSize int64) ([]Record, int64, error) {
	return nil, 0, nil
}
