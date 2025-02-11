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

func getDbByConfig(configName string) (*redis.Client, error) {
	data, err := os.ReadFile(configName)
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

func Connect(configName string) (*RedisDb, error) {
	db, err := getDbByConfig(configName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return NewRedisDb(db), nil
}
