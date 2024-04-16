package logic_test

import (
	"context"
	"log"
	"testing"

	"github.com/lius-new/blog-backend/rpc/utils/internal/logic"
	"github.com/lius-new/blog-backend/rpc/utils/tests"
	"github.com/lius-new/blog-backend/rpc/utils/utils"
)

func TestLogin(t *testing.T) {
	mD5Logic := logic.NewMD5Logic(context.Background(), tests.SVC_CONTEXT)
	resp, err := mD5Logic.MD5(&utils.MD5Reqeust{
		Text: "abc",
	})
	if err != nil {
		panic(err)
	}
	log.Println(resp)
}
