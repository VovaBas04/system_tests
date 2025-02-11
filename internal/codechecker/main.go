package codechecker

import (
	"context"
	"fmt"
	"ginProject1/internal/codechecker/repository"
	"ginProject1/internal/codechecker/service"
	"ginProject1/pkg/database/cache"
	"ginProject1/pkg/database/mongo"
	"ginProject1/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"time"
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
	configPath := "config/codechecker.yaml"
	db, err := mongo.Connect(configPath)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	val := db.Ping(ctx, readpref.Primary())
	fmt.Println("ping", val)
	defer func() {
		ctx := context.Background()
		if err = db.Disconnect(ctx); err != nil {
			log.Fatal("error on disconnect")
		}
	}()
	redisDb, err := cache.Connect(configPath)
	defer redisDb.Close()
	repositoryCodeChecker := repository.NewRepository(db, redisDb)
	serviceCodeChecker := service.NewService(repositoryCodeChecker)
	app.POST("/check", CheckCode(serviceCodeChecker, log))
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(":8081")
	if err != nil {
		log.Fatal(err)
	}

	log.Info("server started")
}
