package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"goweb/internal/conf"
	"goweb/internal/di"
	"goweb/internal/web/server"
)

func main() {
	flag.Parse()
	if err := di.Init(); err != nil {
		panic(err)
	}

	ginserver := server.New(gin.New())
	if err := ginserver.Run(conf.Conf.Server.WebAddress); err != nil {
		panic(err)
	}
}
