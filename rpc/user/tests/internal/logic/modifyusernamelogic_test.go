package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/user/internal/logic"
	"github.com/lius-new/blog-backend/rpc/user/tests"
	"github.com/lius-new/blog-backend/rpc/user/user"
)

func TestModifyUserName(t *testing.T) {
	logicClient := logic.NewModifyUserNameLogic(context.Background(), tests.SVC_CONTEXT)

	resp, err := logicClient.ModifyUserName(&user.ModifyUserNameRequest{
		Id:       "663b9e602a23854f42fc6df8",
		Username: "lius",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
