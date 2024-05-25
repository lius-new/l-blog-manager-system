package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/logic"
	"github.com/lius-new/blog-backend/rpc/analyzer/tests"
)

func TestCreateRecord(t *testing.T) {
	ctx := context.Background()
	createRecordResp, err := logic.NewCreateRecordLogic(ctx, tests.SVC_CONTEXT).CreateRecord(&analyzer.CreateRecordRequest{
		RequestIp:     "localhost",
		RequestMethod: "-1",
		RequestPath:   "0",
	})

	if err != nil {
		fmt.Println("error :", err)
	} else {
		fmt.Println(createRecordResp)
	}
}
