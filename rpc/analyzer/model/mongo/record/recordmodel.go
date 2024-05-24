package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ RecordModel = (*customRecordModel)(nil)

type (
	// RecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecordModel.
	RecordModel interface {
		recordModel
		CountScopeTimeRecordNumber(ctx context.Context, requestIP string, scope time.Duration) (int64, error)
		FindByPage(ctx context.Context, pageNum, pageSize int64) ([]Record, int64, error)
		DeleteScopeTimeRecord(ctx context.Context, requestIP string, scope time.Duration) (int64, error) // 删除指定时间范围内的日志
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

// 统计指定时间间隔的record信息
func (m *customRecordModel) CountScopeTimeRecordNumber(ctx context.Context, requestIP string, scope time.Duration) (int64, error) {
	// 获取现在时间和过去一天时间
	// 事实上存本地时间是CST时间, 而数据库是UTC时间，所以如果要比较那么就需要转换为UTC时间。
	now := time.Now().UTC()
	past := now.Add(-1 * scope).UTC()

	countCondition := bson.M{
		"requestIP": bson.M{"$regex": requestIP, "$options": "i"},
		"createAt": bson.M{
			"$gte": past, // 大于或等于past的时间戳(指定时间间隔之前的时间)
			"$lte": now,  // 小于或等于当前时间戳(现在)
		},
	}

	count, err := m.conn.CountDocuments(ctx, countCondition)

	return count, err
}

// 分页查询
func (m *customRecordModel) FindByPage(ctx context.Context, pageNum, pageSize int64) ([]Record, int64, error) {
	findOptions := options.Find()
	// 如果传入参数小于等于0那么就设置为1
	if pageNum <= 0 {
		pageNum = 1
	}

	findOptions.SetLimit(pageSize)
	findOptions.SetSkip(pageSize * (pageNum - 1))
	findOptions.SetSort(bson.M{"updateAt": -1}) // 根据时间降序排序

	records := make([]Record, pageNum)

	// 查询
	err := m.conn.Find(ctx, &records, bson.M{}, findOptions)
	switch err {
	case nil:
		total, _ := m.conn.CountDocuments(ctx, bson.M{})
		return records, total, nil
	case monc.ErrNotFound:
		return nil, 0, ErrNotFound
	default:
		return nil, 0, err
	}
}

func (m *customRecordModel) DeleteScopeTimeRecord(ctx context.Context, requestIP string, scope time.Duration) (int64, error) {
	now := time.Now().UTC()
	past := now.Add(-1 * scope).UTC()

	deleteCondition := bson.M{
		"requestIP": bson.M{"$regex": requestIP, "$options": "i"},
		"createAt": bson.M{
			"$gte": past, // 大于或等于past的时间戳(指定时间间隔之前的时间)
			"$lte": now,  // 小于或等于当前时间戳(现在)
		},
	}

	count, err := m.conn.DeleteMany(ctx, deleteCondition)
	return count, err
}
