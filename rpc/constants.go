package rpc

import (
	"errors"

	"github.com/zeromicro/go-zero/core/stores/mon"
)

// 数据库相关的的错误常量
var (
	ErrNotFound        = mon.ErrNotFound
	ErrInvalidDisabled = errors.New("invalid disabled")
	ErrInvalidExist    = errors.New("invalid exist")
	ErrInvalidObjectId = errors.New("invalid objectId")
	ErrInvalidDeleted  = errors.New("invalid delete")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidToken    = errors.New("invalid token")
)

// 请求相关的错误常量
var (
	ErrRequestParam = errors.New("param validate failed")
)
