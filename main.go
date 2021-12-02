package main

import (
	"flux_meas/database"
	"flux_meas/detecParams"
	"flux_meas/server"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(CORSMiddleware())
	conn := database.DbConnection("foilsstore")
	storage := database.NewConnection(conn)
	param_service := detecParams.NewParamsService(storage)
	server := server.NewServer(router, param_service)
	server.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
