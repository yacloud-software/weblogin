package requests

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	au "golang.conradwood.net/apis/auth"
	"golang.conradwood.net/apis/h2gproxy"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/http"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/weblogin/common"
	"golang.conradwood.net/weblogin/register"
	"golang.conradwood.net/weblogin/requesttracker"
	"golang.conradwood.net/weblogin/web"
	"golang.yacloud.eu/apis/sessionmanager"
	"html/template"
	"net/url"
	"strings"
	"time"
)

var (
	issue_session_cookie_instead_of_auth = flag.Bool("issue_session_cookie", true, "if true issue a session token instead of an auth token")
	check_captcha_on_login               = flag.Bool("check_captcha_on_login", true, "if true also checks captcha on login")
	captcha_bypass_on_login              = flag.Bool("captcha_bypass_on_login", false, "if true, allows captcha bypass on login as well as register")

// dur_secs                             = flag.Int("session_lifetime", 12*60*60, "session lifetime in `seconds`")
// Cookie_livetime                      *int
)

// implements common.Template_data
type loginrender struct {
	Msg                  string
	weblogin_state_value *pb.State
	magic                string
	ImageURLs            []string
	SiteKey              string
}

func (rr *loginrender) Year() string {
	t := time.Now().Year()
	return fmt.Sprintf("%d", t)
}
func (rr *loginrender) GetQueryValue(key string) string {
	return common.State2URLValues(rr.weblogin_state_value)[key]
}

func (rr *loginrender) TargetURL() string {
	return common.State2URL(rr.weblogin_state_value)
}

func (l *loginrender) Heading() string {
	return "Log In"
}
func (l *loginrender) GetState() *pb.State {
	return l.weblogin_state_value
}
func (l *loginrender) ReferrerHost() string {
	if l.GetState() == nil {
		return ""
	}
	return l.GetState().TriggerHost

}
func (l *loginrender) StateQuery() template.HTMLAttr {
	return template.HTMLAttr("?" + common.WEBLOGIN_STATE + "=" + l.magic)
}

// render l.state into some string
func (l *loginrender) Weblogin_state_value() string {
	if l.magic == "" {
		panic("missing magic")
	}
	return l.magic
}
func (l *loginrender) Weblogin_state_name() string {
	return common.WEBLOGIN_STATE
}
func (l *loginrender) RegistrationEnabled() bool {
	return web.AllowRegister()
}
func (l *loginrender) Username() string {
	return ""
}

// this redirects to sso.yacloud.eu if it is not already and then serves the form
func createLoginPage(cr *requesttracker.Request) (*pb.WebloginResponse, error) {
	cr.Debugf("createLoginPage()\n")
	ctx := cr.Context()
	req := cr.Request()
	if req.Host != web.SSOHost() {
		cr.Debugf("going from host %s to sso.yacloud.eu...\n", req.Host)
		magic, _, err := createState(cr)
		if err != nil {
			return nil, err
		}
		res := NewWebloginResponse()
		res.Body = []byte("redirecting to sso")
		res.RedirectTo = fmt.Sprintf("https://%s/weblogin/login?"+common.WEBLOGIN_STATE+"=%s", web.SSOHost(), magic)
		return res, nil
	}
	//cr.logger.Log(ctx, "Presenting loginpage")
	cr.Debugf("Presenting loginpage...\n")
	// now the magic:
	// are we authenticated already? if so, skip asking and move on
	u := auth.GetUser(ctx)
	if u != nil {
		return skipAuth(cr)
	}
	res := NewWebloginResponse()
	submittedParameters := req.Submitted
	magic := submittedParameters[common.WEBLOGIN_STATE]
	state, err := getMagic(ctx, cr, magic)
	if err != nil {
		cr.Debugf("Whilst presenting login page, tried to recreate state, but error: %s\n", err)
		return nil, err
	}

	l := &loginrender{magic: magic, SiteKey: web.CaptchaKey()}
	l.weblogin_state_value = state
	if l.weblogin_state_value == nil {
		cr.Debugf("[servehtml] WARNING made up new state\n")
		l.weblogin_state_value = &pb.State{TriggerHost: "www.yacloud.eu", TriggerPath: "/weblogin/loggedin"}
	}
	l.ImageURLs, err = createDomainLogins(cr)
	if err != nil {
		cr.Debugf("[servehtml] failed to create domain logins: %s\n", err)
	}
	cr.Debugf("[servehtml] State: %#v\n", l.weblogin_state_value)
	t, err := renderTemplate(cr, l, "loginv2")
	if err != nil {
		cr.Debugf("template error: %s\n", err)
		return nil, err
	}
	res.Body = t
	return res, nil
}

