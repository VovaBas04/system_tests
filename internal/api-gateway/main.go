package api_gateway

import (
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
	configPath := "config/api-gateway.yaml"
	config, err := getConfig(configPath)
	if err != nil {
		log.Fatal(err)

		return
	}

	proxy := NewProxy(config)
	app.Use(CORSMiddleware())
	app.Any("/*path", proxy.proxyHandler)

	err = app.Run(":8080")
	if err != nil {
		log.Fatal(err)

		return
	}

	log.Info("server started")
}
