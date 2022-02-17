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
	"golang.conradwood.net/weblogin/web"
	"net/mail"
	"net/url"
	"time"
)

var (
	CAPTCHA_BYPASS = flag.String("captcha_bypass", "", "if set, this header will allow endusers to bypass the captcha. Must be kept secret and long, but maybe useful for probers")
)

type RegisterRequest struct {
	SSOHost    string
	Host       string
	Error      string
	Email      string
	userid     string
	Password1  string
	Password2  string
	SiteKey    string
	FirstName  string
	LastName   string
	UserExists bool
	VError     string // used in verify registration email
	VReg       string // the email link
}

func (rr *RegisterRequest) StateQuery() string {
	return ""
}
func Registration(ctx context.Context, req *pb.WebloginRequest) (*pb.WebloginResponse, error) {
	if !web.AllowRegister() {
		fmt.Printf("Attempt to register\n")
		return nil, errors.NotFound(ctx, "not found")
	}
	res := &pb.WebloginResponse{}
	rr := &RegisterRequest{
		SSOHost: web.SSOHost(),
		SiteKey: web.CaptchaKey(),
	}
	w := web.NewWebRequest(ctx, req)
	rr.Host = w.GetPara("host")
	if rr.Host == "" {
		rr.Host = web.SSOHost()
	}
	var b []byte
	var err error

	if w.GetPara("v_reg") != "" || w.GetPara("form_submit_Ohg5quei4no2gZZZrgeserg") != "" {
		// this is a link from an email
		return rr.VerifyEmail(w)
	}

	// the register stuff needs redirecting
	if w.GetHost() != web.SSOHost() {
		// first of all - registration ONLY works on the sso host
		u := "?host=" + url.QueryEscape(req.Host)
		res.RedirectTo = "https://" + web.SSOHost() + "/weblogin/register" + u
		return res, nil
	}

	if w.GetPara("form_submit_Ohg5quei4no2grgeserg") != "" || w.GetPara("email") != "" {
		b, err = rr.register1_submitted(w)
	} else {
		// default - no submission
		b, err = w.Render("register1", rr)
	}
	if err != nil {
		return nil, err
	}
	res.Body = b

	return res, nil
}

// user is on the first form entering his/her email
func (rr *RegisterRequest) register1_submitted(w *web.WebRequest) ([]byte, error) {
	fmt.Printf("Register email submitted\n")
	e := w.GetPara("email")
	rr.Email = e
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
		w.BadIP()
	} else if !c {
		notvalid = "google captcha did not verify you as a user. try again please"
		w.BadIP()
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
			rr.Email = ad.Address
		}
	}

	// on error, re-render register1
	if notvalid != "" {
		rr.Error = notvalid
		return w.Render("register1", rr)
	}
	// now send email
	err = rr.send_email(w)
	if err != nil {
		rr.Error = fmt.Sprintf("%s", err)
	}
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
		w.BadIP()
		goto error
	}
	rr.Host = rs.Host
	rr.Email = rs.Email
	rr.Password1 = w.GetPara("password1")
	rr.Password2 = w.GetPara("password2")
	rr.FirstName = w.GetPara("firstname")
	rr.LastName = w.GetPara("lastname")
	rr.UserExists = false
	// it is possible that user re-registered (e.g. already exists)
	ctx = authremote.Context()
	user, err = authremote.GetAuthManagerClient().GetUserByEmail(ctx, &au.ByEmailRequest{Email: rr.Email})
	if err != nil {
		fmt.Printf("error checking user \"%s\": %s\n", rr.Email, utils.ErrorString(err))
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
			return fmt.Errorf("Firstname is mandatory")
		}
		if len(f) < 3 {
			return fmt.Errorf("Firstname is too short")
		}
		l := w.GetPara("lastname")
		if len(l) == 0 {
			return fmt.Errorf("Lastname is mandatory")
		}
		if len(l) < 3 {
			return fmt.Errorf("Lastname is too short")
		}
	}
	p1 := w.GetPara("password1")
	p2 := w.GetPara("password2")
	if len(p1) < 8 {
		return fmt.Errorf("Password must have a minimum length of 8 characters")
	}
	if p1 != p2 {
		return fmt.Errorf("Passwords do not match")
	}
	return nil
}

// update a users' password
func (rr *RegisterRequest) update_user(w *web.WebRequest) error {
	ctx := authremote.Context()
	fmt.Printf("updating password for user \"%s\" (%s)\n", rr.Email, rr.userid)
	upd := &au.ForceUpdatePasswordRequest{UserID: rr.userid, NewPassword: w.GetPara("password1")}
	_, err := authremote.GetAuthManagerClient().ForceUpdatePassword(ctx, upd)
	if err != nil {
		fmt.Printf("failed to update password for user \"%s\": %s\n", rr.Email, utils.ErrorString(err))
	}
	return err
}

// create a new uesr
func (rr *RegisterRequest) create_user(w *web.WebRequest) error {
	cr := &au.CreateUserRequest{
		Email:         rr.Email,
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
