package mongo_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	model "github.com/lius-new/blog-backend/rpc/analyzer/model/mongo/record"
	"github.com/lius-new/blog-backend/rpc/analyzer/tests"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TestInsertWithRecord: 测试插入record
func TestInsertWithRecord(t *testing.T) {
	ctx := context.Background()

	for i := 0; i < 100; i++ {
		err := tests.SVC_CONTEXT.ModelWithRecord.Insert(ctx, &model.Record{
			RequestIP:     "127.0.0.1",
			RequestMethod: "GET",
			RequestPath:   "/",
			RequestCount:  1,
		})

		if err != nil {
			fmt.Println("err: ", err)
		} else {
			fmt.Println("result: 插入成功")
		}
	}
}

// TestDeleteWithRecord: 测试删除record
func TestDeleteWithRecord(t *testing.T) {
	ctx := context.Background()
	count, err := tests.SVC_CONTEXT.ModelWithRecord.Delete(ctx, "6650023bd89f7bf9770ef8e6")
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("result: 删除成功: count( ", count, " )")
	}
}

// TestModifyWithRecord: 测试更新record
func TestModifyWithRecord(t *testing.T) {
	ctx := context.Background()
	id, _ := primitive.ObjectIDFromHex("66500282e47fa995e4a5aaa2")
	modifyResult, err := tests.SVC_CONTEXT.ModelWithRecord.Update(ctx, &model.Record{
		ID:            id,
		RequestIP:     "127.0.0.1:8080",
		RequestMethod: "GET",
		RequestPath:   "/",
		RequestCount:  1,
	})
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("result: 更新成功: modify( ", modifyResult, " )")
	}
}

// TestFindByIdWithRecord: 测试查询record(根据id)
func TestFindByIdWithRecord(t *testing.T) {
	ctx := context.Background()
	record, err := tests.SVC_CONTEXT.ModelWithRecord.FindOne(ctx, "66500282e47fa995e4a5aaa2")
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("result: 查询成功:  ", record)
	}
}

// TestFindByPageWithRecord: 测试查询record(分页)
func TestFindByPageWithRecord(t *testing.T) {
	ctx := context.Background()
	records, total, err := tests.SVC_CONTEXT.ModelWithRecord.FindByPage(ctx, 1, 3)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("result: 查询成功:  ", records, " 总计: ", total)
	}
}

// TestCountDayRecordNumber: 统计一天内访问次数
func TestCountDayRecordNumber(t *testing.T) {
	ctx := context.Background()
	total, err := tests.SVC_CONTEXT.ModelWithRecord.CountScopeTimeRecordNumber(ctx, "127.0.0.1:8080", time.Hour)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("result: 今日访问次数:", total)
	}
}
