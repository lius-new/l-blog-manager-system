package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blocked struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	BlockIP    string    `bson:"blockIp,omitempty"  json:"blockIp,omitempty"`      // 禁用的IP
	BlockEnd   time.Time `bson:"blockEnd,omitempty"   json:"blockEnd,omitempty"`   // 禁用的截止时间
	BlockCount int64     `bson:"blockCount,omitempty" json:"blockCount,omitempty"` // 禁用的次数

	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
