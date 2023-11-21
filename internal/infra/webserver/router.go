package webserver

import (
	"github.com/gabriel-valin/mmr/configs"
	"github.com/gin-gonic/gin"
)

func Start() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	StartRoutes(router)

	port := configs.HTTPPort

	if port == "" {
		port = "8081"
	}

	router.Run("0.0.0.0:" + port)
}
