package cache

import (
	"ginProject1/pkg/database"
	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type RedisDb struct {
	*redis.Client
}

func getDbByConfig() (*redis.Client, error) {
	cwd, _ := os.Getwd()
	log.Print("я тут", cwd)
	data, err := os.ReadFile("config/tasks.yaml")
	if err != nil {
		return nil, err
	}
	config := database.SwitchDbByEnv()
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return redis.NewClient(config.ToStringDriveRedis()), nil
}

func NewRedisDb(db *redis.Client) *RedisDb {
	return &RedisDb{db}
}

func Connect() (*RedisDb, error) {
	db, err := getDbByConfig()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return NewRedisDb(db), nil
}