// we end up here if h2gproxy sends us to sso from a page that was not authenticated.
// we need to get a token, redirect to the original page, set the cookie and reload
func skipAuth(cr *requesttracker.Request) (*pb.WebloginResponse, error) {
	ctx := cr.Context()
	req := cr.Request()
	u := auth.GetUser(ctx)
	cr.Debugf("Skipping auth for user %s\n", auth.Description(u))
	magic := req.Submitted[common.WEBLOGIN_STATE]
	state, err := getMagic(ctx, cr, magic)
	if err != nil {
		return nil, err
	}
	tr, err := authManager.GetTokenForMe(ctx, &au.GetTokenRequest{DurationSecs: uint64(common.AuthTokenLifetime().Seconds())})
	if err != nil {
		return nil, err
	}
	state.Token = tr.Token
	err = putMagic(cr, magic, state)
	if err != nil {
		return nil, err
	}
	res := NewWebloginResponse()
	m := map[string]string{"weblogin": magic}
	target := stateToURL(state, m)
	cr.Debugf("Redirecting to %s\n", target)
	res.RedirectTo = target
	return res, nil
}

/*********************************************
* process a form that was submitted.
**********************************************/
func processLogin(cr *requesttracker.Request) (*pb.WebloginResponse, *au.User, error) {
	req := cr.Request()
	ctx := cr.Context()
	cr.Debugf("[processLogin] path for weblogin: %s/%s\n", req.Host, req.Path)
	paras := req.Submitted
	if *captcha_bypass_on_login && paras["captcha_bypass"] == *register.CAPTCHA_BYPASS {
		fmt.Println("captcha bypassed")
	} else if *check_captcha_on_login {
		err := check_captcha(paras["g_captcha"], cr.Request().Host)
		if err != nil {
			fmt.Printf("Captcha verification failed: %s\n", err)
			return nil, nil, err
		}
		fmt.Printf("Captcha verified and OK\n")
	}
	magic := paras[common.WEBLOGIN_STATE]
	if magic == "" {
		return nil, nil, errors.InvalidArgs(ctx, "ordering mismatch", "no state in processLogin")
	}
	state, err := getMagic(ctx, cr, magic)
	if err != nil {
		return nil, nil, err
	}

	cr.Debugf("[processlogin] request coming from %s %s\n", state.TriggerHost, state.TriggerPath)

	apr := &au.AuthenticatePasswordRequest{Email: paras["email"], Password: paras["password"]}
	u, err := authremote.GetAuthClient().GetByPassword(ctx, apr)
	if err != nil {
		return nil, nil, err
	}
	if !u.Valid {
		return nil, nil, errors.AccessDenied(ctx, "user invalid")
	}
	if u == nil || u.User == nil {
		return nil, nil, errors.AccessDenied(ctx, "no user")
	}
	fmt.Printf("User %s (%s) Logged in\n", auth.Description(u.User), u.User.ID)
	qp := map[string]string{
		common.WEBLOGIN_STATE: magic,
	}
	target := stateToURL(state, qp)
	s := "<html><body>Welcome " + u.User.Email + "<br>\nYou were coming from here:</br>\n" + target + "</body></html>"
	b := []byte(s)
	res := NewWebloginResponse()
	cookie_token := u.Token
	dur := common.AuthCookieLifetime()
	if *issue_session_cookie_instead_of_auth {
		dur = common.SessionCookieLifetime()
		ct, err := session_for_user(ctx, cr, u.User)
		if err != nil {
			fmt.Printf("failed to issue session cookie: %s\n", utils.ErrorString(err))
		} else {
			cookie_token = ct
			state.TokenSource = 1
		}
	}
	addCookies(res, cr.CookiesToSet())
	res.Body = b
	addCookie(res, "Auth-Token", cookie_token, dur)
	state.Token = cookie_token
	err = putMagic(cr, magic, state) // update our store with the state including token
	if err != nil {
		return nil, nil, err
	}
	h := strings.Trim(state.TriggerHost, "/")

	res.RedirectTo = fmt.Sprintf("https://%s/weblogin/setcookie?"+common.WEBLOGIN_STATE+"=%s", h, magic)
	return res, u.User, nil
}

