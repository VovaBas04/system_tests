package repository

import (
	"ginProject1/pkg/database/cache"
	"ginProject1/pkg/database/mongo"
)

type IRepository interface {
	Save() error
}

type Repository struct {
	mongoDb *mongo.MongoDb
	redisDb *cache.RedisDb
}

func NewRepository(mongoDb *mongo.MongoDb, redisDb *cache.RedisDb) IRepository {
	return &Repository{mongoDb: mongoDb, redisDb: redisDb}
}

func (r *Repository) Save() error {
	return nil
}
