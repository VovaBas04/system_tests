package tasks

import (
	"ginProject1/internal/tasks/repository"
	"ginProject1/internal/tasks/service"
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
	configPath := "config/tasks.yaml"
	postgresDb, err := postgres.Connect(configPath)
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *postgres.PostgresDb) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(postgresDb)
	redisDb, err := cache.Connect(configPath)
	defer redisDb.Close()
	if err != nil {
		log.Fatal(err)
	}
	repositoryTask := repository.NewTaskRepository(redisDb, postgresDb)
	serviceNeuron := service.NewNeuronApiServiceService(repositoryTask)
	serviceList := service.NewListService(repositoryTask)
	app.POST("/tasks", Tasks(serviceNeuron, log))
	app.GET("/models", ListModels(serviceList, log))
	err = app.Run(":8082")
	if err != nil {
		log.Fatal(err)
	}

	log.Info("server started")
}
