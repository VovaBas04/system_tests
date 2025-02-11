package codechecker

import (
	"ginProject1/internal/codechecker/service"
	"ginProject1/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckCode(serviceCodeChecker service.IService, logger *logger.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := NewTestRequest()
		err := c.ShouldBindJSON(request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		dto, err := ConvertRequestToDtoCodeCheck(request)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		ok, err := serviceCodeChecker.CheckCode(dto.Code, dto.Tests)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, gin.H{"ok": ok})
	}
}
