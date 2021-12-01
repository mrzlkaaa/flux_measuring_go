package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Router() *gin.Engine {
	router := s.engine
	params := router.Group("/api")
	{
		params.GET("/detector_params/", s.RetrieveAllDetectorParams())
		params.GET("/detector_params/:param/:value", s.RetrieveDetectorParamByQuery())

	}
	foils := router.Group("/api/foil_detectors")
	{
		foils.GET("/:value", s.RetrieveFoilsByType())
	}
	return router
}

func (s *Server) RetrieveAllDetectorParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		all_data, _ := s.paramService.PopulateByAll()
		c.IndentedJSON(http.StatusOK, all_data)
	}
}

func (s *Server) RetrieveDetectorParamByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Param("param")
		value := c.Param("value")
		get_data, _ := s.paramService.PopulateByParam(value, param)
		fmt.Println(get_data)
		c.IndentedJSON(http.StatusOK, get_data)
	}
}

func (s *Server) RetrieveFoilsByType() gin.HandlerFunc {
	return func(c *gin.Context) {
		value := c.Param("value")
		get_data, _ := s.paramService.PopulateAllByFoilType(value)
		fmt.Println(get_data)
		c.IndentedJSON(http.StatusOK, get_data)
	}
}
