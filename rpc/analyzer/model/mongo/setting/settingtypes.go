package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Setting struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	// 记录合并边界
	RecordMergeBoundary int64 `bson:"recordMergeBoundary,omitempty" json:"recordMergeBoundary,omitempty"`
	// 记录合并间隔(指定时间间隔内合并)
	RecordMergeInterval time.Duration `bson:"recordMergeInterval,omitempty" json:"recordMergeInterval,omitempty"`

	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
