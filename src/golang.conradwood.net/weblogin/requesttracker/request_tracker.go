package requesttracker

import (
	"context"
	"fmt"
	"golang.conradwood.net/apis/auth"
	"golang.conradwood.net/apis/h2gproxy"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/utils"
	al "golang.conradwood.net/weblogin/activitylog"
	"golang.conradwood.net/weblogin/opts"
	"golang.conradwood.net/weblogin/trackerif"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	BROWSERID_COOKIE = "yacloud_browserid"
)

type Request struct {
	req        *pb.WebloginRequest
	ctx        context.Context
	magic      string
	state      *pb.State
	email      string
	ip         string
	port       int
	logger     *al.Logger
	browserid  string
	tr         trackerif.TrackerIF
	user       *auth.User
	last_error error
}

func NewRequest(ctx context.Context, req *pb.WebloginRequest) *Request {
	res := &Request{ctx: ctx, req: req}
	if req.Submitted == nil {
		req.Submitted = make(map[string]string)
	}
	xip := net.ParseIP(req.Peer)
	if xip != nil {
		res.ip = req.Peer
	} else {
		ip, port, _ := net.SplitHostPort(req.Peer)
		res.ip = ip
		iport, _ := strconv.Atoi(port)
		res.port = iport
	}
	log_ip := res.ip
	if log_ip == "" {
		log_ip = req.Peer
	}
	res.logger = &al.Logger{IP: log_ip}
	for _, c := range req.Cookies {
		if c.Name == BROWSERID_COOKIE {
			res.browserid = c.Value
		}
	}
	res.tr = &StandardTracker{}
	return res
}
func (r *Request) SetEmail(email string) {
	r.email = email
}
func (r *Request) GetEmail() string {
	return r.email
}

func (r *Request) GetUser() *auth.User {
	return r.user
}
func (r *Request) SetUser(u *auth.User) {
	r.user = u
}
func (r *Request) Context() context.Context {
	return r.ctx
}
func (r *Request) Request() *pb.WebloginRequest {
	return r.req
}
func (r *Request) IP() string {
	return r.ip
}
func (l *Request) prefix() string {
	m := "unknown"
	if l != nil && l.req != nil {
		m = l.req.Method
	}
	s := fmt.Sprintf("[%s/%s] ", l.ip, m)
	return s
}
func (l *Request) Printf(format string, args ...interface{}) {
	fmt.Printf(l.prefix()+format, args...)
}
func (l *Request) Debugf(format string, args ...interface{}) {
	dodebug := opts.IsDebug()
	dodebug = dodebug || strings.HasPrefix(l.req.Peer, "[2001:8b0:1400:279b:5")
	dodebug = dodebug || strings.HasPrefix(l.req.Peer, "81.187.88.146")
	dodebug = dodebug || strings.HasPrefix(l.req.Peer, "81.187.202.194")
	dodebug = dodebug || strings.HasPrefix(l.req.Peer, "137.220.64.19")
	dodebug = dodebug || strings.HasPrefix(l.req.Peer, "2a01:4b00:ab0f:5100:5::5")
	if !dodebug {
		return
	}

	fmt.Printf(l.prefix()+format, args...)
}
func (cr *Request) PrintParas() {
	if !opts.IsDebug() {
		return
	}
	fmt.Printf("Path: https://%s/%s?%s\n", cr.req.Host, cr.req.Path, cr.req.Query)
	for k, v := range cr.req.Submitted {
		if len(v) > 60 {
			v = v[:60] + "..."
		}
		fmt.Printf("Parameter %s: %s\n", k, v)
	}
}
func (cr *Request) SetState(state *pb.State) {
	cr.state = state
}
func (cr *Request) GetState() *pb.State {
	return cr.state
}
func (cr *Request) SetMagic(magic string) {
	cr.magic = magic
}
func (cr *Request) GetMagic() string {
	return cr.magic
}

func (cr *Request) BrowserID() string {
	return cr.browserid
}
func (cr *Request) CookiesToSet() []*h2gproxy.Cookie {
	var res []*h2gproxy.Cookie
	if cr.BrowserID() == "" {
		s := utils.RandomString(128)
		e := uint32(time.Now().Add(time.Duration(365*24*5) * time.Hour).Unix())
		res = append(res, &h2gproxy.Cookie{Name: BROWSERID_COOKIE, Value: s, Expiry: e})
	}
	return res
}
func (cr *Request) UserAgent() string {
	return cr.req.UserAgent
}
func (cr *Request) SetError(err error) {
	if err == nil {
		return
	}
	cr.last_error = err
}
func (cr *Request) ForgotPasswordSent() {
	cr.request_log(pb.AuthAction_FORGOT_PASSWORD_SENT)
}
func (cr *Request) ResetPasswordSent() {
	cr.request_log(pb.AuthAction_RESET_PASSWORD_SENT)
}
func (cr *Request) LoggedOut() {
	cr.request_log(pb.AuthAction_LOGGED_OUT)
}
func (cr *Request) LoginPageSubmitted() {
	cr.request_log(pb.AuthAction_CREDENTIALS_SUBMITTED)
}
func (cr *Request) LoginPageRendered() {
	cr.request_log(pb.AuthAction_LOGIN_RENDERED)
}
func (cr *Request) UnspecifiedRequest() {
	cr.request_log(pb.AuthAction_UNDEFINED)
}
func (cr *Request) RegistrationSubmitted() {
	cr.request_log(pb.AuthAction_SIGNUP_SUBMITTED)
}
func (cr *Request) RegistrationEmailVerified() {
	cr.request_log(pb.AuthAction_SIGNUP_EMAILCLICKED)
}
func (cr *Request) RegistrationRendered() {
	cr.request_log(pb.AuthAction_SIGNUP_RENDERED)
}
func (cr *Request) RegistrationEmailSent() {
	cr.request_log(pb.AuthAction_SIGNUP_EMAILSENT)
}
func (cr *Request) TriggerURL() string {
	req := cr.Request()
	q := ""
	if req.Query != "" {
		q = q + "?" + req.Query
	}
	return req.Scheme + "://" + req.Host + "/" + req.Path + q
}
func (cr *Request) request_log(action pb.AuthAction) {
	em := cr.GetEmail()
	u := cr.GetUser()
	userid := ""
	if u != nil {
		userid = u.ID
		if em == "" {
			em = u.Email
		}
	}
	aar := &pb.AuthActivityRequest{
		Timestamp:      uint32(time.Now().Unix()),
		IP:             cr.IP(),
		PreviousUserID: "weblogin-notimplemented",
		UserID:         userid,
		Email:          em,
		URL:            cr.TriggerURL(),
	}
	cr.tr.LogActivity(action, aar)
	//	text := fmt.Sprintf("%v", action)
	//fmt.Printf("[REQUESTTRACKER "+text+"] %#v\n", aar)

}
