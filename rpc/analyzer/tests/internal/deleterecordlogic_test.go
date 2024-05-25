package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/logic"
	"github.com/lius-new/blog-backend/rpc/analyzer/tests"
)

func TestDeleteRecordById(t *testing.T) {

	ctx := context.Background()
	deleteRecordResp, err := logic.NewDeleteRecordByIdLogic(ctx, tests.SVC_CONTEXT).DeleteRecordById(&analyzer.DeleteRecordByIdRequest{
		Id: "6651da5b200e3d586555a4ef",
	})

	if err != nil {
		fmt.Println("error :", err)
	} else {
		fmt.Println(deleteRecordResp)
	}
}
