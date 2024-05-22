package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/logic"
	"github.com/lius-new/blog-backend/rpc/content/tests"
)

func TestModifyTagVisiable(t *testing.T) {
	ctx := context.Background()

	modifyTagVisiableLogic := logic.NewModifyTagVisiableLogic(ctx, tests.SVC_CONTEXT)

	resp, err := modifyTagVisiableLogic.ModifyTagVisiable(&content.ModifyTagVisiableRequest{
		Id:       "664dad91b821007e111068d2",
		Visiable: false,
	})

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(resp)
	}

}
