package models

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/lius-new/liusnew-blog-backend-server/internal/logger"
	"go.mongodb.org/mongo-driver/bson"
)

// 记录访问和将其添加或删除到Block collections中
func Trace(ip, path string) {
	client := Pool.GetClient()
	defer Pool.ReleaseClient(client)

	coll := client.Database("liusnew-blog").Collection("record")

	ctx := context.Background()

	// 将当前访问的用户ip添加到数据库
	now, dayAgo := getExpireTime()
	coll.InsertOne(ctx, bson.D{
		{"ip", ip},
		{"path", path},
		{"time", now.UnixNano()},
	})

	// countCondition 用于统计24小时内访问条件
	countCondition := bson.D{
		{"ip", ip},
		{"time",
			bson.D{
				{"$gte", dayAgo.UnixNano()}, // 大于或等于dayAgo的时间戳(昨天)
				{"$lte", now.UnixNano()},    // 小于或等于当前时间戳(现在)
			},
		},
	}

	// 对某些path启用更加严格的禁用
	if strict(path) {
		countCondition = append(countCondition, bson.E{Key: "path", Value: path})
	}

	// 统计24小时内访问的次数
	count, err := coll.CountDocuments(ctx, countCondition)

	if err != nil {
		logger.Panic(err.Error())
	}

	// 超过24小时访问的限制数量就添加到blocked
	blockColl := client.Database("liusnew-blog").Collection("blocked")
	// 禁用的代码
	blockHandler := func() {
		if res, err := blockColl.FindOne(ctx, bson.D{
			{"ip", ip},
		}).Raw(); err == nil {
			// 如果被block, 那么就获得时间
			blockTime := res.Lookup("time").Int64()
			hours := float64((now.UnixNano() - blockTime)) / float64(time.Hour)

			BLOCKED_TIME, _ := strconv.Atoi(os.Getenv("BLOCKED_TIME"))
			// 判断block时间和当前时间的差，如果时间差大于BlockedTime常量时间那么就移除出Blcoked
			if hours > float64(BLOCKED_TIME) {
				blockColl.DeleteOne(ctx, bson.D{
					{"ip", ip},
				})
			}
		} else {
			blockColl.InsertOne(ctx, bson.D{
				{"ip", ip},
				{"time", now.UnixNano()},
			})
		}
	}

	VISIT_COMMONPAGE_TO_BLACKLIST, _ := strconv.Atoi(os.Getenv("VISIT_COMMONPAGE_TO_BLACKLIST"))
	VISIT_STRICTPAGE_TO_BLACKLIST, _ := strconv.Atoi(os.Getenv("VISIT_STRICTPAGE_TO_BLACKLIST"))

	if count > int64(VISIT_COMMONPAGE_TO_BLACKLIST) {
		blockHandler()
	} else if strict(path) && count > int64(VISIT_STRICTPAGE_TO_BLACKLIST) {
		blockHandler()
	}
}

// strict: 是否是严格访问(受保护的资源)
func strict(path string) bool {
	stricts := []string{"/api/user/login"}

	for _, v := range stricts {
		if path == v {
			return true
		}
	}
	return false
}

// IsBlocked: 指定的ip是否被block(拦截, 即加入黑名单)
func IsBlocked(ip string) bool {
	client := Pool.GetClient()
	defer Pool.ReleaseClient(client)

	coll := client.Database("liusnew-blog").Collection("blocked")

	ctx := context.Background()
	count, err := coll.CountDocuments(ctx, bson.D{{"ip", ip}})
	if count > 0 || err != nil {
		return true
	}

	return false
}

func getExpireTime() (time.Time, time.Time) {
	now := time.Now()
	dayAgo := now.Add(-24 * time.Hour)
	return now, dayAgo
}
