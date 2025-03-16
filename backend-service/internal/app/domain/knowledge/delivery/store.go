package delivery

import (
	"backend-service/internal/app/domain/knowledge/usecase/store"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madevara24/go-common/request"
	"github.com/madevara24/go-common/response"
)

func Store(service *store.StoreUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := store.Request{}

		if err := request.UnmarshalJSON(c, &req); err != nil {
			return
		}

		err := service.Execute(c.Copy().Request.Context(), req)
		if err != nil {
			response.WriteError(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "successfully store knowledge",
		})
	}
}

func StoreV2(service *store.StoreUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := store.RequestV2{}

		if err := request.UnmarshalJSON(c, &req); err != nil {
			return
		}

		err := service.ExecuteV2(c.Copy().Request.Context(), req)
		if err != nil {
			response.WriteError(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "successfully store knowledge",
		})
	}
}
