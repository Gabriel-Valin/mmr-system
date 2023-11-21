package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMmrPlayerHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": map[string]any{"mmr": "674"}})
}
