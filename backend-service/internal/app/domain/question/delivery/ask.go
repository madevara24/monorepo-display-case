package delivery

import (
	"backend-service/internal/app/domain/question/usecase/ask"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madevara24/go-common/request"
	"github.com/madevara24/go-common/response"
)

func Ask(service *ask.AskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := ask.Request{}

		if err := request.UnmarshalJSON(c, &req); err != nil {
			return
		}

		res, err := service.Execute(c.Copy().Request.Context(), req)
		if err != nil {
			response.WriteError(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "successfully ask question",
			"data":    res,
		})
	}
}
