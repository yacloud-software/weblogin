package requests

import (
	"flag"
	au "golang.conradwood.net/apis/auth"
	"golang.conradwood.net/go-easyops/authremote"
)

const (
	WEBLOGIN_STATE = "weblogin_state_yacloud"
)

var (
	authManager au.AuthManagerServiceClient
	debug       = flag.Bool("debug", false, "debug v2")
)

type RequestHandler struct {
}

func NewHandler() *RequestHandler {
	if authManager == nil {
		authManager = authremote.GetAuthManagerClient()
	}
	res := &RequestHandler{}
	return res
}
