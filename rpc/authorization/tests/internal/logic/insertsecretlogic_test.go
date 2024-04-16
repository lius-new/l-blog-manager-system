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

func TestInsertSecrett(t *testing.T) {
	ctx := context.Background()

	logicClient := logic.NewGenerateSecretLogic(ctx, tests.SVC_CONTEXT)

	resp, err := logicClient.GenerateSecret(&authorization.GenerateSecretRequestWithSecret{
		Uid:    "661c997d54508460ada11b843",
		Issuer: "lius",
		Expire: int64(time.Hour),
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
