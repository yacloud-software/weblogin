package requests

import (
	"context"
	"fmt"
	au "golang.conradwood.net/apis/auth"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/prometheus"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/weblogin/common"
	"golang.conradwood.net/weblogin/requesttracker"
	"golang.conradwood.net/weblogin/web"
	//	"golang.conradwood.net/go-easyops/utils"
	al "golang.conradwood.net/weblogin/activitylog"
	"golang.conradwood.net/weblogin/register"
	sm "golang.yacloud.eu/apis/sessionmanager"
	"google.golang.org/grpc"
	"strings"
	"time"
)

func (r *RequestHandler) StartGRPC(port int) error {
	sd := server.NewServerDef()
	sd.SetPort(port)
	sd.Register = func(server *grpc.Server) error {
		pb.RegisterWebloginServer(server, r)
		return nil
	}
	err := server.ServerStartup(sd)
	if err != nil {
		fmt.Printf("failed to start server: %s\n", err)
	}
	return err
}

// debug function to generate email for a user who's registering
func (w *RequestHandler) CreateRegisterEmail(ctx context.Context, cr *pb.RegisterState) (*pb.Email, error) {
	return register.CreateRegisterEmail(ctx, cr)
}

func (w *RequestHandler) IsBasicAuthValid(ctx context.Context, cr *pb.BasicAuthRequest) (*pb.AuthResponse, error) {
	return nil, errors.NotImplemented(ctx, "weblogin-v2 incomplete")
}

