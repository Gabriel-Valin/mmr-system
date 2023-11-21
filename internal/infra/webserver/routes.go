package webserver

import (
	"github.com/gabriel-valin/mmr/internal/infra/webserver/handler"
	"github.com/gin-gonic/gin"
)

func StartRoutes(router *gin.Engine) {
	basePath := "/api/v1"
	apiV1 := router.Group(basePath)
	{
		apiV1.GET("/mmr/:id", handler.GetMmrPlayerHandler)
		apiV1.PATCH("/mmr/:id", handler.UpdateMmrPlayerHandler)
	}
	{
		apiV1.POST("/player", handler.RegisterPlayerHandler)
	}
}
