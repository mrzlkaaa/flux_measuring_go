package main

import (
	"flux_meas/database"
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
	router.GET("/api/detector_params/:param/:value", RetrieveDetectorByParam)
	// router.GET("/api/detector_params", RetrieveDetectorParam)
	router.Run(":8080")
}

func RetrieveSamples(c *gin.Context) {
	param := c.Param("id")
	id_converted, _ := strconv.ParseInt(param, 10, 64)
	queryResponse := wire.PopulateStruct(id_converted, database.DbConnection)
	c.IndentedJSON(http.StatusOK, queryResponse)
}

func RetrieveDetectorByParam(c *gin.Context) {
	conn := database.NewConnection()
	param := c.Param("param")
	value := c.Param("value")
	s := detector_params.FoilsStoreConstructor()
	// queryResponse := s.Populate(value, param, database.DbConnection)
	queryResponse := s.Populate(value, param, conn.Db)
	// queryResponse := s.Populate(param, DbConnection)
	c.IndentedJSON(http.StatusOK, *queryResponse)
}

// func RetrieveAllDetectorParams(c *gin.Context) {
// 	var s detector_params.QueryPopulater = detector_params.FoilsStoreConstructor()
// 	queryResponse := detector_params.PopulatingByParam(s, value, param, DbConnection)
// 	// queryResponse := s.Populate(param, DbConnection)
// 	c.IndentedJSON(http.StatusOK, *queryResponse)
// }

// func RetrieveDetectorParams(c *gin.Context) {
// 	var ss detector_params.PopulateStruct = []detector_params.FoilsStore{}
// 	queryResponse := ss.PopulateByAll(DbConnection)
// 	c.IndentedJSON(http.StatusOK, *queryResponse)
// }
