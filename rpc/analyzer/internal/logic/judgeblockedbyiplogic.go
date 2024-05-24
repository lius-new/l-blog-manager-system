package logic

import (
	"context"
	"time"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeBlockedByIPLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJudgeBlockedByIPLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeBlockedByIPLogic {
	return &JudgeBlockedByIPLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 判断是否被封禁
func (l *JudgeBlockedByIPLogic) JudgeBlockedByIP(in *analyzer.JudgeBlockedByIPRequest) (*analyzer.JudgeBlockedByIPResponse, error) {
	if len(in.BlockIP) == 0 {
		return nil, rpc.ErrRequestParam
	}

	// 查询到指定IP
	blocked, err := l.svcCtx.ModelWithBlocked.FindByBlockIP(l.ctx, in.BlockIP)

	// 如果查询出现错误且是因为不存在的错误那么不抛出错误并正常返回结果(意味没有被封禁, 应该返回false)
	if err != nil && err == rpc.ErrNotFound {
		return &analyzer.JudgeBlockedByIPResponse{Block: false}, nil
	}

	// 如果不是因为不存在的错误那么就直接返回错误
	if err != nil {
		return nil, err
	}

	// 如果查询成功, 那么意味的确有过封禁或者封禁历史
	// 判断当前时间和封禁截止时间, 条件成立即当前时间大于封禁时间意味封禁结束了。
	if time.Now().UTC().Unix() > blocked.BlockEnd.Unix() {
		return &analyzer.JudgeBlockedByIPResponse{Block: false}, nil
	} else {
		// 如果当前时间小于封禁时间那么意味仍然处于封禁状态。
		return &analyzer.JudgeBlockedByIPResponse{Block: true}, nil
	}
}
