package logic_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/logic"
	"github.com/lius-new/blog-backend/rpc/analyzer/tests"
)

func TestSelectBlockedById(t *testing.T) {

	ctx := context.Background()

	selectBlockedByIdResp, err := logic.NewSelectBlockedByIdLogic(ctx, tests.SVC_CONTEXT).SelectBlockedById(
		&analyzer.SelectBlockedByIdRequest{
			Id: "665012f940afee8fcbbdba19",
		},
	)

	if err != nil {
		fmt.Println("error :", err)
	} else {
		fmt.Println(selectBlockedByIdResp)
	}
}
