package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/user/internal/logic"
	"github.com/lius-new/blog-backend/rpc/user/tests"
	"github.com/lius-new/blog-backend/rpc/user/user"
)

func TestModifyPassword(t *testing.T) {
	logicClient := logic.NewModifyPasswordLogic(context.Background(), tests.SVC_CONTEXT)

	resp, err := logicClient.ModifyPassword(&user.ModifyPasswordRequest{
		Id:       "663b9e602a23854f42fc6df8",
		Password: "12342",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
