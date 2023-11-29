package register

import (
	"context"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	"golang.conradwood.net/apis/email"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/utils"
	"net/url"
)

var (
	enable_signup_rpc = flag.Bool("enable_signup_resend_rpc", false, "if true enable an rpc to send sign-up emails. if false, the rpc will return errors")
)

func SignupEmailRPC(ctx context.Context, req *pb.SignupEmail) (*common.Void, error) {
	if !*enable_signup_rpc {
		return nil, fmt.Errorf("the resending of signup emails is administratively prohibited")
	}
	rs := &pb.RegisterState{
		Host:  req.Host,
		Email: req.Email,
	}
	link_data, err := make_proto(rs)
	if err != nil {
		return nil, err
	}
	link_data = url.QueryEscape(link_data)

	t := &email.TemplateEmailRequest{
		Sender:       "",
		Recipient:    rs.Email,
		TemplateName: "verify_email",
		Values: map[string]string{
			"email":   rs.Email,
			"host":    rs.Host,
			"link":    fmt.Sprintf("https://%s/weblogin/register?v_reg=%s", rs.Host, link_data),
			"subject": fmt.Sprintf("Your Registration at %s", rs.Host),
		},
	}
	r, err := email.GetEmailServiceClient().SendTemplate(ctx, t)
	if err != nil {
		fmt.Printf("Failed to send email to %s: %s\n", t.Recipient, utils.ErrorString(err))
		return nil, err
	}
	if !r.Success {
		fmt.Printf("failed attempt to send email to %s", t.Recipient)
		return nil, fmt.Errorf("of an uncategorised error")
	}

	return &common.Void{}, nil
}
