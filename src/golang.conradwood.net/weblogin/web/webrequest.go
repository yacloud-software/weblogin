package web

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/antidos"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/http"
	"golang.conradwood.net/go-easyops/utils"
	"net"
	"strconv"
	"time"
)

var (
	debug = flag.Bool("debug_web", false, "debug web stuff")
)

type Details interface {
	GetHost() string
	GetPath() string
	GetPeer() string
	GetQuery() string
	GetSubmitted() map[string]string
}
type WebRequest struct {
	ctx     context.Context
	details Details
}

func NewWebRequest(ctx context.Context, details Details) *WebRequest {
	w := &WebRequest{ctx: ctx, details: details}
	return w
}
func (w *WebRequest) GetHost() string {
	return w.details.GetHost()
}
func (w *WebRequest) GetPara(name string) string {
	for k, v := range w.details.GetSubmitted() {
		if k == name {
			return v
		}
	}
	return ""
}

type GResponse struct {
	Success      bool
	Challenge_ts time.Time
	Hostname     string
	Action       string
	Errorcodes   []string
	Score        float64
}

func (w *WebRequest) VerifyCaptcha() (bool, error) {
	grr := w.GetPara("g-recaptcha-response")
	if grr == "" {
		grr = w.GetPara("g_captcha")
		if grr == "" {
			return false, fmt.Errorf("google captcha did not run")
		}
	}
	h := &http.HTTP{}
	h.SetHeader("Content-Type", "application/x-www-form-urlencoded")

	p := map[string]string{
		"secret":   *captcha_secret_key,
		"response": grr,
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
		return false, err
	}
	b := hr.Body()
	g := &GResponse{}
	err = json.Unmarshal(b, g)
	if err != nil {
		return false, err
	}
	fmt.Printf("Google captcha response: %v, score: %0.1f for host \"%s\"\n", g.Success, g.Score, g.Hostname)
	if !g.Success {
		return false, nil
	}
	if w.details.GetHost() != g.Hostname {
		return false, fmt.Errorf("google found a different hostname (%s)", g.Hostname)
	}
	return true, nil
}

// this ip misbehaved, tell antidos
func (w *WebRequest) BadIP(botiness uint32) {
	ctx := authremote.Context()
	peer := w.details.GetPeer()
	ip, port, err := net.SplitHostPort(peer)
	if err != nil {
		fmt.Printf("failed to splithostandport from \"%s\": %s\n", peer, err)
		return
	}
	_, err = strconv.Atoi(port)
	if err != nil {
		fmt.Printf("failed to split port from \"%s\" (%s): %s\n", peer, port, err)
		return
	}
	ifr := &antidos.IPFailureRequest{Message: "weblogin, bad webrequest", IP: ip, Botiness: botiness}
	_, err = antidos.GetAntiDOSClient().IPFailure(ctx, ifr)
	if err != nil {
		fmt.Printf("ANTIDOS failed: %s\n", utils.ErrorString(err))
	} else if *debug {
		fmt.Printf("Told ANTIDOS, that IP %s is misbehaving\n", ip)
	}
}
