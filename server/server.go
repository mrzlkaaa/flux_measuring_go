package server

import (
	"flux_meas/detecParams"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine       *gin.Engine
	paramService detecParams.ParamsService
}

func NewServer(engine *gin.Engine, paramService detecParams.ParamsService) *Server {
	return &Server{engine: engine, paramService: paramService}
}

func (s *Server) Run() error {
	r := s.Router()
	r.Use(cors.Default())
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
	return err
}
