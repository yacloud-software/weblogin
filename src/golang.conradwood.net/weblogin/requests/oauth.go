package requests

import (
	"fmt"
	au "golang.conradwood.net/apis/auth"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/utils"
	//	"net/url"
	"strings"
)

/*
replaces: https://actions.conradwood.net/google-oauth?linkid=oauth&provider=google
*/

const (
	MY_GOOGLE_CLIENT_ID     = "Quiechoh7neSieMued5poukohy2baew5ajameameich3mi8quoh2Aefuzo7xie7u"
	MY_GOOGLE_ROJECT_ID     = "myhome-a97d7"
	MY_GOOGLE_CLIENT_SECRET = ""
)

func googleOAuth(cr *Request) (*pb.WebloginResponse, error) {
	ctx := cr.ctx
	req := cr.req
	u := auth.GetUser(ctx)
	if u == nil {
		return nil, errors.Unauthenticated(ctx, "please log in")
	}
	if singlePara(req, "completed") == "true" {
		return completed(cr, u)
	}
	cid := singlePara(req, "client_id")
	if cid != MY_GOOGLE_CLIENT_ID {
		return nil, errors.InvalidArgs(ctx, "invalid client id", "invalid client id (%s)", cid)
	}
	// create a funny token:
	google_token := utils.RandomString(128)

	// store it
	_, err := authremote.GetAuthManagerClient().StoreRemote(ctx, &au.RemoteStoreRequest{
		UserID:       u.ID,
		Provider:     "GOOGLE",
		OurToken:     google_token,
		RemoteUserID: "",
	})
	if err != nil {
		fmt.Printf("Failed to store token: %s\n", utils.ErrorString(err))
		return nil, err
	}

	fmt.Printf("Google token for %s (%s): %s\n", u.ID, u.Email, google_token)
	paras := make(map[string]string)
	paras["access_token"] = google_token
	//paras["code"] = google_token
	paras["token_type"] = "bearer"
	paras["state"] = singlePara(req, "state")
	//paras["client_id"] = MY_GOOGLE_CLIENT_ID
	//paras["response_type"] = "authorization_code"
	//paras["client_secret"] = MY_GOOGLE_CLIENT_SECRET

	gredir := singlePara(req, "redirect_uri")
	deli := "#"
	gurl := gredir
	paras["completed"] = "true"
	//paras["redirect_uri"] = "https://sso.yacloud.eu/weblogin/oauth"
	for k, v := range paras {
		gurl = gurl + deli + k + "=" + v
		deli = "&"
	}
	gurl = gredir + "#state=" + paras["state"] + "&completed=true" + "&access_token=" + paras["access_token"] + "&token_type=" + paras["token_type"]
	fmt.Printf("Google url for %s (%s): %s\n", u.ID, u.Email, gurl)
	body := `<html><body>
Hello %s %s<br/>
You attempt to interact with singingcat using a Google Account.
In order to access your devices, we need to link your singingcat account with the account you use for Google. This is so that we can ensure only you has access to your singingcat devices (and nobody else).

Please complete your authentication at google, using the below url:<br/>
<a href="%s">Google Authentication OAuth Endpoint thing</a></br>
If at a later stage you want to withdraw consent, visit googles' site:
https://myaccount.google.com/accountlinking <a href="https://myaccount.google.com/accountlinking">Link</a><br/>

</body>
</html>`

	res := &pb.WebloginResponse{
		Headers: make(map[string]string),
	}
	res.Headers["content-type"] = "text/html"
	res.Body = []byte(fmt.Sprintf(body, u.FirstName, u.LastName, gurl))
	// rebuild url for redirect location header:
	/*
		deli = "?"
		gurl = gredir
		for k, v := range paras {
			ev := url.QueryEscape(v)
			gurl = gurl + deli + k + "=" + ev
			deli = "&"
		}
	*/
	res.RedirectTo = gurl
	res.ForceGetAfterRedirect = true
	cr.Debugf("Redirecting to: %s\n", res.RedirectTo)
	return res, nil
}

// find a single parameter in url, case insensitive
func singlePara(req *pb.WebloginRequest, name string) string {
	name = strings.ToLower(name)
	for k, v := range req.Submitted {
		if strings.ToLower(k) != name {
			continue
		}
		return v
	}
	return ""
}
func completed(cr *Request, u *au.User) (*pb.WebloginResponse, error) {
	body := `<html><body>
Welcome %s %s<br/>
Authorization completed.<br/>
</body>
</html>
`
	res := &pb.WebloginResponse{
		Headers: make(map[string]string),
	}
	res.Headers["content-type"] = "text/html"
	res.Body = []byte(fmt.Sprintf(body, u.FirstName, u.LastName))
	return res, nil
}
