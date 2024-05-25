package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/logic"
	"github.com/lius-new/blog-backend/rpc/analyzer/tests"
)

func TestSelectRecordById(t *testing.T) {
	ctx := context.Background()

	selectRecordByIdResp, err := logic.NewSelectRecordByIdLogic(ctx, tests.SVC_CONTEXT).SelectRecordById(&analyzer.SelectRecordByIdRequest{
		Id: "6651dd67628f8c4c59f2138e",
	})

	if err != nil {
		fmt.Println("error :", err)
	} else {
		fmt.Println(selectRecordByIdResp)
	}
}