// we end up in here if a well-known url for weblogin is requested, that is "/weblogin/"
func (w *RequestHandler) ServeHTML(ctx context.Context, req *pb.WebloginRequest) (*pb.WebloginResponse, error) {
	fmt.Printf("Requested path \"%s\"\n", req.Path)
	cr := requesttracker.NewRequest(ctx, req)
	requestCounter.With(prometheus.Labels{"counter": "total"}).Inc()
	CountURL(cr)
	e := IsDosing(cr)
	if e != nil {
		requestCounter.With(prometheus.Labels{"counter": "fail-dos"}).Inc()
		cr.SetError(e)
		cr.UnspecifiedRequest()
		cr.Debugf("not serving path \"%s\": %s\n", req.Path, e)
		return ServeError(ctx, req, e)
	}
	cr.Debugf("weblogin.ServeHTML(), serving %s\n", req.Path)
	xs := strings.Split(req.Path, "/")
	if len(xs) >= 2 {
		fname := xs[len(xs)-1]
		if xs[len(xs)-2] == "assets" {
			res, err := ServeAsset(ctx, cr, fname)
			if err != nil {
				requestCounter.With(prometheus.Labels{"counter": "fail-asset"}).Inc()
			}
			return res, err
		}
	}

	wl, err := w.ServeHTMLWithError(ctx, req)
	cr.SetError(err)
	cr.UnspecifiedRequest()
	if err != nil {
		requestCounter.With(prometheus.Labels{"counter": "fail"}).Inc()
		u := auth.GetUser(ctx)
		cr.Debugf("weblogin.ServeHTML: error serving https://%s/%s?%s for user %s: %s\n", req.Host, req.Path, req.Query, auth.Description(u), utils.ErrorString(err))
		wr, err2 := ServeError(ctx, req, err)
		if err2 != nil {
			// serveerror failed, we cannot render this ;(
			fmt.Printf("ServeError failed: %s\n", err)
			return wr, err2
		}

		if wr != nil {
			// if error, make sure we don't mislead h2groxy
			wr.Authenticated = false
		}
		return wr, nil
	}

	return wl, nil
}
func (w *RequestHandler) ServeHTMLWithError(ctx context.Context, req *pb.WebloginRequest) (*pb.WebloginResponse, error) {
	cr := requesttracker.NewRequest(ctx, req)
	u := auth.GetUser(ctx)
	cr.SetUser(u)
	host := req.Host
	q := cr.Request().Query
	if len(q) > 50 {
		q = q[:50] + "..."
	}
	cr.Debugf("weblogin.ServeHTMLWithError: serving https://%s/%s?%s for user %s\n", req.Host, req.Path, q, auth.Description(u))
	cr.PrintParas()
	if strings.HasSuffix(req.Path, "/oauth") { // oauththing
		if u != nil {
			return googleOAuth(cr)
		} else {
			// we have to generate some magic here for the rest to work
			err := generateMagicIfNecessary(cr)
			if err != nil {
				return nil, err
			}
		}
	}
	initMagic(ctx, req, cr)

	str, err := serveThemes(ctx, cr)
	if err != nil {
		return nil, err
	}
	if str != nil {
		return str, nil
	}

	if strings.Contains(req.Path, "/register") {
		return register.Registration(ctx, req, cr)
	}

	if strings.HasSuffix(req.Path, "/setcookie") { // after authentication
		return setCookiePage(cr)
	}
	if strings.HasSuffix(req.Path, "/preauthimg") { // set a cookie with a temporary session, which we only link to an auth cookie AFTER authentication completed
		return preAuthCookie(cr)
	}
	if strings.HasSuffix(req.Path, "/logout") { // user clicked "logout"
		res, err := logoutPage(cr)
		cr.SetError(err)
		cr.LoggedOut()
		return res, err
	}
	if strings.HasSuffix(req.Path, "/forgotPassword") { // user clicked "forgot password"
		res, err := forgotpasswordPage(cr)
		cr.SetError(err)
		cr.ForgotPasswordSent()
		return res, err
	}
	if strings.HasSuffix(req.Path, "/forgotpw") { // user clicked on link in reset password email
		res, err := resetpasswordPage(cr)
		cr.SetError(err)
		cr.ResetPasswordSent()
		return res, err
	}
	// .../login
	msg := ""
	if host != web.SSOHost() {
		msg = fmt.Sprintf("This page must be loaded at %s (not \"%s\"). (path=%s)", web.SSOHost(), host, req.Path)
		fmt.Println(msg)
	}

	if strings.HasSuffix(req.Path, "/needsession") { // user clicked on link in reset password email
		res, err := needSessionPage(cr)
		cr.SetError(err)
		cr.SessionSet()
		return res, err
	}

	paras := req.Submitted

	if paras["email"] != "" {
		s := ""
		cr.SetEmail(paras["email"]) // so it shows up in the activity log with an email, even if it does not become a user through authentication
		if cr.GetState() != nil {
			s = cr.GetState().TriggerHost
		}
		logger := &al.Logger{
			IP:          cr.Request().Peer,
			TriggerHost: s,
			Email:       paras["email"],
			BrowserID:   cr.BrowserID(),
			UserAgent:   cr.UserAgent(),
		}
		r, user, err := processLogin(cr)
		cr.SetUser(user)
		cr.SetError(err)
		if err != nil {
			cr.LoginPageSubmitted()
			werr, castable := err.(*common.WError)
			if (castable && werr.Urgent) || (!castable) {
				logger.Log(ctx, fmt.Sprintf("login failed: %s", err))
			}
		} else {
			login_success(ctx, user, logger)
			cr.LoginPageSubmitted()
		}
		return r, err
	}
	// any other path: serve a login page
	res, err := createLoginPage(cr)
	if err != nil {
		cr.SetError(err)
		cr.LoginPageRendered()
		cr.Debugf("createLoginPage() failed: %s\n", err)
	}
	return res, err

}

func (w *RequestHandler) GetVerifyEmail(ctx context.Context, req *pb.WebloginRequest) (*pb.EmailPageResponse, error) {
	cr := requesttracker.NewRequest(ctx, req)
	CountURL(cr)
	e := IsDosing(cr)
	if e != nil {
		return nil, e
	}

	u := auth.GetUser(ctx)
	res := &pb.EmailPageResponse{HTML: "Logged in", Verified: true, User: u}
	if !u.EmailVerified {
		cr.Debugf("Non-Verified User \"%s\" attempted to browse %s/%s?%s\n", auth.Description(u), req.Host, req.Path, req.Query)
		res = &pb.EmailPageResponse{Verified: false, User: u, HTML: "account locked - login and verify your account please"}
	}
	res.Headers = map[string]string{"weblogin": "true"}
	return res, nil
}
func (w *RequestHandler) SaveState(ctx context.Context, req *pb.WebloginRequest) (*pb.StateResponse, error) {
	cr := requesttracker.NewRequest(ctx, req)
	CountURL(cr)
	e := IsDosing(cr)
	if e != nil {
		return nil, e
	}

	m, _, err := createState(cr)
	if err != nil {
		return nil, err
	}
	res := &pb.StateResponse{YacloudWebloginState: m, URLStateName: common.WEBLOGIN_STATE}
	fmt.Printf("created new state (by rpc): %s\n", m)
	return res, nil
}

