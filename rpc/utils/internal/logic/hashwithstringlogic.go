package logic

import (
	"context"
	"fmt"
	"hash/fnv"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/utils/internal/svc"
	"github.com/lius-new/blog-backend/rpc/utils/utils"
)

type HashWithStringLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHashWithStringLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HashWithStringLogic {
	return &HashWithStringLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HashWithStringLogic) HashWithString(
	in *utils.HashWithStringReqeust,
) (*utils.HashWithStringReponse, error) {
	if len(in.Content) == 0 {
		return nil, rpc.ErrRequestParam
	}
  // 创建FNV-1a哈希算法实例
	hasher := fnv.New32a()
  // 将字符串写入实例
	hasher.Write([]byte(in.Content))
  // 计算hash
	hash := hasher.Sum32()
	return &utils.HashWithStringReponse{
		Content: fmt.Sprintf("%x", hash),
	}, nil
}
