package main

import (
	"github.com/gabriel-valin/mmr/configs"
	"github.com/gabriel-valin/mmr/internal/infra/webserver"
)

func main() {
	configs.DbConn()
	webserver.Start()
}
