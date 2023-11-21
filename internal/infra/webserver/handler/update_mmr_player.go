package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateMmrPlayerHandler(ctx *gin.Context) {
	var request any
	ctx.BindJSON(&request)
	ctx.JSON(http.StatusOK, gin.H{"data": request})
}
