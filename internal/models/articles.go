package models

import (
	"context"
	"errors"
	"time"

	"github.com/lius-new/liusnew-blog-backend-server/internal/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Article struct {
	Id      string
	Title   string
	Content string
	Tags    []string
	Covers  []string
	Status  bool
	Time    int64
}

func getItemWithBson(b bson.M, item string) (interface{}, error) {
	res := b[item]
	if res != nil {
		return res, nil
	}
	return nil, errors.New("key not found")

}
func BsonToArticle(b bson.M) *Article {

	primitiveASlcieToStringSlice := func(pSlice primitive.A) []string {
		s := make([]string, 0)

		for _, v := range pSlice {
			if itemStr, ok := v.(string); ok {
				s = append(s, itemStr)
			}
		}

		return s
	}

	a := &Article{}
	if value, err := getItemWithBson(b, "_id"); err == nil {
		a.Id = value.(primitive.ObjectID).Hex()
	}
	if value, err := getItemWithBson(b, "title"); err == nil {
		a.Title = value.(string)
	}
	if value, err := getItemWithBson(b, "content"); err == nil {
		a.Content = value.(string)
	}
	if value, err := getItemWithBson(b, "tags"); err == nil {
		a.Tags = primitiveASlcieToStringSlice(value.(primitive.A))
	} else {
		a.Tags = make([]string, 0)
	}
	if value, err := getItemWithBson(b, "covers"); err == nil {
		a.Covers = primitiveASlcieToStringSlice(value.(primitive.A))
	} else {
		a.Covers = make([]string, 0)
	}
	if value, err := getItemWithBson(b, "status"); err == nil {
		a.Status = value.(bool)
	}
	if value, err := getItemWithBson(b, "time"); err == nil {
		a.Time = value.(int64)
	}
	return a
}

func (a *Article) ToBson() (d bson.D) {
	d = append(d, bson.E{Key: "title", Value: a.Title})
	d = append(d, bson.E{Key: "content", Value: a.Content})
	d = append(d, bson.E{Key: "tags", Value: a.Tags})
	d = append(d, bson.E{Key: "covers", Value: a.Covers})
	d = append(d, bson.E{Key: "status", Value: a.Status})
	d = append(d, bson.E{Key: "time", Value: a.Time})
	return
}

func CreateArticles(title, content string, tags, covers []string) (*Article, error) {
	save := func(a *Article) (string, error) {
		client := Pool.GetClient()
		Pool.ReleaseClient(client)

		coll := client.Database("liusnew-blog").Collection("articles")

		ctx := context.Background()

		if count, err := coll.CountDocuments(ctx, bson.D{{"title", a.Title}, {"content", a.Content}}); err != nil {
			panic(err)
		} else if count > 0 {
			return "", errors.New("article exist")
		}

		insertRes, err := coll.InsertOne(ctx, a.ToBson())

		if err != nil {
			panic(err)
		}

		return insertRes.InsertedID.(primitive.ObjectID).Hex(), nil
	}

	a := Article{
		"", title, content, tags, covers, true, time.Now().UnixNano(),
	}
	id, err := save(&a)
	a.Id = id

	return &a, err
}
func ModifyArticleStatus(id string, status bool) (*Article, error) {
	client := Pool.GetClient()
	Pool.ReleaseClient(client)
	coll := client.Database("liusnew-blog").Collection("articles")
	ctx := context.Background()

	modify := func() error {
		objectId, _ := primitive.ObjectIDFromHex(id)

		if count, err := coll.CountDocuments(ctx, bson.D{{"_id", objectId}}); err != nil {
			panic(err)
		} else if count == 0 {
			return errors.New("article not found")
		}

		_, err := coll.UpdateOne(ctx, bson.D{{"_id", objectId}}, bson.D{{"$set", bson.D{{"status", status}, {"time", time.Now().UnixNano()}}}})

		if err != nil {
			panic(err)
		}

		return nil
	}

	err := modify()
	return ViewArticle(id), err
}

