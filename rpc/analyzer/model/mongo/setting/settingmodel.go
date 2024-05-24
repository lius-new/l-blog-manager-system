package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ SettingModel = (*customSettingModel)(nil)

type (
	// SettingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSettingModel.
	SettingModel interface {
		settingModel
		FindLastSetting(ctx context.Context) (Setting, error) // 获取最后一条数据
	}

	customSettingModel struct {
		*defaultSettingModel
	}
)

// NewSettingModel returns a model for the mongo.
func NewSettingModel(url, db, collection string, c cache.CacheConf) SettingModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customSettingModel{
		defaultSettingModel: newDefaultSettingModel(conn),
	}
}

// 获取最后一条数据, 注意要忽略not found的错误， 因为可能存在集合中没有任何数据的情况, 此时会设置默认值（具体看checkrecordmergelogic.go）。
func (m *customSettingModel) FindLastSetting(ctx context.Context) (Setting, error) {
	opts := options.FindOne().SetSort(bson.M{"updateAt": -1})
	var setting Setting
	err := m.conn.FindOneNoCache(ctx, &setting, bson.M{}, opts)
	return setting, err
}
