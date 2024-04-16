package logic_test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/authorization/authorization"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/logic"
	"github.com/lius-new/blog-backend/rpc/authorization/tests"
)

func TestGenerateSecret(t *testing.T) {
	ctx := context.Background()

	logicClient := logic.NewGenerateSecretLogic(ctx, tests.SVC_CONTEXT)
	resp, err := logicClient.GenerateSecret(&authorization.GenerateSecretRequestWithSecret{
		Uid:    "661c997d54508460adab84",
		Issuer: "shfj",
		Expire: int64(time.Hour),
	})

	if err != nil {
		switch {
		case strings.Contains(err.Error(), rpc.ErrInvalidObjectId.Error()):
			fmt.Println(">>>> :", "错误id")
		default:
			fmt.Println("err :", err)
		}
	}
	fmt.Println(resp)
}
