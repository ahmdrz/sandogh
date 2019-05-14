package fileserver

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/ahmdrz/sandogh/src/config"
)

type Server struct {
	engine *gin.Engine
	config *config.Configuration
}

func Initialize() (*Server, error) {
	var err error
	s := &Server{}

	s.config, err = config.Read()
	if err != nil {
		return nil, err
	}

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.GET("/status", s.healthCheckHandler)
	engine.Use(gin.Logger())

	files := engine.Group("/files", s.authMiddleware(), s.directoryCheck())
	{
		files.POST("/:directory/:name", s.postHandler)
		files.DELETE("/:directory/:name", s.deleteHandler)
		files.GET("/:directory/:name", s.getHandler)
	}
	s.engine = engine

	return s, nil
}

func (s *Server) Run() error {
	log.Println("Running on", s.config.ListenAddr)
	return s.engine.Run(s.config.ListenAddr)
}
