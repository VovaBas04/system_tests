package repository

import (
	"context"
	"fmt"
	"ginProject1/pkg/database/cache"
	"ginProject1/pkg/database/postgres"
	"time"
)

const ActionTable = "actions"

const ModelsTable = "models"

type IRepository interface {
	GetTokenFromCache() (string, error)
	SetTokenToCache(token string)
	SaveAction(input string, output string) error
	ListModels() ([]Model, error)
}

type TaskRepository struct {
	redisDb    *cache.RedisDb
	postgresDb *postgres.PostgresDb
}

func NewTaskRepository(redisDb *cache.RedisDb, postgresDb *postgres.PostgresDb) *TaskRepository {
	return &TaskRepository{redisDb: redisDb, postgresDb: postgresDb}
}

func (repository *TaskRepository) GetTokenFromCache() (string, error) {
	cacheKey := "token"
	ctx := context.Background()

	return repository.redisDb.Get(ctx, cacheKey).Result()
}

func (repository *TaskRepository) SetTokenToCache(token string) {
	cacheKey := "token"
	ctx := context.Background()

	repository.redisDb.Set(ctx, cacheKey, token, 30*time.Minute)
}

func (repository *TaskRepository) SaveAction(input string, output string) error {
	_, err := repository.postgresDb.Exec(fmt.Sprintf("INSERT INTO %s (input, output) VALUES ('%s', '%s')", ActionTable, input, output))
	fmt.Println(fmt.Sprintf("INSERT INTO %s (input, output) VALUES ('%s', '%s')", ActionTable, input, output))
	if err != nil {
		return err
	}

	return nil
}

type Model struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

func (repository *TaskRepository) ListModels() ([]Model, error) {
	var models []Model
	err := repository.postgresDb.Select(&models, fmt.Sprintf("SELECT id, name FROM %s", ModelsTable))
	if err != nil {
		return nil, err
	}

	return models, nil
}
