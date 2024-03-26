package models

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/lius-new/liusnew-blog-backend-server/db"
)

var Pool db.MongoDBPool

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(fmt.Sprintf("No .env file : %s", err.Error()))
	}
	Pool = *db.NewMongoDBPool()
}
