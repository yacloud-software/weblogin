package handler

import (
	"fmt"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/server"
	"google.golang.org/grpc"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) StartGRPC(port int) error {
	sd := server.NewServerDef()
	sd.SetPort(port)
	sd.SetRegister(func(server *grpc.Server) error {
		pb.RegisterWebloginServer(server, h)
		return nil
	})
	err := server.ServerStartup(sd)
	if err != nil {
		s := fmt.Sprintf("failed to start server: %s\n", err)
		panic(s)
	}
	return nil
}
