package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/user/internal/logic"
	"github.com/lius-new/blog-backend/rpc/user/tests"
	"github.com/lius-new/blog-backend/rpc/user/user"
)

func TestModifySecretId(t *testing.T) {
	logicClient := logic.NewModifySecretIdLogic(context.Background(), tests.SVC_CONTEXT)

	resp, err := logicClient.ModifySecretId(&user.ModifySecretRequest{
		Id:       "663b9e602a23854f42fc6df8",
		SecretId: "663b9e7a3107cbf02b0b113c",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
