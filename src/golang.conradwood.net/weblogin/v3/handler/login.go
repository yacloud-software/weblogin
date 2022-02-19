package handler

import (
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/weblogin/v3/common"
)

type login_render struct {
}

// render a login page.
func renderLoginPage(request *common.Request) error {
	lr := &login_render{}
	b, err := request.Render("login", lr)
	if err != nil {
		return err
	}
	w := &pb.WebloginResponse{
		Body:     b,
		MimeType: "text/html",
	}
	request.SetResponse(w)
	return nil
}
