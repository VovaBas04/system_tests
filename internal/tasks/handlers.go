package tasks

import (
	tasks "ginProject1/internal/tasks/neuron_model"
	"ginProject1/internal/tasks/postprocess"
	"ginProject1/pkg/database/cache"
	"ginProject1/pkg/database/postgres"
	"ginProject1/pkg/logger"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Tasks(db *postgres.PostgresDb, logger *logger.Logger, redisDb *cache.RedisDb) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := &PromptRequest{}
		err := c.ShouldBindJSON(request)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gigaChatRequest := tasks.NewGigaChatRequest()
		dto := ConvertRequestToDtoConstraint(request)
		err = gigaChatRequest.PreparedRequest(dto.task, dto.constraints)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response, err := gigaChatRequest.Do()
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		gigaChatResponse, err := ReadNeuronModelAnswer(response)
		log.Println(gigaChatResponse.GetNeuronRawAnswer())
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		testCases, err := postprocess.GetTestsByResponse(gigaChatResponse)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, ToResponse(testCases))
	}
}

func Test(db *postgres.PostgresDb, logger *logger.Logger, redisDb *cache.RedisDb) func(c *gin.Context) {
	return func(c *gin.Context) {
		token, err := tasks.GetToken()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": token,
		})
	}
}
