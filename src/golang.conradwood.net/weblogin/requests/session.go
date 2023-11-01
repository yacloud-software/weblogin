package requests

import (
	"fmt"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/weblogin/requesttracker"
)

// this needs to check if there is a session (e.g. as cookie).
// if not, create a new one
// if there is, reuse it
// then, redirect browser back to where it came from (using the same method, payload etc)
// the last redirect needs to include some short-lived token, which h2gproxy can use to
// lookup the session
// (it is h2gproxy's job to then set the session cookie on the domain)
func needSessionPage(cr *requesttracker.Request) (*pb.WebloginResponse, error) {
	return nil, fmt.Errorf("need session not implemented")
}
