package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
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

func (m *defaultUserModel) FindByPage(ctx context.Context, pageNum, pageSize int64) ([]User, int64, error) {
	findOptions := options.Find()
	if pageNum <= 0 {
		pageNum = 1
	}
	findOptions.SetLimit(pageSize)
	findOptions.SetSkip(pageSize * (pageNum - 1))
	findOptions.SetSort(bson.M{"time": -1})

	data := make([]User, 0)
	err := m.conn.Find(ctx, &data, bson.D{{}}, findOptions)

	total, _ := m.conn.CountDocuments(ctx, bson.D{{}})

	switch err {
	case nil:
		return data, total, nil
	case monc.ErrNotFound:
		return nil, 0, ErrNotFound
	default:
		return nil, 0, err
	}
}