func ModifyArticles(id, title, content string, tags, covers []string, status bool) (*Article, error) {
	client := Pool.GetClient()
	Pool.ReleaseClient(client)
	coll := client.Database("liusnew-blog").Collection("articles")
	ctx := context.Background()

	modify := func(a *Article) error {
		objectId, _ := primitive.ObjectIDFromHex(a.Id)

		if count, err := coll.CountDocuments(ctx, bson.D{{"_id", objectId}}); err != nil {
			panic(err)
		} else if count == 0 {
			return errors.New("article not found")
		}

		_, err := coll.UpdateOne(ctx, bson.D{{"_id", objectId}}, bson.D{{"$set", a.ToBson()}})

		if err != nil {
			panic(err)
		}

		return nil
	}

	a := Article{
		id, title, content, tags, covers, status, time.Now().UnixNano(),
	}

	err := modify(&a)

	return ViewArticle(id), err
}

func DeleteArticles(id string) {
	client := Pool.GetClient()
	defer Pool.ReleaseClient(client)
	coll := client.Database("liusnew-blog").Collection("articles")
	ctx := context.Background()

	delete := func(id string) {
		objectId, _ := primitive.ObjectIDFromHex(id)
		deleteRes, err := coll.DeleteOne(ctx, bson.D{{"_id", objectId}})
		if err != nil {
			panic(err)
		}
		logger.Debug(deleteRes)
	}
	delete(id)
}

func ViewArticle(id string) *Article {
	client := Pool.GetClient()
	defer Pool.ReleaseClient(client)
	coll := client.Database("liusnew-blog").Collection("articles")
	ctx := context.Background()

	view := func() *Article {
		objectId, _ := primitive.ObjectIDFromHex(id)
		res := coll.FindOne(ctx, bson.D{{"_id", objectId}})
		if res.Err() != nil {
			panic(res.Err())
		}
		var tempResult bson.M
		res.Decode(&tempResult)
		return BsonToArticle(tempResult)
	}
	return view()
}
func ViewArticles(pageSize, pageNumber int64, showDisable bool) ([]Article, int64) {
	client := Pool.GetClient()
	defer Pool.ReleaseClient(client)
	coll := client.Database("liusnew-blog").Collection("articles")
	ctx := context.Background()

	view := func() (articles []Article) {
		findOptions := options.Find()

		if pageNumber <= 0 {
			pageNumber = 1
		}

		findOptions.SetLimit(pageSize)
		findOptions.SetSkip(pageSize * (pageNumber - 1))
		findOptions.SetSort(bson.M{"time": 1})

		var (
			cur *mongo.Cursor
			err error
		)
		if showDisable {
			cur, err = coll.Find(ctx, bson.D{{}}, findOptions)
		} else {
			cur, err = coll.Find(ctx, bson.D{{"status", true}}, findOptions)
		}

		if err != nil {
			panic(err)
		}

		for cur.Next(ctx) {
			var tempResult bson.M
			err := cur.Decode(&tempResult)

			if err != nil {
				logger.Debug(err)
			}
			articles = append(articles, *BsonToArticle(tempResult))
		}

		return
	}

	count := func() int64 {
		c, err := coll.CountDocuments(ctx, bson.D{{}})
		if err != nil {
			panic(err)
		}
		return int64(c/pageSize) + 1
	}

	return view(), count()
}

func SearchArticle(title string) []Article {
	// {title: {$regex: /go/i}}
	client := Pool.GetClient()
	defer Pool.ReleaseClient(client)
	coll := client.Database("liusnew-blog").Collection("articles")
	ctx := context.Background()

	view := func() (articles []Article) {

		regex := primitive.Regex{Pattern: title, Options: "i"}

		cur, err := coll.Find(ctx, bson.D{{"title", regex}})
		if err != nil {
			panic(err)
		}

		for cur.Next(ctx) {
			var tempResult bson.M
			err := cur.Decode(&tempResult)

			if err != nil {
				logger.Debug(err)
			}
			articles = append(articles, *BsonToArticle(tempResult))
		}

		return
	}

	return view()
}
