package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Secret struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	SecretInner string             `bson:"secretInner,omitempty" json:"secretInner,omitempty"`
	SecretOuter string             `bson:"secretOuter,omitempty" json:"secretOuter,omitempty"`
	Expire      int64              `bson:"expire,omitempty" json:"expire,omitempty"`
	Issuer      string             `bson:"issuer,omitempty" json:"issuer,omitempty"`
	UserId      string             `bson:"userId,omitempty" json:"userId,omitempty"`
	UpdateAt    time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt    time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
