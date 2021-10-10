package requests

// user made it past the login form, got authenticated and the weblogin server
// redirected to the original host which triggered the request.
import (
	"fmt"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/errors"
	//	"golang.conradwood.net/go-easyops/utils"
)

// got here after login. url is /weblogin/setcookie and we redirect to url as requested by user in the first request
func setCookiePage(cr *Request) (*pb.WebloginResponse, error) {
	ctx := cr.ctx
	req := cr.req
	paras := req.Submitted
	magic := paras[WEBLOGIN_STATE]
	if magic == "" {
		return nil, errors.InvalidArgs(ctx, "missing state", "setcookiepage - missing state")
	}
	state, err := cr.getMagic(ctx, magic)
	if err != nil {
		return nil, err
	}
	m := map[string]string{}
	target := stateToURL(state, m)
	fmt.Printf("Setting cookie and redirecting to %s....\n", target)
	res := NewWebloginResponse()
	res.RedirectTo = target
	addCookie(res, "Auth-Token", state.Token)
	if state.Token == "" {
		return nil, errors.AccessDenied(ctx, "missing token")
	}

	return res, nil
}
