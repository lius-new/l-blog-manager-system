package api

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
	ErrInvalidNotFound = errors.New("not found current resource")
)

// 请求相关的错误常量
var (
	ErrRequestParam = errors.New("param validate failed")
	ErrHostNotFound = errors.New("request.host failed")
)

// 权限校验的错误常量
var (
	ErrInvalidToken        = errors.New("invalid token")
	ErrFailedAuthorization = errors.New("authorization failed")
)

// 运行时的错误
var (
	ErrInvalidMemoryNil = errors.New("nil pointer dereference")
)
