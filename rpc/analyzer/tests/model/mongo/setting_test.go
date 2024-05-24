package mongo_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	model "github.com/lius-new/blog-backend/rpc/analyzer/model/mongo/setting"
	"github.com/lius-new/blog-backend/rpc/analyzer/tests"
)

func TestFindLastSetting(t *testing.T) {
	ctx := context.Background()

	setting, err := tests.SVC_CONTEXT.ModelWithSetting.FindLastSetting(ctx)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println(setting)
	}
}

func TestInsertSetting(t *testing.T) {
	ctx := context.Background()

	err := tests.SVC_CONTEXT.ModelWithSetting.Insert(ctx, &model.Setting{
		RecordMergeBoundary: 20,
		RecordMergeInterval: time.Hour * 24,
	})

	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("添加成功！")
	}
}
