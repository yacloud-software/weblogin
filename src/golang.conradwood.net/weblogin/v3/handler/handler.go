package handler

import (
	"context"
	pb "golang.conradwood.net/apis/weblogin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) StartGRPC(port int) error {
	return nil
}
func (h *Handler) GetLoginPage(ctx context.Context, req *pb.WebloginRequest) (*pb.WebloginResponse, error) {
	return nil, nil
}
