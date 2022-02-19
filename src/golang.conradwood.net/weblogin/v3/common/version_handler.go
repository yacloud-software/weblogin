package common

import (
	"context"
	pb "golang.conradwood.net/apis/weblogin"
)

type VersionHandler interface {
	StartGRPC(port int) error
	GetLoginPage(ctx context.Context, req *pb.WebloginRequest) (*pb.WebloginResponse, error)
}
