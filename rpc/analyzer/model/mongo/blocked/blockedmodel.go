package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ BlockedModel = (*customBlockedModel)(nil)

type (
	// BlockedModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBlockedModel.
	BlockedModel interface {
		blockedModel
		FindByBlockIP(ctx context.Context, blockIp string) (Blocked, error)
		FindByPage(ctx context.Context, pageNum, pageSize int64) ([]Blocked, int64, error)
		ModifyBlockByBlockIPWithCount(ctx context.Context, blockIp string, count int64) error                               // 在已有的block数据上加1
		ModifyBlockByBlockIPWithBlockend(ctx context.Context, blockIp string, endTime time.Time) error                      // 在已有的block数据上设置Blockend
		ModifyBlockByBlockIPWithCountAndBlockend(ctx context.Context, blockIp string, endTime time.Time, count int64) error // 在已有的block数据上设置Blockend
		DeleteBlockByBlockIP(ctx context.Context, blockIp string) (*mongo.UpdateResult, error)                              // 根据ip删除blocked
	}

	customBlockedModel struct {
		*defaultBlockedModel
	}
)

// NewBlockedModel returns a model for the mongo.
func NewBlockedModel(url, db, collection string, c cache.CacheConf) BlockedModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customBlockedModel{
		defaultBlockedModel: newDefaultBlockedModel(conn),
	}
}

func (m *customBlockedModel) FindByBlockIP(ctx context.Context, blockIp string) (Blocked, error) {
	var blocked Blocked
	err := m.conn.FindOneNoCache(ctx, &blocked, bson.M{"blockIp": blockIp})
	switch err {
	case nil:
		return blocked, nil
	case monc.ErrNotFound:
		return blocked, ErrNotFound
	default:
		return blocked, err
	}
}
func (m *customBlockedModel) FindByPage(ctx context.Context, pageNum, pageSize int64) ([]Blocked, int64, error) {
	return nil, 0, nil
}

func (m *customBlockedModel) ModifyBlockByBlockIPWithCount(ctx context.Context, blockIp string, count int64) error {
	return nil
}
func (m *customBlockedModel) ModifyBlockByBlockIPWithBlockend(ctx context.Context, blockIp string, endTime time.Time) error {
	return nil
}
func (m *customBlockedModel) ModifyBlockByBlockIPWithCountAndBlockend(ctx context.Context, blockIp string, endTime time.Time, count int64) error {
	return nil
}

func (m *customBlockedModel) DeleteBlockByBlockIP(ctx context.Context, blockIp string) (*mongo.UpdateResult, error) {
	updateAt := time.Now()
	res, err := m.conn.UpdateOneNoCache(ctx, bson.M{"blockIp": blockIp}, bson.M{"blockEnd": 0, "updateAt": updateAt})
	return res, err
}
