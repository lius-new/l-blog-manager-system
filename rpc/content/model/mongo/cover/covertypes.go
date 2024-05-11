package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cover struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"      json:"id,omitempty"`
	Content  string             `bson:"content,omitempty"  json:"content,omitempty"`
	Hash     string             `bson:"hash,omitempty"     json:"hash,omitempty"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
