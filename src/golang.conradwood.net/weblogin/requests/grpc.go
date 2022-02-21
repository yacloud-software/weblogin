package requests

import (
	"context"
	"fmt"
	au "golang.conradwood.net/apis/auth"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/weblogin/common"
	"golang.conradwood.net/weblogin/web"
	//	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/weblogin/register"
	"google.golang.org/grpc"
	"strings"
)

func (r *RequestHandler) StartGRPC(port int) error {
	sd := server.NewServerDef()
	sd.Port = port
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
	cr := NewRequest(ctx, req)
	CountURL(cr)
	e := IsDosing(cr)
	if e != nil {
		cr.Debugf("not serving path \"%s\": %s\n", req.Path, e)
		return ServeError(ctx, req, e)
	}
	if *debug {
		cr.Debugf("weblogin.ServeHTML(), serving %s\n", req.Path)
	}
	wl, err := w.ServeHTMLWithError(ctx, req)
	if err != nil {
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
	cr := NewRequest(ctx, req)
	u := auth.GetUser(ctx)
	host := req.Host
	q := cr.req.Query
	if len(q) > 50 {
		q = q[:50] + "..."
	}
	cr.Debugf("weblogin.ServeHTMLWithError: serving https://%s/%s?%s for user %s\n", req.Host, req.Path, q, auth.Description(u))
	cr.printParas()
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
		return register.Registration(ctx, req)
	}

	if strings.HasSuffix(req.Path, "/setcookie") { // after authentication
		return setCookiePage(cr)
	}
	if strings.HasSuffix(req.Path, "/preauthimg") { // set a cookie with a temporary session, which we only link to an auth cookie AFTER authentication completed
		return preAuthCookie(cr)
	}
	if strings.HasSuffix(req.Path, "/logout") { // user clicked "logout"
		return logoutPage(cr)
	}
	if strings.HasSuffix(req.Path, "/forgotPassword") { // user clicked "forgot password"
		return forgotpasswordPage(cr)
	}
	if strings.HasSuffix(req.Path, "/forgotpw") { // user clicked on link in reset password email
		return resetpasswordPage(cr)
	}
	// .../login
	msg := ""
	if host != web.SSOHost() {
		msg = fmt.Sprintf("This page must be loaded at %s (not \"%s\"). (path=%s)", web.SSOHost(), host, req.Path)
		fmt.Println(msg)
	}
	paras := req.Submitted

	if paras["email"] != "" {
		return processLogin(cr)
	}
	// any other path: serve a login page
	res, err := cr.createLoginPage()
	if err != nil {
		cr.Debugf("createLoginPage() failed: %s\n", err)
	}
	return res, err

}
func (w *RequestHandler) GetVerifyEmail(ctx context.Context, req *pb.WebloginRequest) (*pb.EmailPageResponse, error) {
	cr := NewRequest(ctx, req)
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
	cr := NewRequest(ctx, req)
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
	cr := NewRequest(ctx, req)
	CountURL(cr)
	e := IsDosing(cr)
	if e != nil {
		cr.Debugf("GetLoginPage(): determined peer is dos'ing. not serving path \"%s\": %s\n", req.Path, e)
		return nil, e
	}

	host := req.Host
	u := auth.GetUser(ctx)
	cr.Debugf("(rpc) host \"%s\" login request Peer=%s for user %s\n", host, req.Peer, auth.Description(u))
	res, err := cr.createLoginPage()
	if err != nil {
		cr.Debugf("createLoginPage() failed: %s\n", err)
	}
	return res, err

}
func (w *RequestHandler) VerifyURL(ctx context.Context, req *pb.WebloginRequest) (*pb.WebloginResponse, error) {
	cr := NewRequest(ctx, req)
	CountURL(cr)
	e := IsDosing(cr)
	if e != nil {
		return nil, e
	}

	cr.Debugf("verifying url")
	wls := req.Submitted["weblogin"]
	if wls == "" {
		return nil, errors.InvalidArgs(ctx, "missing weblogin", "missing weblogin parameter")
	}
	state, err := cr.getMagic(ctx, wls)
	if err != nil {
		return nil, err
	}
	res := NewWebloginResponse()
	if state.Token == "" {
		return res, nil
	}
	apr := &au.AuthenticateTokenRequest{Token: state.Token}
	u, err := authremote.GetAuthClient().GetByToken(ctx, apr)
	if err != nil {
		return nil, err
	}
	if !u.Valid {
		return res, nil
	}
	res.User = u.User
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

func initMagic(ctx context.Context, req *pb.WebloginRequest, cr *Request) {
	cr.magic = cr.req.Submitted[common.WEBLOGIN_STATE]
	if cr.magic != "" {
		cr.state, _ = common.ParseMagic(ctx, cr.magic)
		return
	}
	s := cr.req.Submitted["v_reg"] // email link
	if s != "" {
		p, err := register.DecodeEmailLink(ctx, s)
		if err != nil {
			fmt.Printf("failed to decode email link: %s\n", err)
		}
		if p != nil {
			fmt.Printf("[regverify] user verified link. state: %#v\n", p)
			cr.state = &pb.State{}
		}
	}
}
