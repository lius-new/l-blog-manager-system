package logic

import (
	"context"
	"time"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/analyzer/model/mongo/blocked"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBlockedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBlockedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBlockedLogic {
	return &CreateBlockedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ================ Blocked  ================
func (l *CreateBlockedLogic) CreateBlocked(in *analyzer.CreateBlockedRequest) (*analyzer.CreateBlockedResponse, error) {
	// 查询指定ip是否存在
	selectBlockedByBlockIPLogic := NewSelectBlockedByBlockIPLogic(l.ctx, l.svcCtx)
	selectResp, err := selectBlockedByBlockIPLogic.SelectBlockedByBlockIP(&analyzer.SelectBlockedByBlockIPRequest{
		BlockIp: in.BlockIp,
	})

	if err != nil && err != rpc.ErrNotFound {
		return nil, err
	}

	// 设置封禁结束时间
	setting, err := l.svcCtx.ModelWithSetting.FindLastSetting(l.ctx)
	if err != nil && err != rpc.ErrNotFound {
		return nil, err
	}

	// 如果没有添加设置，那么就使用默认封禁时间
	if err == rpc.ErrNotFound {
		setting.RecordMergeInterval = time.Hour * 12
	}
	endTime := time.Now().Add(setting.RecordMergeInterval)

	// 判断是否是notfound, 那么就是创建新的
	if err == rpc.ErrNotFound {
		err = l.svcCtx.ModelWithBlocked.Insert(l.ctx, &model.Blocked{
			BlockIP:    in.BlockIp,
			BlockCount: 1,
			BlockEnd:   endTime,
		})
	} else {
		// 也就是说存在过封禁历史, 此时修改历史
		err = l.svcCtx.ModelWithBlocked.ModifyBlockByBlockIPWithCountAndBlockend(l.ctx, in.BlockIp, endTime, selectResp.BlockCount+1)
	}

	if err != nil {
		return nil, err
	}

	return &analyzer.CreateBlockedResponse{}, nil
}
