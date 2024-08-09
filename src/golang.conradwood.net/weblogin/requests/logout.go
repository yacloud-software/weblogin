package requests

import (
	"fmt"
	apb "golang.conradwood.net/apis/auth"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/utils"
	cm "golang.conradwood.net/weblogin/common"
	"golang.conradwood.net/weblogin/requesttracker"
	"golang.yacloud.eu/apis/sessionmanager"
	"html/template"
	"time"
)

type LogoutStruct struct {
	Msg   string
	user  *apb.User
	state *pb.State
}

func (rr *LogoutStruct) Year() string {
	t := time.Now().Year()
	return fmt.Sprintf("%d", t)
}

func (rr *LogoutStruct) GetQueryValue(key string) string {
	return cm.State2URLValues(rr.state)[key]
}

func (rr *LogoutStruct) TargetURL() string {
	return cm.State2URL(rr.state)
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

func logoutPage(cr *requesttracker.Request) (*pb.WebloginResponse, error) {
	ctx := cr.Context()
	u := auth.GetUser(ctx)
	if u == nil {
		u = cr.GetUser()
		if u == nil {
			return nil, cm.Errorf("cannot log you out because you are not yet logged in")
		}
	}
	state, err := getState(ctx, cr)
	l := &LogoutStruct{user: u, state: state}
	res := NewWebloginResponse()
	addCookies(res, cr.CookiesToSet())
	addCookie(res, "Auth-Token", "", time.Duration(10)*time.Second) // clear cookie soonish
	t, err := renderTemplate(cr, l, "loggedout")
	if err != nil {
		fmt.Printf("template error: %s\n", err)
		return nil, err
	}
	res.Body = t

	var a_err error
	var s_err error
	_, a_err = authManager.LogMeOut(ctx, &common.Void{})
	if a_err != nil {
		fmt.Printf("Failed to log out: %s\n", utils.ErrorString(a_err))
	}

	token := cr.GetAuthCookie()
	if token == "" {
		fmt.Printf("no auth cookie to log out with\n")
	} else {
		st := &sessionmanager.SessionToken{Token: token}
		_, s_err = sessionmanager.GetSessionManagerClient().DisassociateUserFromSession(ctx, st)
		if s_err != nil {
			fmt.Printf("failed to log out of session: %s\n", utils.ErrorString(s_err))
		}
	}
	if a_err != nil && s_err != nil {
		return nil, a_err
	}
	return res, nil
}
