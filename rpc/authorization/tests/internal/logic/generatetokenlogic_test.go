package logic_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/lius-new/blog-backend/rpc/authorization/authorization"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/logic"
	"github.com/lius-new/blog-backend/rpc/authorization/tests"
)

func TestGenerateToken(t *testing.T) {
	ctx := context.Background()

	logicClient := logic.NewGenerateSecretLogic(ctx, tests.SVC_CONTEXT)

	resp, err := logicClient.GenerateSecret(&authorization.GenerateSecretRequestWithSecret{
		Expire: int64(time.Hour),
		Issuer: "lius",
		Uid:    "661c997d54508460ada11b84",
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
