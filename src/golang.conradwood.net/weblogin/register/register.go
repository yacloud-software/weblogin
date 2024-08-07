package register

import (
	"context"
	"flag"
	"fmt"
	au "golang.conradwood.net/apis/auth"
	"golang.conradwood.net/apis/h2gproxy"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/weblogin/activitylog"
	"golang.conradwood.net/weblogin/common"
	"golang.conradwood.net/weblogin/requesttracker"
	"golang.conradwood.net/weblogin/web"
	"html/template"
	"net/mail"
	"net/url"
	"time"
)

var (
	CAPTCHA_BYPASS = flag.String("captcha_bypass", "", "if set, this header will allow endusers to bypass the captcha. Must be kept secret and long, but maybe useful for probers")
)

type RegisterRequest struct {
	SSOHost       string
	Host          string
	Error         string
	userid        string
	Password1     string
	Password2     string
	SiteKey       string
	FirstName     string
	LastName      string
	UserExists    bool
	VError        string // used in verify registration email
	VReg          string // the email link
	magic         string
	state         *pb.State
	RegisterState *pb.RegisterState
	logger        *activitylog.Logger
	cr            *requesttracker.Request
}

func (rr *RegisterRequest) Year() string {
	t := time.Now().Year()
	return fmt.Sprintf("%d", t)
}

func (rr *RegisterRequest) GetQueryValue(key string) string {
	return common.State2URLValues(rr.state)[key]
}

func (rr *RegisterRequest) TargetURL() string {
	return common.State2URL(rr.state)
}

func (rr *RegisterRequest) ReferrerHost() string {
	if rr.state != nil {
		return rr.state.TriggerHost
	}
	if rr.RegisterState != nil {
		return rr.RegisterState.Host
	}
	fmt.Printf("WARNING - no referrer host!!\n")
	return ""

}
func (rr *RegisterRequest) Username() string {
	return ""
}
func (rr *RegisterRequest) StateQuery() template.HTMLAttr {
	return template.HTMLAttr("?" + common.WEBLOGIN_STATE + "=" + rr.magic)
}
func (rr *RegisterRequest) Heading() string {
	return "Register Account"
}
func Registration(ctx context.Context, req *pb.WebloginRequest, cr *requesttracker.Request) (*pb.WebloginResponse, error) {
	var err error
	if !web.AllowRegister() {
		fmt.Printf("Attempt to register\n")
		return nil, errors.NotFound(ctx, "not found")
	}
	logger := &activitylog.Logger{
		IP: req.Peer,
	}
	res := &pb.WebloginResponse{}
	rr := &RegisterRequest{
		SSOHost: web.SSOHost(),
		SiteKey: web.CaptchaKey(),
		cr:      cr,
	}
	rr.logger = logger
	w := web.NewWebRequest(ctx, req)
	if w.GetPara(common.WEBLOGIN_STATE) != "" {
		rr.magic = w.GetPara(common.WEBLOGIN_STATE)
		rr.state, err = common.ParseMagic(ctx, rr.magic)
		if err != nil {
			return nil, err
		}
	}
	if rr.state == nil {
		rr.state = &pb.State{}
	}
	fmt.Printf("[registration] coming from host \"%s\"\n", rr.state.TriggerHost)

	rr.Host = w.GetPara("host")
	if rr.Host == "" {
		rr.Host = web.SSOHost()
	}
	var b []byte

	if w.GetPara("v_reg") != "" || w.GetPara("form_submit_Ohg5quei4no2gZZZrgeserg") != "" {
		logger.Log(ctx, fmt.Sprintf("Registration email verification request from host \"%s\"", rr.state.TriggerHost))
		// this is a link from an email
		response, err := rr.VerifyEmail(w)
		cr.SetError(err)
		cr.RegistrationEmailVerified()
		return response, err
	}

	// the register stuff needs redirecting
	if w.GetHost() != web.SSOHost() {
		// first of all - registration ONLY works on the sso host
		u := "?host=" + url.QueryEscape(req.Host)
		res.RedirectTo = "https://" + web.SSOHost() + "/weblogin/register" + u
		return res, nil
	}

	if w.GetPara("form_submit_Ohg5quei4no2grgeserg") != "" || w.GetPara("email") != "" {
		logger.Log(ctx, fmt.Sprintf("Registration email from submitted from host \"%s\" for email \"%s\"", rr.state.TriggerHost, w.GetPara("email")))
		b, err = rr.register1_submitted(ctx, logger, w)
		cr.SetError(err)
		cr.RegistrationSubmitted()
	} else {
		// default - no submission
		b, err = w.Render("register1", rr)
		cr.SetError(err)
		cr.RegistrationRendered()
	}
	if err != nil {
		return nil, err
	}
	res.Body = b

	return res, nil
}

