package logic_test

import (
	"fmt"
	"testing"

	"context"

	"github.com/lius-new/blog-backend/rpc/utils/internal/logic"
	"github.com/lius-new/blog-backend/rpc/utils/tests"
	"github.com/lius-new/blog-backend/rpc/utils/utils"
)

func TestHashWithString(t *testing.T) {
	hashLogic := logic.NewHashWithStringLogic(context.Background(), tests.SVC_CONTEXT)
	resp, err := hashLogic.HashWithString(&utils.HashWithStringReqeust{
		Content: "",
	})
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(resp)
}
