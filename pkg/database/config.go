package database

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

type ConfigDb interface {
	ToStringDrivePostgres() string
	ToStringDriveRedis() *redis.Options
}

type ReplicateConfigDb interface {
	ToStringDriveMongo() string
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

type MongoParamsDbConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"name"`
}

type DatabasesConfig struct {
	Redis    RedisParamsDbConfig    `yaml:"redis"`
	Postgres PostgresParamsDbConfig `yaml:"postgres"`
}

type ReplicatedDatabasesConfig struct {
	Redis RedisParamsDbConfig `yaml:"redis"`
	Mongo MongoParamsDbConfig `yaml:"mongo"`
}

type LocalhostDb struct {
	Db DatabasesConfig `yaml:"localhost"`
}

type DockerDb struct {
	Db DatabasesConfig `yaml:"docker"`
}

type ReplicatedLocalhostDb struct {
	Db ReplicatedDatabasesConfig `yaml:"localhost"`
}

type ReplicatedDockerDb struct {
	Db ReplicatedDatabasesConfig `yaml:"docker"`
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

func (DbConfig *ReplicatedLocalhostDb) ToStringDriveMongo() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?directConnection=true&serverSelectionTimeoutMS=2000&authSource=admin", DbConfig.Db.Mongo.Username, DbConfig.Db.Mongo.Password, DbConfig.Db.Mongo.Host, DbConfig.Db.Mongo.Port, DbConfig.Db.Mongo.Database)
}

func (DbConfig *ReplicatedDockerDb) ToStringDriveMongo() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?directConnection=true&serverSelectionTimeoutMS=2000&authSource=admin", DbConfig.Db.Mongo.Username, DbConfig.Db.Mongo.Password, DbConfig.Db.Mongo.Host, DbConfig.Db.Mongo.Port, DbConfig.Db.Mongo.Database)
}

func (DbConfig *ReplicatedLocalhostDb) ToStringDriveRedis() *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", DbConfig.Db.Redis.Host, DbConfig.Db.Redis.Port),
		Password: "",
		DB:       0,
	}
}

func (DbConfig *ReplicatedDockerDb) ToStringDriveRedis() *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", DbConfig.Db.Redis.Host, DbConfig.Db.Redis.Port),
		Password: "",
		DB:       0,
	}
}
