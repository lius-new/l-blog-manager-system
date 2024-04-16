package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/user/internal/logic"
	"github.com/lius-new/blog-backend/rpc/user/tests"
	"github.com/lius-new/blog-backend/rpc/user/user"
)

func TestModify(t *testing.T) {

	logicClient := logic.NewModifyLogic(context.Background(), tests.SVC_CONTEXT)
	resp, err := logicClient.Modify(&user.ModifyUserRequest{
		Id:       "661e450ec428a0e0abce3e4c",
		Username: "lius6666",
		Password: "57df302a4e041396b534d27ec64eb0ac",
		Status:   false,
		SecretId: "661e33b895fb0ef92c541012",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
