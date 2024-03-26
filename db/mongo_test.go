package db_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/lius-new/liusnew-blog-backend-server/db"
	"github.com/lius-new/liusnew-blog-backend-server/internal/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		panic(fmt.Sprintf("No .env file : %s", err.Error()))
	}
}

func TestMongo(t *testing.T) {

	pool := db.NewMongoDBPool()

	mongoClient, ctx := pool.GetClient()

	err := mongoClient.Ping(ctx, readpref.Primary())

	if err != nil {
		logger.Debug(err.Error())
	} else {
		logger.Info("mongodb server connect success")
	}
	coll := mongoClient.Database("liusnew-blog").Collection("user")
	_, err = coll.InsertOne(context.Background(), bson.D{
		{"username", "lius6666"},
		{"password", "14569636547"},
	})

	log.Println(err)

	pool.ReleaseClient(mongoClient)
}

func TestMongoInserUser(t *testing.T) {

	pool := db.NewMongoDBPool()

	client, ctx := pool.GetClient()

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
		{"password", "14569636547"},
	})

	pool.ReleaseClient(client)

}
