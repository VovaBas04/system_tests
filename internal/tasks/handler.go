package tasks

import (
	tasks "ginProject1/internal/tasks/neuron_model"
	"ginProject1/internal/tasks/postprocess"
	"ginProject1/internal/tasks/service"
	"ginProject1/pkg/logger"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Tasks(service service.IService, logger *logger.Logger) func(c *gin.Context) {
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
		response, err := service.Do(gigaChatRequest)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		gigaChatResponse, err := service.ReadNeuronModelAnswer(response)
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
		c.JSON(http.StatusOK, service.ToResponse(testCases))
	}
}
