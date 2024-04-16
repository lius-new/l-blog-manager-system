package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/authorization/authorization"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/logic"
	"github.com/lius-new/blog-backend/rpc/authorization/tests"
)

func TestDeleteSecret(t *testing.T) {
	ctx := context.Background()

	logicClient := logic.NewDeleteSecretLogic(ctx, tests.SVC_CONTEXT)

	resp, err := logicClient.DeleteSecret(&authorization.DeleteSecretRequestWithSecret{
		Id: "661cf647f1b85bf75a0be728",
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
