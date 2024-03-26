package db_test

import (
	"context"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/lius-new/liusnew-blog-backend-server/db"
	"github.com/lius-new/liusnew-blog-backend-server/internal/logger"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		panic("No .env file")
	}
}
func TestMongo(t *testing.T) {

	pool := db.NewMongoDBPool()

	mongoClient := pool.GetClient()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := mongoClient.Ping(ctx, readpref.Primary())

	if err != nil {
		logger.Debug(err.Error())
	} else {
		logger.Info("mongodb server connect success")
	}

	pool.ReleaseClient(mongoClient)
}
