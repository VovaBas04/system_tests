package database

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

type ConfigDb interface {
	ToStringDrivePostgres() string
	ToStringDriveRedis() *redis.Options
}

type PostgresParamsDbConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"name"`
}

type RedisParamsDbConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DatabasesConfig struct {
	Redis    RedisParamsDbConfig    `yaml:"redis"`
	Postgres PostgresParamsDbConfig `yaml:"postgres"`
}

type LocalhostDb struct {
	Db DatabasesConfig `yaml:"localhost"`
}

type DockerDb struct {
	Db DatabasesConfig `yaml:"docker"`
}

func (DbConfig *LocalhostDb) ToStringDrivePostgres() string {
	return fmt.Sprintf("user = %s password = %s dbname = %s port = %s host = %s sslmode=disable", DbConfig.Db.Postgres.Username, DbConfig.Db.Postgres.Password, DbConfig.Db.Postgres.Database, DbConfig.Db.Postgres.Port, DbConfig.Db.Postgres.Host)
}

func (DbConfig *DockerDb) ToStringDrivePostgres() string {
	return fmt.Sprintf("user = %s password = %s dbname = %s port = %s host = %s sslmode=disable", DbConfig.Db.Postgres.Username, DbConfig.Db.Postgres.Password, DbConfig.Db.Postgres.Database, DbConfig.Db.Postgres.Port, DbConfig.Db.Postgres.Host)
}

func (DbConfig *LocalhostDb) ToStringDriveRedis() *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", DbConfig.Db.Redis.Host, DbConfig.Db.Redis.Port),
		Password: "",
		DB:       0,
	}
}

func (DbConfig *DockerDb) ToStringDriveRedis() *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", DbConfig.Db.Redis.Host, DbConfig.Db.Redis.Port),
		Password: "",
		DB:       0,
	}
}
