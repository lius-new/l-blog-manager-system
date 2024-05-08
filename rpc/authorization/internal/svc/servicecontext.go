package svc

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/lius-new/blog-backend/rpc/authorization/internal/config"
	model "github.com/lius-new/blog-backend/rpc/authorization/model/mongo"
	"github.com/lius-new/blog-backend/rpc/user/userer"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Model  model.SecretModel
	Userer userer.Userer
	Redis  *redis.Redis
}

// 随机生成token secret
func GenerateRandomKey(length int) (string, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}

// 建立redis连接
func connRedis(c cache.CacheConf) *redis.Redis {
	rds := redis.MustNewRedis(redis.RedisConf{
		Host:        c[0].Host,
		Type:        "node",
		Pass:        "",
		Tls:         false,
		NonBlock:    false,
		PingTimeout: time.Second,
	})
	return rds
}

// 刷新redis中的outer secret
func refreshSecretOuter(rds *redis.Redis) string {
	var secretCacheKey = "cache:secret:outer"

	// TODO: 实施上这里好像不用设置超时。 如果超时那么就意味REDIS已经挂了。
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancel()

	secretOuter, err := rds.GetCtx(ctx, secretCacheKey)

	// 查询异常或者secretOuter为空那么就刷新redis中数据
	if err != nil || secretOuter == "" {
		rds.DelCtx(ctx, secretCacheKey)
		secret, _ := GenerateRandomKey(32)
		rds.SetCtx(ctx, secretCacheKey, secret)
		return secret
	} else {
		return secretOuter
	}
}

// 获取redis中的outer secret
func (s ServiceContext) GetSecretOuter() string {
	return refreshSecretOuter(s.Redis)
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 在redis中添加一份secret outer, 用于校验token时候的外层密钥
	rds := connRedis(c.Cache)
	defer refreshSecretOuter(rds)
	return &ServiceContext{
		Config: c,
		Model:  model.NewSecretModel(c.MongoURL, c.DBName, "secret", c.Cache),
		Userer: userer.NewUserer(zrpc.MustNewClient(c.User)),
		Redis:  rds,
	}
}
