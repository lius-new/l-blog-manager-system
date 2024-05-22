package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
)

type DeleteArtilceByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteArtilceByIdLogic(
	ctx context.Context,
	svcCtx *svc.ServiceContext,
) *DeleteArtilceByIdLogic {
	return &DeleteArtilceByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据删除文章
func (l *DeleteArtilceByIdLogic) DeleteArtilceById(
	in *content.DeleteArticleRequest,
) (*content.DeleteArticleResponse, error) {
	findArticle, err := l.svcCtx.ModelWithArticle.FindOne(l.ctx, in.Id)

	if err == rpc.ErrNotFound {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	// 根据id删除指定文章
	count, err := l.svcCtx.ModelWithArticle.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	// 删除article对应的covers
	deleteCoverLogic := NewDeleteCoverLogic(l.ctx, l.svcCtx)
	for _, coverId := range findArticle.Covers {
		deleteCoverLogic.DeleteCover(&content.DeleteCoverRequest{
			Id: coverId,
		})
	}
	// 从tag中移除article
	modifyTagRemoveArticleLogic := NewModifyTagRemoveArticleLogic(l.ctx, l.svcCtx)
	for _, v := range findArticle.Tags {
		modifyTagRemoveArticleLogic.ModifyTagRemoveArticle(&content.ModifyTagRemoveArticleRequest{
			Id:      v,
			Article: findArticle.ID.Hex(),
		})
	}

	return &content.DeleteArticleResponse{Count: count}, nil
}