// we end up here if h2gproxy calls a backend which then says 'please authenticate me'
func (w *RequestHandler) GetLoginPage(ctx context.Context, req *pb.WebloginRequest) (*pb.WebloginResponse, error) {
	cr := requesttracker.NewRequest(ctx, req)
	CountURL(cr)
	e := IsDosing(cr)
	if e != nil {
		cr.Debugf("GetLoginPage(): determined peer is dos'ing. not serving path \"%s\": %s\n", req.Path, e)
		return nil, e
	}

	host := req.Host
	u := auth.GetUser(ctx)
	cr.Debugf("(rpc) host \"%s\" login request Peer=%s for user %s\n", host, req.Peer, auth.Description(u))
	res, err := createLoginPage(cr)
	if err != nil {
		cr.Debugf("createLoginPage() failed: %s\n", err)
	}
	return res, err

}
func (w *RequestHandler) VerifyURL(ctx context.Context, req *pb.WebloginRequest) (*pb.WebloginResponse, error) {
	cr := requesttracker.NewRequest(ctx, req)
	CountURL(cr)
	e := IsDosing(cr)
	if e != nil {
		return nil, e
	}

	cr.Debugf("verifying url\n")
	wls := req.Submitted["weblogin"]
	if wls == "" {
		return nil, errors.InvalidArgs(ctx, "missing weblogin", "missing weblogin parameter")
	}
	state, err := getMagic(ctx, cr, wls)
	if err != nil {
		return nil, err
	}
	res := NewWebloginResponse()
	if state.Token == "" {
		return res, nil
	}
	u, err := UserByToken(ctx, state.Token)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return res, nil
	}
	res.User = u
	addCookie(res, "Auth-Token", state.Token)
	return res, nil
}

func NewWebloginResponse() *pb.WebloginResponse {
	res := &pb.WebloginResponse{}
	addStandardHeader(res)
	return res
}
func addStandardHeader(r *pb.WebloginResponse) {
	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}
	r.Headers["weblogin"] = "true"
}

func initMagic(ctx context.Context, req *pb.WebloginRequest, cr *requesttracker.Request) {
	cr.SetMagic(cr.Request().Submitted[common.WEBLOGIN_STATE])
	if cr.GetMagic() != "" {
		xstate, _ := common.ParseMagic(ctx, cr.GetMagic())
		cr.SetState(xstate)
		if xstate != nil && xstate.Token != "" && cr.GetUser() == nil {
			u, err := UserByToken(ctx, xstate.Token)
			if err == nil && u != nil {
				cr.SetUser(u)
			}
		}
		return
	}
	s := cr.Request().Submitted["v_reg"] // email link
	if s != "" {
		p, err := register.DecodeEmailLink(ctx, s)
		if err != nil {
			fmt.Printf("failed to decode email link: %s\n", err)
		}
		if p != nil {
			fmt.Printf("[regverify] user verified link. state: %#v\n", p)
			cr.SetState(&pb.State{})
		}
	}
}

func login_success(ctx context.Context, user *au.User, logger *al.Logger) {
	sr := &sm.NewSessionRequest{
		IPAddress:   logger.IP,
		BrowserID:   logger.BrowserID,
		UserAgent:   logger.UserAgent,
		UserID:      user.ID,
		Useremail:   user.Email,
		TriggerHost: logger.TriggerHost,
		Username:    fmt.Sprintf("%s %s", user.FirstName, user.LastName),
	}
	sb, err := sm.GetSessionManagerClient().NewSession(ctx, sr)
	if err != nil {
		fmt.Printf("Failed to get session: %s\n", utils.ErrorString(err))
		logger.Log(ctx, fmt.Sprintf("login suceeded - session failed (%s)", err))
	} else {
		fmt.Printf("Session created: %#v\n", sb)
		if noteworthy_login(ctx, user, sb) {
			logger.Log(ctx, fmt.Sprintf("login suceeded (to %s)", logger.TriggerHost))
		}
	}
}

func noteworthy_login(ctx context.Context, user *au.User, sr *sm.SessionResponse) bool {
	if user.ID != "1" && user.ID != "7" {
		return true
	}
	if sr.NewDevice {
		return true
	}
	t := time.Unix(int64(sr.LastSessionTimestamp), 0)
	if time.Since(t) > time.Duration(25)*time.Hour {
		return true
	}
	return false
}
