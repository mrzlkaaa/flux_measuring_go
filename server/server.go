package server

import (
	"flux_meas/detecParams"
	"flux_meas/wire"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine       *gin.Engine
	paramService detecParams.ParamsService
	wireService  wire.WireService
}

func NewServer(engine *gin.Engine, paramService detecParams.ParamsService, wireService wire.WireService) *Server {
	return &Server{engine: engine, paramService: paramService, wireService: wireService}
}

func (s *Server) Run() error {

	r := s.Router()
	// r.Use(cors.Default())
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
	return err
}