func addCookies(wr *pb.WebloginResponse, cookies []*h2gproxy.Cookie) {
	wr.Cookies = append(wr.Cookies, cookies...)

}
func addCookie(wr *pb.WebloginResponse, name string, value string, dur time.Duration) {
	exp := time.Now().Add(dur).Unix()
	hc := &h2gproxy.Cookie{Name: name, Value: value, Expiry: uint32(exp)}
	wr.Cookies = append(wr.Cookies, hc)
}
func stateToURL(state *pb.State, qparas map[string]string) string {
	q := ""

	if state.TriggerQuery != "" {
		q = "?" + state.TriggerQuery
	}
	if len(qparas) != 0 {
		if len(q) == 0 {
			q = "?"
		} else {
			q = q + "&"
		}
		uv := url.Values{}
		for k, v := range qparas {
			uv[k] = []string{v}
		}
		q = q + uv.Encode()
	}
	h := strings.Trim(state.TriggerHost, "/")
	p := strings.Trim(state.TriggerPath, "/")
	if len(p) != 0 {
		p = "/" + p
	}
	return fmt.Sprintf("https://%s%s%s", h, p, q)
}
func createState(cr *requesttracker.Request) (string, *pb.State, error) {
	req := cr.Request()
	state := &pb.State{Method: req.Method, TriggerHost: req.Host, TriggerPath: req.Path, TriggerQuery: req.Query}
	magic := utils.RandomString(60)
	err := putMagic(cr, magic, state)
	if err != nil {
		return "", nil, err
	}
	cr.Debugf("Created state (%s)\n", magic)
	return magic, state, nil
}

type GResponse struct {
	Success      bool
	Challenge_ts time.Time
	Hostname     string
	Action       string
	Errorcodes   []string
	Score        float64
}

func check_captcha(response string, hostname string) error {
	if response == "" {
		return common.Errorf("invalid captcha")
	}
	h := &http.HTTP{}
	h.SetHeader("Content-Type", "application/x-www-form-urlencoded")

	p := map[string]string{
		"secret":   web.CaptchaSecretKey(),
		"response": response,
	}
	deli := ""
	d := ""
	for k, v := range p {
		d = d + deli + k + "=" + v
		deli = "&"
	}
	fmt.Printf("Posting: \"%s\"\n", d)
	hr := h.Post("https://www.google.com/recaptcha/api/siteverify", []byte(d))
	err := hr.Error()
	if err != nil {
		return err
	}
	b := hr.Body()
	g := &GResponse{}
	err = json.Unmarshal(b, g)
	if err != nil {
		return err
	}
	fmt.Printf("Google captcha response: %v, score: %0.1f for host \"%s\"\n", g.Success, g.Score, g.Hostname)
	if !g.Success {
		return nil
	}
	if hostname != g.Hostname {
		return common.UrgentErrorf("google found a different hostname (%s)", g.Hostname)
	}
	return nil

}

func session_for_user(ctx context.Context, cr *requesttracker.Request, u *au.User) (string, error) {
	ns := &sessionmanager.NewSessionRequest{
		IPAddress:   cr.IP(),
		UserID:      u.ID,
		Username:    u.FirstName + " " + u.LastName,
		Useremail:   u.Email,
		TriggerHost: cr.TriggerURL(),
		UserAgent:   cr.UserAgent(),
		BrowserID:   cr.BrowserID(),
	}
	sv, err := sessionmanager.GetSessionManagerClient().NewSession(ctx, ns)
	if err != nil {
		return "", err
	}
	return sv.Token, nil
}
