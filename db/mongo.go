package db

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBPool struct {
	Ctx  context.Context
	pool *sync.Pool
}

func NewMongoDBPool() *MongoDBPool {
	uri := os.Getenv("MONGODB_URI")
	if len(uri) == 0 {
		panic("MONGODB_URI variable not found")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	pool := &sync.Pool{
		New: func() any {

			defer func() {
				if ok := recover(); ok != nil {
					log.Panic("mongodb server connect error")
				}
			}()

			client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
			if err != nil {
				panic(err)
			}

			return client

		},
	}

	return &MongoDBPool{ctx, pool}
}

// GetClient: 获取连接池中的mongodb client
func (p *MongoDBPool) GetClient() *mongo.Client {
	return p.pool.Get().(*mongo.Client)
}

// ReleaseClient: 将连接池中的mongodb client放回去
func (p *MongoDBPool) ReleaseClient(client *mongo.Client) {
	p.pool.Put(client)
}
