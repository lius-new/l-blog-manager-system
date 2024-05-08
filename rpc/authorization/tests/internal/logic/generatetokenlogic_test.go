package logic_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lius-new/blog-backend/rpc/authorization/authorization"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/logic"
	"github.com/lius-new/blog-backend/rpc/authorization/tests"
)

func TestGenerateToken(t *testing.T) {
	ctx := context.Background()

	logicClient := logic.NewGenerateTokenLogic(ctx, tests.SVC_CONTEXT)

	resp, err := logicClient.GenerateToken(&authorization.GenerateJwtRequestWithJwt{
		Id:       "663ae6880cf2ed35a8e362ea",
		Uid:      "663ae6779d7c676bf233dfc6",
		Uesrname: "胡娜",
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Token)
}

type Claims struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	jwt.RegisteredClaims
}

// 测试下生成TOKEN, 因为之前随机的secret可以在浏览器在线解析token解析，我以为是secret的问题，所以下面测试方法写了个死的secret
func TestGenerateTokenWithJwt(t *testing.T) {
	claims := Claims{
		Key:   "lius",
		Value: "value",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 8)),
			Issuer:    "issuer",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte("*secret."))
	if err != nil {
		panic(err)
	}

	fmt.Println(tokenStr)
}
