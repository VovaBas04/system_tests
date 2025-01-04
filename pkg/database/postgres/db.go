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

func getDbByConfigPostgres() (*sqlx.DB, error) {
	data, err := os.ReadFile("config/tasks.yaml")
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

func Connect() (*PostgresDb, error) {
	db, err := getDbByConfigPostgres()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return NewPostgresDB(db), nil
}
