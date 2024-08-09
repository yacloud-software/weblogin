package requests

import (
	"context"
	"fmt"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/weblogin/opts"
	"golang.conradwood.net/weblogin/web"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"html/template"
	"time"
)

type errordata struct {
	ErrorText string
}

func (rr *errordata) Year() string {
	t := time.Now().Year()
	return fmt.Sprintf("%d", t)
}

func (rr *errordata) GetQueryValue(key string) string {
	return ""
}

func (rr *errordata) TargetURL() string {
	return "/"
}

func (e *errordata) Username() string {
	return ""
}
func (e *errordata) ReferrerHost() string {
	return ""
}
func (e *errordata) Heading() string {
	return ""
}
func (e *errordata) StateQuery() template.HTMLAttr {
	return template.HTMLAttr("")
}

// this must not return an error ever. (an error is badly displayed as plain text to user)
func ServeError(ctx context.Context, req *pb.WebloginRequest, err error) (*pb.WebloginResponse, error) {
	if opts.IsDebug() {
		fmt.Printf("Serving error: %v\n", err)
	}
	wr := web.NewWebRequest(ctx, req)
	res := &pb.WebloginResponse{
		Authenticated: false,
		User:          nil,
		Token:         "",
		PeerIP:        "",
		HTTPCode:      400,
	}
	st := status.Convert(err)
	ed := &errordata{ErrorText: fmt.Sprintf("%s", st.Message())}
	if st.Code() == codes.ResourceExhausted {
		ed.ErrorText = "You have had too many unsuccessful login attempts. Please try later."
	}
	b, err := wr.Render("error", ed)
	if err != nil {
		fmt.Printf("error rendering failed: %s\n", err)
		res.Body = []byte(`<html><body>cannot render error message</body></html>`)
	} else {
		res.Body = b
	}
	return res, nil
}
