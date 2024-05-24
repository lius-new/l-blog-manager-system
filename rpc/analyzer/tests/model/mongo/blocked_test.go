package mongo_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	model "github.com/lius-new/blog-backend/rpc/analyzer/model/mongo/blocked"
	"github.com/lius-new/blog-backend/rpc/analyzer/tests"
)

// TestFindByBlockIP: 测试根据IP查找block
func TestFindByBlockIP(t *testing.T) {
	ctx := context.Background()
	blocked, err := tests.SVC_CONTEXT.ModelWithBlocked.FindByBlockIP(ctx, "127.0.0.1")
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("result: 查询成功: ", blocked)
	}
}

// TestFindByPage: 测试分页
func TestFindByPage(t *testing.T) {
	ctx := context.Background()
	blockeds, total, err := tests.SVC_CONTEXT.ModelWithBlocked.FindByPage(ctx, 1, 2)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("result: 查询成功:  ", blockeds, " 总计: ", total)
	}
}

// TestModifyBlockByBlockIPWithCount: 测试根据IP修改count
func TestModifyBlockByBlockIPWithCount(t *testing.T) {
	ctx := context.Background()
	err := tests.SVC_CONTEXT.ModelWithBlocked.ModifyBlockByBlockIPWithCount(ctx, "127.0.0.1", 3)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("result: 更新成功 ")
	}
}

// TestModifyBlockByBlockIPWithBlockend: 测试根据IP修改blockend
func TestModifyBlockByBlockIPWithBlockend(t *testing.T) {
	ctx := context.Background()
	err := tests.SVC_CONTEXT.ModelWithBlocked.ModifyBlockByBlockIPWithBlockend(ctx, "127.0.0.1", time.Now().Add(time.Hour*48))
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("result: 更新成功 ")
	}
}

// TestModifyBlockByBlockIPWithCountAndBlockend: 测试根据IP修改blockend和count
func TestModifyBlockByBlockIPWithCountAndBlockend(t *testing.T) {
	ctx := context.Background()
	err := tests.SVC_CONTEXT.ModelWithBlocked.ModifyBlockByBlockIPWithCountAndBlockend(ctx, "127.0.0.1", time.Now().Add(time.Hour*72), 3)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("result: 更新成功 ")
	}
}

// TestDeleteBlockByBlockIP: 测试根据IP删除
func TestDeleteBlockByBlockIP(t *testing.T) {
	ctx := context.Background()
	_, err := tests.SVC_CONTEXT.ModelWithBlocked.DeleteBlockByBlockIP(ctx, "127.0.0.1")
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("result: 删除成功  ")
	}
}

// TestInsertBlocked: 测试插入blocked数据
func TestInsertBlocked(t *testing.T) {
	ctx := context.Background()
	err := tests.SVC_CONTEXT.ModelWithBlocked.Insert(ctx, &model.Blocked{
		BlockIP:    "127.0.0.1",
		BlockEnd:   time.Now().Add(time.Hour * 24),
		BlockCount: 1,
	})
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("result: 插入成功")
	}
}
