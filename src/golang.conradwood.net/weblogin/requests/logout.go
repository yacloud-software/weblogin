package requests

import (
	"fmt"
	apb "golang.conradwood.net/apis/auth"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/utils"
	"html/template"
)

type LogoutStruct struct {
	Msg   string
	user  *apb.User
	state *pb.State
}

func (l *LogoutStruct) Heading() string {
	return "Log out"
}
func (l *LogoutStruct) GetState() *pb.State {
	return l.state
}
func (l *LogoutStruct) ReferrerHost() string {
	if l.GetState() == nil {
		return ""
	}
	return l.GetState().TriggerHost

}
func (l *LogoutStruct) StateQuery() template.HTMLAttr {
	return template.HTMLAttr("")
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
	state, err := cr.getState(ctx)
	l := &LogoutStruct{user: u, state: state}
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
