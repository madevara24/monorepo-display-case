package delivery

import (
	"backend-service/internal/app/domain/system/usecase/healthcheck"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madevara24/go-common/response"
)

func HealthCheck(service *healthcheck.HealthCheckService) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := service.Execute(c.Copy().Request.Context())
		if err != nil {
			response.WriteError(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "successfully get healthcheck",
			"data":    res,
		})
	}
}
