package logic_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/user/internal/logic"
	"github.com/lius-new/blog-backend/rpc/user/tests"
	"github.com/lius-new/blog-backend/rpc/user/user"
)

func TestLogin(t *testing.T) {
	loginLogic := logic.NewLoginLogic(context.Background(), tests.SVC_CONTEXT)
	resp, err := loginLogic.Login(&user.LoginUserRequest{
		Username: "lius6666",
		Password: "14569636547",
	})
	if err != nil {
		if err == rpc.ErrNotFound {
			fmt.Println("lius: ", err)
		} else {
			panic(err)
		}
	}
	log.Println(resp)
}
