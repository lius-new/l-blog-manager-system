package models

import (
	"context"

	"github.com/lius-new/liusnew-blog-backend-server/internal/logger"
	"github.com/lius-new/liusnew-blog-backend-server/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	client := Pool.GetClient()
	ctx := context.Background()

	coll := client.Database("liusnew-blog").Collection("user")

	count, err := coll.CountDocuments(ctx, bson.D{})
	if err != nil {
		logger.Warn(err.Error())
	}

	if count != 0 {
		return
	}

	coll.InsertOne(ctx, bson.D{
		{"username", "lius6666"},
		{"password", utils.MD5("14569636547")},
	})

	Pool.ReleaseClient(client)
}

func Login(username, password string) ([]string, error) {
	client := Pool.GetClient()
	defer Pool.ReleaseClient(client)

	ctx := context.Background()

	coll := client.Database("liusnew-blog").Collection("user")

	if r, err := coll.FindOne(ctx, bson.D{
		{"username", username},
		{"password", utils.MD5(password)},
	}).Raw(); err != nil {
		return []string{}, err
	} else {
		return []string{r.Lookup("_id").String(), r.Lookup("username").String()}, nil
	}
}
