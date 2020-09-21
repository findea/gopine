package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"goweb/internal/web/controller"
	"goweb/internal/web/jsonrpc/service"
	"goweb/internal/web/middleware"
	"goweb/pkg/util/path"
	"net/http"
	"path/filepath"
	"time"
)

func (s *Server) initRouter() {
	s.OPTIONS("/*all", middleware.AccessControlAllow)

	s.initHTMLRouter()

	s.initAPIRouter()

	s.initJSONRPCRouter()
}

func (s *Server) initHTMLRouter() {
	p, err := path.FindPath("template", 5)
	if err != nil {
		panic("template not found")
	}

	s.LoadHTMLGlob(fmt.Sprintf("%s%c*", p, filepath.Separator))

	group := s.Group("/")

	group.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.gohtml", gin.H{
			"message": time.Now().Format(time.RFC3339),
		})
	})
}

func (s *Server) initAPIRouter() {
	group := s.Group("/api", middleware.AccessControlAllow)

	// email code
	group.POST("/emailcode/send", middleware.NewCounterHandler(10, 10*time.Minute), controller.EmailCodeController.Send)

	// user
	group.POST("/user/register", controller.UserController.Register)
	group.POST("/user/login", controller.UserController.Login)
}

func (s *Server) initJSONRPCRouter() {
	rs := rpc.NewServer()

	group := s.Group("/")
	group.POST("/jsonrpc", func(c *gin.Context) {
		rs.ServeHTTP(c.Writer, c.Request)
	})

	rs.RegisterCodec(json.NewCodec(), "application/json")
	rs.RegisterService(new(service.HelloService), "")
}