// user is on the first form entering his/her email
func (rr *RegisterRequest) register1_submitted(ctx context.Context, logger *activitylog.Logger, w *web.WebRequest) ([]byte, error) {
	e := w.GetPara("email")
	fmt.Printf("Register email submitted (%s)\n", e)
	rr.SetEmail(e)
	logger.Email = e
	notvalid := ""
	var c bool
	var err error
	if w.GetPara("captcha_bypass") == *CAPTCHA_BYPASS {
		c = true
		err = nil
	} else {
		c, err = w.VerifyCaptcha()
	}
	if err != nil {
		notvalid = fmt.Sprintf("%s", err)
		w.BadIP(40)
	} else if !c {
		notvalid = "google captcha did not verify you as a user. try again please"
		w.BadIP(90)
	} else if len(e) == 0 {
		notvalid = "email address is mandatory"
	} else if len(e) < 3 {
		notvalid = "email address needs to be at least 3 characters"
	}
	if notvalid == "" {
		ad, err := mail.ParseAddress(e)
		if err != nil {
			notvalid = fmt.Sprintf("%s", err)
		}
		if ad != nil {
			rr.SetEmail(ad.Address)
		}
	}

	// on error, re-render register1
	if notvalid != "" {
		logger.Log(ctx, fmt.Sprintf("Registration request for host \"%s\" invalid: %s", rr.state.TriggerHost, notvalid))
		rr.Error = notvalid
		return w.Render("register1", rr)
	}
	logger.Log(ctx, fmt.Sprintf("Registration request for host \"%s\" successful (sending email)", rr.state.TriggerHost))

	// now send email
	err = rr.send_email(w)
	if err != nil {
		rr.cr.SetError(err)
		rr.Error = fmt.Sprintf("%s", err)
	}
	rr.cr.RegistrationEmailSent()
	return w.Render("register2", rr)
}

// user clicked on link in email
func (rr *RegisterRequest) VerifyEmail(w *web.WebRequest) (*pb.WebloginResponse, error) {
	var err error
	var ctx context.Context
	var user *au.User
	var b []byte
	res := &pb.WebloginResponse{}
	p := w.GetPara("v_reg")
	rr.VReg = p
	rs, err := decode_email_link(p)
	if err != nil {
		w.BadIP(60)
		goto error
	}
	rr.state = &pb.State{
		TriggerHost: rs.Host,
	}
	ctx = authremote.Context()
	rr.logger.Email = rs.Email
	rr.logger.Log(ctx, "user clicked on email link")
	rr.magic = rs.Magic
	rr.RegisterState = rs
	rr.Host = rs.Host
	rr.SetEmail(rs.Email)
	rr.Password1 = w.GetPara("password1")
	rr.Password2 = w.GetPara("password2")
	rr.FirstName = w.GetPara("firstname")
	rr.LastName = w.GetPara("lastname")
	rr.UserExists = false
	// it is possible that user re-registered (e.g. already exists)

	user, err = authremote.GetAuthManagerClient().GetUserByEmail(ctx, &au.ByEmailRequest{Email: rr.GetEmail()})
	if err != nil {
		fmt.Printf("error checking user \"%s\": %s\n", rr.GetEmail(), utils.ErrorString(err))
	} else if user != nil {
		rr.UserExists = true
		rr.userid = user.ID
	}
	if w.GetPara("form_submit_Ohg5quei4no2gZZZrgeserg") != "" {
		// it was submitted at least once
		err = rr.verify_submit_reg_form(w)
		if err != nil {
			goto verror
		}
		if rr.UserExists {
			err = rr.update_user(w)
		} else {
			err = rr.create_user(w)
		}
		if err != nil {
			goto verror
		}
		// user created or updated, log him/er in
		return rr.login(w)
	}
	goto render
verror:
	if err == nil {
		panic("error but no error")
	}
	rr.VError = fmt.Sprintf("%s", err)
	goto render
error:
	rr.Error = fmt.Sprintf("%s", err)
render:
	b, err = w.Render("verify_registration_email", rr)
	if err != nil {
		return nil, err
	}
	res.Body = b
	return res, nil
}

