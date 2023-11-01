package requests

import (
	au "golang.conradwood.net/apis/auth"
	"golang.conradwood.net/go-easyops/authremote"
)

var (
	authManager au.AuthManagerServiceClient
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
