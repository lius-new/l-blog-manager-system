package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ TagModel = (*customTagModel)(nil)

type (
	// TagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTagModel.
	TagModel interface {
		tagModel
		InsertReturnId(ctx context.Context, data *Tag) (id string, err error)
		FindByName(ctx context.Context, name string) (*Tag, error)
		FindByPage(
			ctx context.Context,
			pageNum, pageSize int64,
			hideShow bool,
		) ([]Tag, int64, error)
		UpdateVisiable(ctx context.Context, data *Tag) (*mongo.UpdateResult, error)
	}
	customTagModel struct {
		*defaultTagModel
	}
)

// NewTagModel returns a model for the mongo.
func NewTagModel(url, db, collection string, c cache.CacheConf) TagModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customTagModel{
		defaultTagModel: newDefaultTagModel(conn),
	}
}

// InsertWithId: 插入tag且一定要包含id
func (m *customTagModel) InsertReturnId(ctx context.Context, data *Tag) (id string, err error) {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	key := prefixTagCacheKey + data.ID.Hex()
	_, err = m.conn.InsertOne(ctx, key, data)
	return data.ID.Hex(), err
}

// FindByName: 根据tag name来查询
func (m *customTagModel) FindByName(ctx context.Context, name string) (*Tag, error) {
	var tag Tag
	err := m.conn.FindOneNoCache(ctx, &tag, bson.M{"name": name})

	switch err {
	case nil:
		return &tag, nil
	case monc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindByPage: 查看文章, 分页
// hideShow : 是否包含显示被设置为visiable = false的标签
// 写入到缓存中
func (m *customTagModel) FindByPage(
	ctx context.Context,
	pageNum, pageSize int64,
	hideShow bool,
) ([]Tag, int64, error) {
	tags := make([]Tag, pageNum)

	findOptions := options.Find()
	// 如果传入参数小于等于0那么就设置为1
	if pageNum <= 0 {
		pageNum = 1
	}

	findOptions.SetLimit(pageSize)
	findOptions.SetSkip(pageSize * (pageNum - 1))
	findOptions.SetSort(bson.M{"updateAt": -1}) // 根据时间降序排序

	// 查询
	var err error
	if hideShow {
		err = m.conn.Find(ctx, &tags, bson.M{}, findOptions)
	} else {
		err = m.conn.Find(ctx, &tags, bson.M{"visiable": true}, findOptions)
	}

	// TODO: 应该将查询结果缓存起来, 但是因为如果缓存我就要去修改更新和删除的代码.

	switch err {
	case nil:
		total, _ := m.conn.CountDocuments(ctx, bson.M{})
		return tags, total, nil
	case monc.ErrNotFound:
		return nil, 0, ErrNotFound
	default:
		return nil, 0, err
	}
}

func (m *customTagModel) UpdateVisiable(ctx context.Context, data *Tag) (*mongo.UpdateResult, error) {

	data.UpdateAt = time.Now()
	key := prefixTagCacheKey + data.ID.Hex()

	res, err := m.conn.UpdateOne(
		ctx,
		key,
		bson.M{"_id": data.ID},
		bson.M{"$set": bson.M{"visiable": data.Visiable}},
	)

	return res, err
}
