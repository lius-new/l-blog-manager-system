package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT *jwtStruct = &jwtStruct{}

const TokenExpireDuration = time.Hour * 12

type Claims struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Token    string `json:"token"`
	jwt.RegisteredClaims
}
type jwtStruct struct{}

// 获取过期时间
// 因为数据库也存, 所以通过service获取后再传递过来给GenerateJwtToken方法
func (j jwtStruct) GetExpiresAt() time.Time {
	return time.Now().Add(TokenExpireDuration)
}

// 生成token
func (j jwtStruct) GenerateJwtToken(
	userId, username, secret, issuer string, expiresAt time.Time,
) (string, error) {
	claims := Claims{
		UserID:   userId,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			Issuer:    issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func (j jwtStruct) ParseJwtToken(secret, tokenStr string) (Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		err = errors.New("invalid Token")
	}

	return *claims, err
}

func (j jwtStruct) GenerateJwtTokenSecond(
	secret, issuer string, expiresAt time.Time,
	username, tokenString string,
) (string, error) {
	claims := Claims{
		Username: username,
		Token:    tokenString,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			Issuer:    issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// TODO: salt
	return token.SignedString([]byte(secret))
}
