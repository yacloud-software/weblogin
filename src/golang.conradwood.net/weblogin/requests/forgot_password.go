package requests

import (
	"fmt"
	apb "golang.conradwood.net/apis/auth"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/weblogin/common"
	"golang.conradwood.net/weblogin/requesttracker"
	"html/template"
)

type ForgotStruct struct {
	Msg   string
	magic string
	user  *apb.User
	state *pb.State
	PW1   string
	PW2   string
}

func (rr *ForgotStruct) GetQueryValue(key string) string {
	return common.State2URLValues(rr.state)[key]
}

func (rr *ForgotStruct) TargetURL() string {
	return common.State2URL(rr.state)
}

func (l *ForgotStruct) Heading() string {
	return "Reset Password"
}

func (l *ForgotStruct) GetState() *pb.State {
	return l.state
}
func (l *ForgotStruct) ReferrerHost() string {
	if l.GetState() == nil {
		return ""
	}
	return l.GetState().TriggerHost

}
func (l *ForgotStruct) StateQuery() template.HTMLAttr {
	return template.HTMLAttr("?" + common.WEBLOGIN_STATE + "=" + l.magic)
}

// render l.state into some string
func (l *ForgotStruct) Weblogin_state_value() string {
	if l.magic == "" {
		panic("missing magic")
	}
	return l.magic
}
func (l *ForgotStruct) Weblogin_state_name() string {
	return common.WEBLOGIN_STATE
}

func (l *ForgotStruct) State() string {
	if l.magic == "" {
		panic("no magic")
	}
	return l.magic
}
func (l *ForgotStruct) Username() string {
	if l.user == nil {
		return "nobody"
	}
	return l.user.Email
}

func forgotpasswordPage(cr *requesttracker.Request) (*pb.WebloginResponse, error) {
	cr.Debugf("Forgotpassword request\n")
	ctx := cr.Context()
	req := cr.Request()
	if req.Submitted["email"] != "" {
		return forgotSent(cr)
	}
	// render template to type in an email address
	u := auth.GetUser(ctx)
	state, err := getState(ctx, cr)
	l := &ForgotStruct{user: u, state: state}
	res := NewWebloginResponse()
	addCookies(res, cr.CookiesToSet())
	t, err := renderTemplate(cr, l, "forgotv2")
	if err != nil {
		fmt.Printf("template error: %s\n", err)
		return nil, err
	}
	res.Body = t

	return res, nil
}
func forgotSent(cr *requesttracker.Request) (*pb.WebloginResponse, error) {
	req := cr.Request()
	ctx := cr.Context()
	email := req.Submitted["email"]
	_, err := authremote.GetAuthManagerClient().ResetPasswordEmail(ctx, &apb.ResetRequest{Email: email})
	if err != nil {
		l := &ForgotStruct{Msg: fmt.Sprintf("%s", err)}
		res := NewWebloginResponse()
		t, err := renderTemplate(cr, l, "forgotv2")
		if err != nil {
			fmt.Printf("template error: %s\n", err)
			return nil, err
		}
		res.Body = t
		return res, nil
	}
	l := &ForgotStruct{Msg: fmt.Sprintf("email sent to %s", email)}
	res := NewWebloginResponse()
	t, err := renderTemplate(cr, l, "forgotv2.1")
	if err != nil {
		fmt.Printf("template error: %s\n", err)
		return nil, err
	}
	res.Body = t
	return res, nil

}
func resetpasswordPage(cr *requesttracker.Request) (*pb.WebloginResponse, error) {
	ctx := cr.Context()
	req := cr.Request()
	pw1 := req.Submitted["password1"]
	pw2 := req.Submitted["password2"]
	if pw1 != "" && pw2 != "" && pw1 == pw2 {
		return resettingPage(cr)
	}
	apikey := req.Submitted["apikey"]
	if apikey == "" {
		return forgotpasswordPage(cr)
	}
	u, err := authremote.GetAuthClient().GetByToken(ctx, &apb.AuthenticateTokenRequest{Token: apikey})
	if err != nil {
		return nil, err
	}
	if !u.Valid {
		return nil, errors.AccessDenied(ctx, "user api key not valid. Old email?")
	}
	state := &pb.State{Token: apikey}
	magic := utils.RandomString(60)
	err = putMagic(cr, magic, state)
	if err != nil {
		return nil, err
	}

	l := &ForgotStruct{
		Msg:   "resetting password for " + auth.Description(u.User),
		magic: magic,
		PW1:   req.Submitted["password1"],
		PW2:   req.Submitted["password2"],
	}
	if pw1 != pw2 {
		l.Msg = "Passwords do not match"
	}
	if pw1 == "" {
		l.Msg = "Please enter passwords"
	}
	res := NewWebloginResponse()
	addCookies(res, cr.CookiesToSet())
	t, err := renderTemplate(cr, l, "forgotv2.2")
	if err != nil {
		fmt.Printf("template error: %s\n", err)
		return nil, err
	}
	res.Body = t
	return res, nil
}
func resettingPage(cr *requesttracker.Request) (*pb.WebloginResponse, error) {
	req := cr.Request()
	ctx := cr.Context()
	s := req.Submitted[common.WEBLOGIN_STATE]
	state, err := getMagic(ctx, cr, s)
	if err != nil {
		return nil, err
	}
	u, err := authremote.GetAuthClient().GetByToken(ctx, &apb.AuthenticateTokenRequest{Token: state.Token})
	if err != nil {
		return nil, err
	}
	if !u.Valid {
		return nil, errors.AccessDenied(ctx, "user api key not valid. Old email?")
	}
	ctx, err = authremote.ContextForUserID(u.User.ID)

	pw := req.Submitted["password1"]
	pw2 := req.Submitted["password2"]
	l := &ForgotStruct{Msg: "Password reset for user " + auth.Description(u.User)}
	if err != nil || u == nil || !u.Valid {
		l.Msg = "Temporary login via URL failed. please try again"
	} else if pw != pw2 {
		l.Msg = "Passwords do not match"
	} else if len(pw) < 6 {
		l.Msg = "Password too short"
	} else {
		_, err = authremote.GetAuthManagerClient().UpdateMyPassword(ctx, &apb.UpdatePasswordRequest{NewPassword: pw})
	}
	if err != nil {
		l.Msg = "Failed to change password"
		fmt.Printf("Error changing password: %s\n", err)
	}
	res := NewWebloginResponse()
	t, err := renderTemplate(cr, l, "forgotv2.3")
	if err != nil {
		fmt.Printf("template error: %s\n", err)
		return nil, err
	}
	res.Body = t
	return res, nil

}
