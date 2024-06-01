package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindByUserName(ctx context.Context, username string) (*User, error)
		FindByPage(ctx context.Context, pageNum, pageSize int64) ([]User, int64, error)
		UpdateStatus(ctx context.Context, data *User) (*mongo.UpdateResult, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the mongo.
func NewUserModel(url, db, collection string, c cache.CacheConf) UserModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customUserModel{
		defaultUserModel: newDefaultUserModel(conn),
	}
}

// FindByUserName: 根据用户名来查询用户
func (m *customUserModel) FindByUserName(ctx context.Context, username string) (*User, error) {
	var data User
	err := m.conn.FindOneNoCache(ctx, &data, bson.M{"username": username})

	switch err {
	case nil:
		return &data, nil
	case monc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindByPage: 分页查询
func (m *defaultUserModel) FindByPage(
	ctx context.Context,
	pageNum, pageSize int64,
) ([]User, int64, error) {
	findOptions := options.Find()
	if pageNum <= 0 {
		pageNum = 1
	}
	findOptions.SetLimit(pageSize)
	findOptions.SetSkip(pageSize * (pageNum - 1))
	findOptions.SetSort(bson.M{"time": -1})

	data := make([]User, 0)
	err := m.conn.Find(ctx, &data, bson.M{}, findOptions)

	total, _ := m.conn.CountDocuments(ctx, bson.M{})

	switch err {
	case nil:
		return data, total, nil
	case monc.ErrNotFound:
		return nil, 0, ErrNotFound
	default:
		return nil, 0, err
	}
}

// UpdateStatus: 更新用户状态(是否可用)
func (m *customUserModel) UpdateStatus(
	ctx context.Context,
	data *User,
) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()
	key := prefixUserCacheKey + data.ID.Hex()
	res, err := m.conn.UpdateOne(
		ctx,
		key,
		bson.M{"_id": data.ID},
		bson.M{"$set": bson.M{"status": data.Status}},
	)
	return res, err
}
