package handler

import (
	"context"
	"fmt"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/weblogin/v3/common"
)

func (w *Handler) ServeHTML(ctx context.Context, req *pb.WebloginRequest) (*pb.WebloginResponse, error) {
	request := common.NewRequest(ctx, req)
	request.Debugf("request started")
	if request.Req.Host == common.SSOHost() && request.Req.Path == "/weblogin/login" {
		return w.GetLoginPage(ctx, req)
	}
	return nil, nil
}

func (w *Handler) GetVerifyEmail(ctx context.Context, req *pb.WebloginRequest) (*pb.EmailPageResponse, error) {
	return nil, nil
}
func (w *Handler) SaveState(ctx context.Context, req *pb.WebloginRequest) (*pb.StateResponse, error) {
	return nil, nil
}
func (w *Handler) GetLoginPage(ctx context.Context, req *pb.WebloginRequest) (*pb.WebloginResponse, error) {
	request := common.NewRequest(ctx, req)
	request.Debugf("request started")
	if !request.State().HaveTriedGlobalCookie {
		if request.Req.Host == common.SSOHost() {
			request.State().HaveTriedGlobalCookie = true
			request.SaveState()
			request.RedirectTo(fmt.Sprintf("https://%s/weblogin/login%s", request.State().TriggerHost, request.GetURLQuery()))
		} else {
			// try redirecting to sso, so that we can retry with a global cookie
			request.RedirectTo(fmt.Sprintf("https://%s/weblogin/login%s", common.SSOHost(), request.GetURLQuery()))
		}
	}

	// we tried the global login, but it did not help
	// next step: render a login page
	request.Error(renderLoginPage(request))
	return request.Close()
}
func (w *Handler) VerifyURL(ctx context.Context, req *pb.WebloginRequest) (*pb.WebloginResponse, error) {
	return nil, nil
}
