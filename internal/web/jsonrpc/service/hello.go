package service

import (
	"goweb/internal/web/jsonrpc/model"
	"net/http"
)

type HelloService struct {

}

func (h *HelloService) Say(_ *http.Request, args *model.HelloArgs, reply *model.HelloReply) error {
	reply.Message = "Hello, " + args.Who + "!"
	return nil
}
