package logic_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/lius-new/blog-backend/rpc/user/internal/logic"
	"github.com/lius-new/blog-backend/rpc/user/tests"
	"github.com/lius-new/blog-backend/rpc/user/user"
)

func TestSelect(t *testing.T) {
	selectLogic := logic.NewSelectLogic(context.Background(), tests.SVC_CONTEXT)

	uid := "sdf"
	// uid := "661c997d54508460ada11b84"
	resp, err := selectLogic.Select(&user.SelectUserRequest{Uid: uid})

	if err != nil {
		fmt.Println(">>>>>>>>>", err)
	}
	log.Println(resp)
}
