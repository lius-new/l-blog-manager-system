package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
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
	return nil, nil
}

func (m *customCoverModel) InsertReturnId(ctx context.Context, data *Cover) (id string, err error) {
	return "", nil
}
