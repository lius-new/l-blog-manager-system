package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

var _ SecretModel = (*customSecretModel)(nil)

type (
	// SecretModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSecretModel.
	SecretModel interface {
		secretModel
	}

	customSecretModel struct {
		*defaultSecretModel
	}
)

// NewSecretModel returns a model for the mongo.
func NewSecretModel(url, db, collection string, c cache.CacheConf) SecretModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customSecretModel{
		defaultSecretModel: newDefaultSecretModel(conn),
	}
}
