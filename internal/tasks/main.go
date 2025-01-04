package tasks

import (
	"ginProject1/pkg/database/cache"
	"ginProject1/pkg/database/postgres"
	"ginProject1/pkg/logger"
	"github.com/gin-gonic/gin"
)

type App struct {
	*gin.Engine
}

func NewApp() *App {
	return &App{
		gin.Default(),
	}
}
func Run() {
	log := logger.NewLogger()
	app := NewApp()
	db, err := postgres.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *postgres.PostgresDb) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	redisDb, err := cache.Connect()
	if err != nil {
		log.Fatal(err)
	}

	app.POST("/", Tasks(db, log, redisDb))
	app.GET("/test", Test(db, log, redisDb))
	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("server started")
}