// are the parameters good enough to create a user?
func (rr *RegisterRequest) verify_submit_reg_form(w *web.WebRequest) error {
	if !rr.UserExists {
		// if user exists, don't bother with firstname/lastname
		f := w.GetPara("firstname")
		if len(f) == 0 {
			return common.UrgentErrorf("Firstname is mandatory")
		}
		if len(f) < 3 {
			return common.UrgentErrorf("Firstname is too short")
		}
		l := w.GetPara("lastname")
		if len(l) == 0 {
			return common.UrgentErrorf("Lastname is mandatory")
		}
		if len(l) < 3 {
			return common.UrgentErrorf("Lastname is too short")
		}
	}
	p1 := w.GetPara("password1")
	p2 := w.GetPara("password2")
	if len(p1) < 8 {
		return common.UrgentErrorf("Password must have a minimum length of 8 characters")
	}
	if p1 != p2 {
		return common.UrgentErrorf("Passwords do not match")
	}
	return nil
}

// update a users' password
func (rr *RegisterRequest) update_user(w *web.WebRequest) error {
	ctx := authremote.Context()
	fmt.Printf("updating password for user \"%s\" (%s)\n", rr.GetEmail(), rr.userid)
	upd := &au.ForceUpdatePasswordRequest{UserID: rr.userid, NewPassword: w.GetPara("password1")}
	_, err := authremote.GetAuthManagerClient().ForceUpdatePassword(ctx, upd)
	if err != nil {
		fmt.Printf("failed to update password for user \"%s\": %s\n", rr.GetEmail(), utils.ErrorString(err))
	}
	return err
}

// create a new uesr
func (rr *RegisterRequest) create_user(w *web.WebRequest) error {
	cr := &au.CreateUserRequest{
		Email:         rr.GetEmail(),
		FirstName:     rr.FirstName,
		LastName:      rr.LastName,
		Password:      w.GetPara("password1"),
		Abbrev:        utils.RandomString(7),
		EmailVerified: true,
	}
	ctx := authremote.Context()
	u, err := authremote.GetAuthManagerClient().CreateUser(ctx, cr)
	if err != nil {
		fmt.Printf("Unable to create user: %s\n", utils.ErrorString(err))
		return err
	}
	fmt.Printf("Created new user %s, ID=%s\n", u.Email, u.ID)
	return nil
}

// user has been created or has changed his/her password. tell him and log him in
func (rr *RegisterRequest) login(w *web.WebRequest) (*pb.WebloginResponse, error) {
	wr := &pb.WebloginResponse{}
	value := "foobar"
	exp := time.Now().Add(time.Duration(8) * time.Hour).Unix()
	hc := &h2gproxy.Cookie{Name: "Auth-Token", Value: value, Expiry: uint32(exp)}
	wr.Cookies = append(wr.Cookies, hc)
	b, err := w.Render("registration_complete", rr)
	if err != nil {
		return nil, err
	}
	wr.Body = b
	return wr, nil
}

func (rr *RegisterRequest) SetEmail(email string) {
	rr.cr.SetEmail(email)
}
func (rr *RegisterRequest) Email() string {
	return rr.GetEmail()
}
func (rr *RegisterRequest) GetEmail() string {
	return rr.cr.GetEmail()
}
