package main

import (
	"fmt"
	apb "golang.conradwood.net/apis/auth"
	pb "golang.conradwood.net/apis/weblogin"
	"io"
	"net/http"
	"strings"
)

type PageContext struct {
	resp   *pb.WebloginResponse
	req    *pb.WebloginRequest
	peerip string
	Path   string
	PW1    string
	PW2    string
	Error  bool
	Msg    string
	Key    string
	Email  string // email for forgot-password
	w      io.Writer
	hrw    http.ResponseWriter
	r      *http.Request
	user   *apb.User
}

func (p *PageContext) reload() {
	if p.hrw == nil || p.w == nil {
		fmt.Printf("Cannot reload - no request or writer\n")
		return
	}

	u := p.FormValue("url")
	fmt.Printf("Redirecting to \"%s\"\n", u)
	http.Redirect(p.hrw, p.r, u, http.StatusSeeOther)

}
func (p *PageContext) printForm() {
	var ms map[string]string
	ms = p.req.Submitted

	if ms != nil {
		for k, v := range ms {
			fmt.Printf(" %s -> \"%s\"\n", k, v)
		}
		return
	}
	if p.r == nil {
		fmt.Printf("No form\n")
		return
	}
	for k, v := range p.r.Form {
		if len(v) == 1 {
			fmt.Printf("%s = \"%s\"\n", k, v[0])
		} else {
			fmt.Printf("%s: (%d values)\n", k, len(v))
			for _, l := range v {
				fmt.Printf("   %s\n", l)
			}
		}
	}
}
func (p *PageContext) error(err error) {
	if p.hrw != nil {
		p.hrw.WriteHeader(500)
	}
	fmt.Printf("Error: %s\n", err)
}

func (p *PageContext) Username() string {
	if p.user == nil {
		return "nobody"
	}
	return p.user.FirstName + " " + p.user.LastName + " (" + p.user.Email + ")"
}
func (p *PageContext) FormValue(name string) string {
	return p.req.Submitted[name]
}

func (p *PageContext) PeerIP() string {
	pe := p.peerip
	i := strings.LastIndex(pe, ":")
	if i != -1 {
		pe = pe[:i]
	}
	return pe
}
