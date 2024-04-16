package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/authorization/authorization"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/logic"
	"github.com/lius-new/blog-backend/rpc/authorization/tests"
)

func TestSelectSecret(t *testing.T) {
	ctx := context.Background()

	logicClient := logic.NewGenerateSecretLogic(ctx, tests.SVC_CONTEXT)

	resp, err := logicClient.GenerateSecret(&authorization.GenerateSecretRequestWithSecret{})

	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
