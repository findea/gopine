package server

import (
	"github.com/gin-gonic/gin"
	"goweb/internal/web/middleware"
	gintrace "goweb/pkg/lighttracer/gin"
)

type Server struct {
	*gin.Engine
}

func New(engine *gin.Engine) *Server {
	s := &Server{engine}
	s.initMiddleWare()
	s.initRouter()
	return s
}

func (s *Server) initMiddleWare() {
	s.Use(
		gintrace.TraceHandler,
		middleware.LoggerHandler,
		middleware.RecoverHandler,
	)
}
