package requests

import (
	"fmt"
	apb "golang.conradwood.net/apis/auth"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/utils"
)

type LogoutStruct struct {
	Msg  string
	user *apb.User
}

func (l *LogoutStruct) Username() string {
	if l.user == nil {
		return "nobody"
	}
	return l.user.Email
}

func logoutPage(cr *Request) (*pb.WebloginResponse, error) {
	ctx := cr.ctx
	u := auth.GetUser(ctx)
	if u == nil {
		return nil, fmt.Errorf("cannot log you out because you are not yet logged in")
	}
	l := &LogoutStruct{user: u}
	res := NewWebloginResponse()
	addCookie(res, "Auth-Token", "")
	t, err := cr.renderTemplate(l, "loggedout")
	if err != nil {
		fmt.Printf("template error: %s\n", err)
		return nil, err
	}
	res.Body = t
	_, e := authManager.LogMeOut(ctx, &common.Void{})
	if e != nil {
		fmt.Printf("Failed to log out: %s\n", utils.ErrorString(e))
		return nil, e
	}
	return res, nil
}
