package handler

import (
	"fmt"
	"net/http"

	"github.com/gabriel-valin/mmr/configs"
	"github.com/gabriel-valin/mmr/internal/dtos"
	"github.com/gabriel-valin/mmr/internal/services"
	"github.com/gabriel-valin/mmr/pkg/cryptography"
	"github.com/gin-gonic/gin"
)

func RegisterPlayerHandler(ctx *gin.Context) {
	db, err := configs.DbConn()
	if err != nil {
		fmt.Print(err)
	}
	request := dtos.RequestRegisterPlayerDTO{}
	ctx.BindJSON(&request)

	if err = request.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	hash, err := cryptography.HashPassword(request.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Try again soon later"})
		return
	}

	createPlayer := services.NewCreatePlayerService(db)
	err = createPlayer.Create(request.Username, string(hash), request.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user created with successfull"})
}
