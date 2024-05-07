package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/lius-new/blog-backend/rpc"
)

type JWTGenerate interface {
	ParseJwtToken(secret, token string) (JWTGenerate, error)
	GenerateJwtToken(key, value string) (string, error)
}

type Claims struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	jwt.RegisteredClaims
}

type JwtUtil struct {
	SecretInner string
	SecretOuter string
	Expire      time.Time
	Issuer      string
}

func NewJwtUtil(secret1, secret2 string, expire time.Time, issuer string) JwtUtil {
	return JwtUtil{
		SecretInner: secret1,
		SecretOuter: secret2,
		Expire:      expire,
		Issuer:      issuer,
	}
}

// buildGenerateJwtTokenFunc: 生成GenerateJwtTokenFunc的工厂方法, 支持修改signed method的类型
func (j JwtUtil) buildGenerateJwtTokenFunc(
	signedMethod *jwt.SigningMethodHMAC,
	secret string,
) func(key, value string) (string, error) {
	return func(key, value string) (string, error) {
		if len(key) == 0 || len(value) == 0 {
			return "", errors.New("claim param empty")
		}
		claims := Claims{
			Key:   key,
			Value: value,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(j.Expire),
				Issuer:    j.Issuer,
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		}

		token := jwt.NewWithClaims(signedMethod, claims)
		return token.SignedString([]byte(secret))
	}
}

// generateJwtTokenInner: 生成内部token
func (j JwtUtil) generateJwtTokenInner(key, value string) (string, error) {
	return j.buildGenerateJwtTokenFunc(jwt.SigningMethodHS384, j.SecretInner)(key, value)
}

// generateJwtTokenOuter: 生成外部token
func (j JwtUtil) generateJwtTokenOuter(key, value string) (string, error) {
	return j.buildGenerateJwtTokenFunc(jwt.SigningMethodHS256, j.SecretOuter)(key, value)
}

// ParseJwtToken: 解析token
func (j JwtUtil) ParseJwtToken(token string) (*Claims, error) {
	parse := func(secret, token string, c *Claims) (*jwt.Token, error) {
		res, err := jwt.ParseWithClaims(token, c, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !res.Valid {
			return res, rpc.ErrInvalidToken
		}

		return res, err
	}

	// 解析外层
	outerClaims := Claims{}
	_, err := parse(j.SecretOuter, token, &outerClaims)
	if err != nil {
		return nil, err
	}

	// 解析内层
	innerClaims := Claims{}
	_, err = parse(j.SecretInner, outerClaims.Value, &innerClaims)
	if err != nil {
		return nil, err
	}

	return &innerClaims, err
}

// GenerateJwtToken: 解析token
func (j JwtUtil) GenerateJwtToken(key, value string) (string, error) {
	innerToken, err := j.generateJwtTokenInner(key, value)
	if err != nil {
		return "", err
	}
	outerToken, err := j.generateJwtTokenOuter(value, innerToken)
	if err != nil {
		return "", err
	}

	return outerToken, nil
}
