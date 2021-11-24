package main

import (
	"flux_meas/detector_params"
	"flux_meas/wire"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api/experiment/:id", RetrieveSamples)
	router.GET("/api/detector_params/:nuclide", RetrieveDetectorParams)
	router.Run(":8080")
}

func RetrieveSamples(c *gin.Context) {
	param := c.Param("id")
	id_converted, _ := strconv.ParseInt(param, 10, 64)
	queryResponse := wire.PopulateStruct(id_converted, DbConnection)
	c.IndentedJSON(http.StatusOK, queryResponse)
}

func RetrieveDetectorParams(c *gin.Context) {
	param := c.Param("nuclide")
	queryResponse := detector_params.PopulateStruct(param, DbConnection)
	c.IndentedJSON(http.StatusOK, *queryResponse)
}
