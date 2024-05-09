package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/user/internal/logic"
	"github.com/lius-new/blog-backend/rpc/user/tests"
	"github.com/lius-new/blog-backend/rpc/user/user"
)

func TestModifyStatus(t *testing.T) {
	logicClient := logic.NewModifyStatusLogic(context.Background(), tests.SVC_CONTEXT)

	resp, err := logicClient.ModifyStatus(&user.ModifyUserStatusRequest{
		Id:     "663b9e602a23854f42fc6df8",
		Status: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
