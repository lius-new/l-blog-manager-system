package models

import (
	"context"
	"strings"

	"github.com/lius-new/liusnew-blog-backend-server/internal/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag struct {
	Name   string
	Status bool
}

func BsonToTags(b bson.M) *Tag {
	t := &Tag{}

	t.Name = b["name"].(string)
	t.Status = b["status"].(bool)
	return t
}

func (t *Tag) ToBson() (d bson.D) {
	d = append(d, bson.E{Key: "name", Value: t.Name})
	d = append(d, bson.E{Key: "status", Value: t.Status})
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
		if len(v) == 0 {
			continue
		}
		tag := Tag{
			Name:   v,
			Status: true,
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

func ViewTags() []string {
	client := Pool.GetClient()
	defer Pool.ReleaseClient(client)
	coll := client.Database("liusnew-blog").Collection("tags")
	ctx := context.Background()

	view := func() (tags []string) {
		cur, err := coll.Find(ctx, bson.D{{}})
		if err != nil {
			panic(err)
		}

		for cur.Next(ctx) {
			var tempResult bson.M
			err := cur.Decode(&tempResult)

			if err != nil {
				logger.Debug(err)
			}
			tag := *BsonToTags(tempResult)
			if tag.Status {
				tags = append(tags, tag.Name)
			}
		}
		return
	}

	return view()
}

func ViewArticlesTags(tagIds []string) (tagNames []string) {
	client := Pool.GetClient()
	defer Pool.ReleaseClient(client)
	coll := client.Database("liusnew-blog").Collection("tags")
	ctx := context.Background()

	objectIds := make([]primitive.ObjectID, 0)
	for _, v := range tagIds {
		objectId, _ := primitive.ObjectIDFromHex(v)
		objectIds = append(objectIds, objectId)
	}

	view := func() []string {
		tags := make([]string, 0)
		cur, err := coll.Find(ctx, bson.D{{"_id", bson.D{{"$in", objectIds}}}})
		if err != nil {
			panic(err)
		}

		for cur.Next(ctx) {
			var tempResult bson.M
			err := cur.Decode(&tempResult)

			if err != nil {
				logger.Debug(err)
			}
			tag := *BsonToTags(tempResult)
			if tag.Status {
				tags = append(tags, tag.Name)
			}
		}
		return tags
	}

	return view()
}
