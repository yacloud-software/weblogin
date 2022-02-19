package common

import (
	"context"
	"fmt"
	au "golang.conradwood.net/apis/auth"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/utils"
)

// this is a single "request", a bit like an extended context
type Request struct {
	ctx      context.Context
	Req      *pb.WebloginRequest
	failure  error // if nil we have to serve an error
	state    *pb.V3State
	ref      string // the string with which we refer to the state
	user     *au.User
	response *pb.WebloginResponse // if set, will be sent by Close()
}

func NewRequest(ctx context.Context, req *pb.WebloginRequest) *Request {
	r := &Request{ctx: ctx, Req: req}
	r.user = auth.GetUser(ctx)

	ref := r.Req.Submitted[WEBLOGIN_STATE]
	if ref != "" {
		r.ref = ref
		state, err := ParseMagic(r.ctx, ref)
		if err != nil {
			r.failure = err
			r.state = r.buildNewState()
			return r
		}
		r.state = state
		return r
	}
	r.state = r.buildNewState()
	r.ref = utils.RandomString(32)
	err := SaveMagic(ctx, r.ref, r.state)
	if err != nil {
		r.failure = err
	}
	return r
}

// build a new state from request
func (r *Request) buildNewState() *pb.V3State {
	res := &pb.V3State{
		TriggerHost:  r.Req.Host,
		TriggerPath:  r.Req.Path,
		TriggerQuery: r.Req.Query,
	}
	return res
}

func (r *Request) State() *pb.V3State {
	return r.state
}
func (r *Request) Debugf(format string, args ...interface{}) {
	uss := "nobody"
	if r.user != nil {
		uss = fmt.Sprintf("%s(%s)", r.user.ID, r.user.Email)
	}
	s := fmt.Sprintf("[user=%s, host=%s, path=%s] ", uss, r.State().TriggerHost, r.State().TriggerPath)
	fmt.Printf(s+format+"\n", args...)
}

func (r *Request) GetURLQuery() string {
	return "?" + WEBLOGIN_STATE + "=" + r.ref
}

// build a weblogin response to redirect to url
func (r *Request) RedirectTo(url string) (*pb.WebloginResponse, error) {
	res := &pb.WebloginResponse{RedirectTo: url}
	r.Debugf("redirecting to: %s", url)
	return res, nil
}

func (r *Request) SaveState() {
	err := SaveMagic(r.ctx, r.ref, r.state)
	if err != nil {
		r.failure = err
	}
}

// respond with html
func (r *Request) SetResponse(b *pb.WebloginResponse) {
	r.response = b
}

func (r *Request) Close() (*pb.WebloginResponse, error) {
	if r.failure == nil {
		if r.response != nil {
			return r.response, nil
		}
		r.failure = fmt.Errorf("nothing to show here")
	}
	res := &pb.WebloginResponse{
		Body:     []byte(fmt.Sprintf("%s", r.failure)),
		MimeType: "text/plain",
	}
	return res, nil
}

func (r *Request) Error(err error) {
	if r.failure == nil && err != nil {
		r.failure = err
	}
}
