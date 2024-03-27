package models

import (
	"context"
	"strings"

	"github.com/lius-new/liusnew-blog-backend-server/internal/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag struct {
	name   string
	status bool
}

func (t *Tag) ToBson() (d bson.D) {
	d = append(d, bson.E{Key: "name", Value: t.name})
	d = append(d, bson.E{Key: "status", Value: t.status})
	return
}

func SaveTags(tagStrs []string) []string {
	client := Pool.GetClient()
	defer Pool.ReleaseClient(client)
	coll := client.Database("liusnew-blog").Collection("tags")
	ctx := context.Background()

	save := func(t Tag) string {
		res, err := coll.FindOne(ctx, t.ToBson()).Raw()
		if err != nil && strings.Contains(err.Error(), "no documents") {
			if insertRes, err := coll.InsertOne(ctx, t.ToBson()); err == nil {
				return insertRes.InsertedID.(primitive.ObjectID).Hex()
			}
		} else if err != nil {
			panic(err)
		}
		return res.Lookup("_id").ObjectID().Hex()
	}

	tagsId := make([]string, 0)
	for _, v := range tagStrs {
		tag := Tag{
			name:   v,
			status: true,
		}

		tagsId = append(tagsId, save(tag))
	}

	return tagsId
}

func DeleteTags(tagStrs []string) {
	client := Pool.GetClient()
	defer Pool.ReleaseClient(client)
	coll := client.Database("liusnew-blog").Collection("tags")
	ctx := context.Background()

	delete := func(name string) error {
		deleteRes, err := coll.DeleteOne(ctx, bson.D{{"name", name}})
		if err != nil {
			return err
		}
		logger.Debug(deleteRes)
		return nil
	}

	for _, v := range tagStrs {
		if err := delete(v); err != nil {
			panic(err)
		}
	}
}
