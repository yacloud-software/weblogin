package main

import (
	"flag"
	"fmt"
	apb "golang.conradwood.net/apis/auth"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/weblogin/requests"
	"golang.conradwood.net/weblogin/web"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

/*
the h2gproxy forwards requests for authentication to us.
it's a rather old single-site buggy code. A new implementation is being worked on in  weblogin2-server.go

This is how it _should_ work:
weblogin will handle the authentication process and then redirect
to the URL originally asked for.
The authentication works as follows:
1) weblogin will redirect immediately to https://sso.singingcat.net
2) weblogin will look for cookie - if so authenticate immediately
2) weblogin will serve a login page
3) if submitted, weblogin will authenticate the user (error if not valid)
4) if authenticated, weblogin will set a cookie (valid for sso.singingcat.net)
5) redirect to the [domain]/weblogin?otp=[OTP] the user wanted to browse with a OTP in url
6) using the OTP, authenticate the request
7) set cookie (valid for the domain)
8) redirect to url user originally requested...

*/
// static variables for flag parser
var (
	port            = flag.Int("port", 10001, "The server port")
	httpport        = flag.Int("http_port", 8091, "The port to start the HTTP listener on")
	cookie_livetime = flag.Int("cookie_expiry", 30*60, "cookie expiry time in seconds")
	authClient      apb.AuthenticationServiceClient
	authMgr         apb.AuthManagerServiceClient
	rh              *requests.RequestHandler
)

func main() {
	flag.Parse() // parse stuff. see "var" section above
	web.InitKey()
	requests.Cookie_livetime = cookie_livetime
	rh = requests.NewHandler()
	go rh.StartGRPC(*port)
	fmt.Printf("Starting http server on port %d\n", *httpport)
	sd := server.NewHTMLServerDef("weblogin.Weblogin")
	sd.Port = *httpport
	server.AddRegistry(sd)
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", *httpport))
	if err != nil {
		panic(err)
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", sd.Port),
		Handler: &webhandler{},
	}
	err = srv.Serve(conn)
	utils.Bail("failed to server", err)
}

type webhandler struct {
}

// we come in here if h2gproxy forwards to an http target and that target requires authentication
// at that point we can't do HTTP because h2gproxy is in a 'http proxy' thus it needs an http target
func (w *webhandler) ServeHTTP(response http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Body read Error: %s\n", utils.ErrorString(err))
		response.Write([]byte("weblogin error"))
		return
	}
	url := req.URL
	host := "unknown-host"
	peerip := "unknown-ip"
	loc := "unknown-location"
	for h, vl := range req.Header {
		if len(vl) == 0 {
			continue
		}
		v := vl[0]
		lh := strings.ToLower(h)
		if lh == "x-forwarded-host" {
			host = v
		}

		if lh == "x-requested-location" {
			loc = v
		}
		if lh == "remote_addr" {
			ip, _, err := net.SplitHostPort(v)
			if err != nil {
				fmt.Printf("Invalid remote_addr \"%s\": %s\n", v, err)
			} else {
				peerip = ip
			}
		}

		fmt.Printf("header: \"%s\" == \"%s\"\n", h, v)
	}
	fmt.Printf("URL: %#v\n", url)
	ctx := authremote.Context()
	greq := &pb.WebloginRequest{
		Method: "get",
		Scheme: "http",
		Host:   host,
		Path:   loc,
		Query:  url.RawQuery,
		Body:   string(body),
		Peer:   peerip,
	}
	fmt.Printf("HTTP proxy request for %s://%s/%s from ip %s\n", greq.Scheme, greq.Host, greq.Path, peerip)
	wr, err := rh.GetLoginPage(ctx, greq)
	if err != nil {
		fmt.Printf("Error: %s\n", utils.ErrorString(err))
		response.Write([]byte("weblogin error"))
		return
	}
	if wr.RedirectTo != "" {
		fmt.Printf("Redirecting to %s\n", wr.RedirectTo)
		http.Redirect(response, req, wr.RedirectTo, 302)
	}
	fmt.Printf("Response: %s\n", string(wr.Body))
	response.Write(wr.Body)
}
