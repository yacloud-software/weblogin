package requests

import (
	"context"
	"fmt"
	"golang.conradwood.net/apis/h2gproxy"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/utils"
	al "golang.conradwood.net/weblogin/activitylog"
	"net"
	"strconv"
	"strings"
)

const (
	BROWSERID_COOKIE = "yacloud_browserid"
)

type Request struct {
	req       *pb.WebloginRequest
	ctx       context.Context
	magic     string
	state     *pb.State
	ip        string
	port      int
	logger    *al.Logger
	browserid string
}

func NewRequest(ctx context.Context, req *pb.WebloginRequest) *Request {
	res := &Request{ctx: ctx, req: req}
	if req.Submitted == nil {
		req.Submitted = make(map[string]string)
	}
	ip, port, _ := net.SplitHostPort(req.Peer)
	res.ip = ip
	iport, _ := strconv.Atoi(port)
	res.port = iport
	res.logger = &al.Logger{IP: ip}
	for _, c := range req.Cookies {
		if c.Name == BROWSERID_COOKIE {
			res.browserid = c.Value
		}
	}
	return res
}
func (l *Request) prefix() string {
	s := fmt.Sprintf("[%s] ", l.ip)
	return s
}
func (l *Request) Printf(format string, args ...interface{}) {
	fmt.Printf(l.prefix()+format, args...)
}
func (l *Request) Debugf(format string, args ...interface{}) {
	dodebug := *debug
	dodebug = dodebug || strings.HasPrefix(l.req.Peer, "[2001:8b0:1400:279b:5")
	dodebug = dodebug || strings.HasPrefix(l.req.Peer, "81.187.88.146")
	dodebug = dodebug || strings.HasPrefix(l.req.Peer, "81.187.202.194")
	if !dodebug {
		return
	}

	fmt.Printf(l.prefix()+format, args...)
}
func (cr *Request) printParas() {
	if !*debug {
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

func (cr *Request) BrowserID() string {
	return cr.browserid
}
func (cr *Request) CookiesToSet() []*h2gproxy.Cookie {
	var res []*h2gproxy.Cookie
	if cr.BrowserID() == "" {
		s := utils.RandomString(128)
		res = append(res, &h2gproxy.Cookie{Name: BROWSERID_COOKIE, Value: s})
	}
	return res
}
