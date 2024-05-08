package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/authorization/authorization"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/logic"
	"github.com/lius-new/blog-backend/rpc/authorization/tests"
)

func TestJudgeToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJrZXkiOiLog6HlqJwiLCJ2YWx1ZSI6ImV5SmhiR2NpT2lKSVV6TTROQ0lzSW5SNWNDSTZJa3BYVkNKOS5leUpyWlhraU9pSTJOak5oWlRZM056bGtOMk0yTnpaaVpqSXpNMlJtWXpZaUxDSjJZV3gxWlNJNkl1aURvZVdvbkNJc0ltbHpjeUk2SW14dloybHVJaXdpWlhod0lqb3pORE13TXpBMU9UZzJMQ0pwWVhRaU9qRTNNVFV4TkRneU5UQjkuNGhvU1JkU1d4cFhRNE1SQXFfbHAwRHFURl9vOXpRZ0dCSWREX0hZYkFWeUpyUFNZOTg3bk9Ea2piYWYxMERrMyIsImlzcyI6ImxvZ2luIiwiZXhwIjozNDMwMzA1OTg2LCJpYXQiOjE3MTUxNDgyNTB9.9sEQSH6iVSJbT6_cSWJULeja9P_bx5-qAJ2DKYCuwlU"

	ctx := context.Background()

	logicClient := logic.NewJudgeTokenLogic(ctx, tests.SVC_CONTEXT)

	resp, err := logicClient.JudgeToken(&authorization.JudgeJwtRequestWithJwt{
		Token: token,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
