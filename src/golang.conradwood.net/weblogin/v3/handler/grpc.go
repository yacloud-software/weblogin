package handler

import (
	"context"
	"fmt"
	"golang.conradwood.net/apis/themes"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/weblogin/v3/common"
)

func (w *Handler) ServeHTML(ctx context.Context, req *pb.WebloginRequest) (*pb.WebloginResponse, error) {
	if req.Host == common.SSOHost() && req.Path == "/weblogin/login" {
		return w.GetLoginPage(ctx, req)
	}
	request := common.NewRequest(ctx, req)
	request.Debugf("request started to servehtml for \"%s\"", req.Path)
	if req.Path == "/weblogin/forgotPassword" {
		return forgotPassword(request)
	}
	var err error
	res := &pb.WebloginResponse{}
	htr := &themes.HostThemeRequest{Host: request.TriggerHost()}
	if req.Path == "/weblogin/stylesheet.css" {
		b, err := themes.GetThemesClient().GetCSS(ctx, htr)
		if err == nil {
			res.Body = b.Data
			res.MimeType = "text/css"
		}
	}
	request.SetResponse(res)
	request.Error(err)
	return request.Close()
}
func (w *Handler) CreateRegisterEmail(ctx context.Context, cr *pb.RegisterState) (*pb.Email, error) {
	return nil, fmt.Errorf("CreateRegisteremail not implemented in v3")
}
func (w *Handler) IsBasicAuthValid(ctx context.Context, cr *pb.BasicAuthRequest) (*pb.AuthResponse, error) {
	return nil, fmt.Errorf("isbasicauth not implemented in v3")
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
func forgotPassword(request *common.Request) (*pb.WebloginResponse, error) {
	b, err := request.Render("forgot_password", "")
	if err != nil {
		request.Error(err)
		return request.Close()
	}
	w := &pb.WebloginResponse{
		Body:     b,
		MimeType: "text/html",
	}
	request.SetResponse(w)
	return request.Close()

}
