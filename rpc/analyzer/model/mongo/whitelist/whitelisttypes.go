package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Whitelist struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	Ip   string `bson:"ip,omitempty"   json:"ip,omitempty"`   // 白名单中的ip, 该ip不会被记录
	Path string `bson:"path,omitempty" json:"path,omitempty"` // 白名单中的地址, 请求该地址不会被记录

	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
