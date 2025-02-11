package tasks

import (
	tasks "ginProject1/internal/tasks/neuron_model"
	"ginProject1/internal/tasks/postprocess"
	"ginProject1/internal/tasks/service"
	"ginProject1/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Tasks(taskService service.INeuronApiService, logger *logger.Logger) func(c *gin.Context) {
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
		response, err := taskService.Do(gigaChatRequest)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		gigaChatResponse, err := taskService.ReadNeuronModelAnswer(response)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = taskService.SaveAction(gigaChatRequest.Prompt, gigaChatResponse.GetNeuronRawAnswer())
		if err != nil {
			logger.Error(err)
		}

		testCases, err := postprocess.GetTestsByResponse(gigaChatResponse)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, taskService.ToResponse(testCases))
	}
}

func ListModels(modelService service.IListService, logger *logger.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		answer, err := modelService.List()
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, answer)
	}
}
