package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ CoverModel = (*customCoverModel)(nil)

type (
	// CoverModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCoverModel.
	CoverModel interface {
		coverModel
		FindOneByHash(ctx context.Context, hash string) (*Cover, error)
		InsertReturnId(ctx context.Context, data *Cover) (id string, err error)
	}

	customCoverModel struct {
		*defaultCoverModel
	}
)

// NewCoverModel returns a model for the mongo.
func NewCoverModel(url, db, collection string, c cache.CacheConf) CoverModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customCoverModel{
		defaultCoverModel: newDefaultCoverModel(conn),
	}
}

// FindOneByHash: 根据hash来查找
func (m *customCoverModel) FindOneByHash(ctx context.Context, hash string) (*Cover, error) {

	var data Cover
	err := m.conn.FindOneNoCache(ctx, &data, bson.M{"hash": hash})

	switch err {
	case nil:
		return &data, nil
	case monc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customCoverModel) InsertReturnId(ctx context.Context, data *Cover) (id string, err error) {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	_, err = m.conn.InsertOneNoCache(ctx, data)
	return data.ID.Hex(), err
}
