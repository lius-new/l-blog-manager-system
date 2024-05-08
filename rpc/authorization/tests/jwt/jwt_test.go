package jwt_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/lius-new/blog-backend/rpc/authorization/internal/jwt"
)

func TestParseJwtToken(t *testing.T) {
	// jwtUtil := jwt.NewJwtUtil("inner", "outer", time.Now().Add(time.Hour), "lius")
	// token, err := jwtUtil.GenerateJwtToken("uid:123456", "username:user-lius")
	// if err != nil {
	// 	panic(err)
	// }
	// claims, err := jwtUtil.ParseJwtTokenOuter()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(claims)
}

func TestGenerateJwtToken(t *testing.T) {
	jwtUtil := jwt.NewJwtUtil("inner", "outer", time.Now().Add(time.Duration(3600)), "lius")
	token, err := jwtUtil.GenerateJwtToken("uid:123456", "username:user-lius")
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
