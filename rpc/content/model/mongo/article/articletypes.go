package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title    string             `bson:"title,omitempty" json:"title,omitempty"`
	Desc     string             `bson:"desc,omtempty" json:"desc,omitempty"`
	Content  string             `bson:"content,omtempty" json:"content,omitempty"`
	Tags     []string           `bson:"tags,omtempty" json:"tags,omitempty"`
	Covers   []string           `bson:"covers,omtempty" json:"covers,omitempty"`
	Visiable bool               `bson:"visiable,omtempty" json:"visiable,omitempty"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
