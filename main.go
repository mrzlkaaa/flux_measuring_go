package main

import (
	"flux_meas/database"
	"flux_meas/detecParams"
	"flux_meas/server"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	conn := database.DbConnection("foilsstore")
	storage := database.NewConnection(conn)
	param_service := detecParams.NewParamsService(storage)
	server := server.NewServer(router, param_service)
	server.Run()
}
