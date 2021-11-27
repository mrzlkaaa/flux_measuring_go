package main

import (
	"flux_meas/database"
	"flux_meas/detector_params"
	"flux_meas/wire"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServiceMethods interface {
	Populate(value, param string, db *gorm.DB) *detector_params.FoilsStore
}

type Service interface {
	PopulateByAll(db *gorm.DB) *[]detector_params.FoilsStore
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api/experiment/:id", RetrieveSamples)
	router.GET("/api/detector_params/:param/:value", RetrieveDetectorByParam)
	router.GET("/api/detector_params", RetrieveAllDetectorParams)
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
	var s ServiceMethods = detector_params.FoilsStoreConstructor()
	queryResponse := s.Populate(value, param, conn.Db)
	c.IndentedJSON(http.StatusOK, *queryResponse)
}

func RetrieveAllDetectorParams(c *gin.Context) {
	conn := database.NewConnection()
	queryResponse := detector_params.PopulateByAll(conn.Db)
	c.IndentedJSON(http.StatusOK, *queryResponse)
}
