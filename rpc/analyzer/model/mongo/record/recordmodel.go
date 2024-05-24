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
		CountDayRecordNumber(ctx context.Context, requestIP string) (int64, error)
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

func (m *customRecordModel) CountDayRecordNumber(ctx context.Context, requestIP string) (int64, error) {
	// 获取现在时间和过去一天时间
	now, dayAgo := getNowAndOldDayTime()

	countCondition := bson.M{
		"requestIP": bson.M{"$regex": requestIP, "$options": "i"},
		"createAt": bson.M{
			"$gte": dayAgo, // 大于或等于dayAgo的时间戳(昨天)
			"$lte": now,    // 小于或等于当前时间戳(现在)
		},
	}

	count, err := m.conn.CountDocuments(ctx, countCondition)

	return count, err
}

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

// getNowAndOldDayTime: 获取当前时间和前一天的时间
// 事实上存本地时间是CST时间, 而数据库是UTC时间，所以如果要比较那么就需要转换为UTC时间。
func getNowAndOldDayTime() (time.Time, time.Time) {
	now := time.Now().UTC()
	dayAgo := now.Add(-24 * time.Hour).UTC()
	return now, dayAgo
}
