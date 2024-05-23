package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Record struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	RequestIP     string `bson:"requestIP,omitempty"        json:"requestIP,omitempty"`     // 请求者的IP
	RequestMethod string `bson:"requestMethod,omitempty"    json:"requestMethod,omitempty"` // 请求者请求资源的请求方法
	RequestPath   string `bson:"requestPath,omitempty"      json:"requestPath,omitempty"`   // 请求者请求资源的请求地址
	RequestCount  int64  `bson:"requestCount,omitempty"     json:"requestCount,omitempty"`  // 请求者请求资源的次数，设置默认为1. 当在指定时间段达到阈值的时候阈值数量合并并修改count.
	// RequestFromLevel int64  `bson:"requestFromLevel,omitempty" json:"requestFromLevel,omitempty"` // 请求者的级别, 比如服务器后台访问级别高，游客浏览器访问级别低，博客所有者也有不同的访问级别。

	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
