package logic

import (
	"context"
	"errors"
	"strings"

	"github.com/lius-new/blog-backend/api"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"
)

type ViewImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewImageLogic {
	return &ViewImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewImageLogic) ViewImage(req *types.ViewImageRequest) (resp *types.ViewImageResponse, err error) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			catchErr := catchErr.(error)
			switch {
			case strings.Contains(catchErr.Error(), api.ErrNotFound.Error()):
				err = errors.New(api.ErrInvalidNotFound.Error())
			}
		} else if err != nil {
			err = errors.New(strings.Replace(err.Error(), "rpc error: ", "", 1))
		}
	}()

	coverResp, err := l.svcCtx.Content.SelectCoverByHash(l.ctx, &content.SelectCoverByHashRequest{
		Hash: req.Hash,
	})

	if err != nil {
		panic(err)
	}

	return &types.ViewImageResponse{
		Base64: coverResp.Cover.Content,
	}, nil
}
