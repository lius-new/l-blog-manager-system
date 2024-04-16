package model

import (
	"errors"

	"github.com/zeromicro/go-zero/core/stores/mon"
)

var (
	ErrNotFound        = mon.ErrNotFound
	ErrExist           = errors.New("invalid exist")
	ErrInvalidObjectId = errors.New("invalid objectId")
)
