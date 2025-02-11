package mongo

import (
	"ginProject1/pkg/database"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type MongoDb struct {
	*mongo.Client
}

func getDbByConfigMongo(configName string) (*mongo.Client, error) {
	data, err := os.ReadFile(configName)
	if err != nil {
		return nil, err
	}
	config := database.SwitchReplicatedDbByEnv()
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return mongo.Connect(options.Client().ApplyURI(config.ToStringDriveMongo()))
}

func NewMongoDB(db *mongo.Client) *MongoDb {
	return &MongoDb{db}
}

func Connect(configName string) (*MongoDb, error) {
	db, err := getDbByConfigMongo(configName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return NewMongoDB(db), nil
}
