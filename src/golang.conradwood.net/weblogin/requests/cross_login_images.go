package requests

import (
	"fmt"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/weblogin/common"
)

var (
	smallimg = []byte{
		0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, 0x00, 0x00, 0x00,
		0x0D, 0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
		0x00, 0x01, 0x08, 0x02, 0x00, 0x00, 0x00, 0x90, 0x77, 0x53, 0xDE,
		0x00, 0x00, 0x00, 0x0C, 0x49, 0x44, 0x41, 0x54, 0x08, 0xD7, 0x63,
		0xF8, 0xFF, 0xFF, 0x3F, 0x00, 0x05, 0xFE, 0x02, 0xFE, 0xDC, 0xCC,
		0x59, 0xE7, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4E, 0x44, 0xAE,
		0x42, 0x60, 0x82,
	}
)

// given the host the user requested we should authenticate some other domains here as well
// for now, it is hardcoded...
func createDomainLogins(cr *Request) ([]string, error) {
	submittedParameters := cr.req.Submitted
	state := submittedParameters[common.WEBLOGIN_STATE]

	domains := []string{"sso.yacloud.eu", "api.conradwood.net", "api.yacloud.eu", "api.yacloud.eu", "api.singingcat.net"}
	var res []string
	for _, d := range domains {
		res = append(res, fmt.Sprintf("https://%s/weblogin/preauthimg?"+common.WEBLOGIN_STATE+"=%s", d, state))
	}
	return res, nil
}

// we set a cookie here with a random string. this random string will be associated once (and only once)
// the authentication is complete.
// there is also is a distinct possibility that we have already authenticated the user (because a cookie for sso.yacloud.eu was set when they got here)
// if so, we just need to set more auth cookies
func preAuthCookie(cr *Request) (*pb.WebloginResponse, error) {
	res := NewWebloginResponse()
	res.Body = smallimg
	res.MimeType = "image/png"
	addCookie(res, "Pre-Auth-Token", "weblogin_is_incomplete_cross_domain_pre_auth_not_working_yet")
	return res, nil
}
