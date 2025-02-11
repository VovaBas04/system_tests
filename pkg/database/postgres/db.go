package postgres

import (
	"ginProject1/pkg/database"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type PostgresDb struct {
	sqlx.DB
}

func getDbByConfigPostgres(configName string) (*sqlx.DB, error) {
	data, err := os.ReadFile(configName)
	if err != nil {
		return nil, err
	}
	config := database.SwitchDbByEnv()
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return sqlx.Connect("postgres", config.ToStringDrivePostgres())
}

func NewPostgresDB(db *sqlx.DB) *PostgresDb {
	return &PostgresDb{*db}
}

func Connect(configName string) (*PostgresDb, error) {
	db, err := getDbByConfigPostgres(configName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return NewPostgresDB(db), nil
}
