package register

import (
	"context"
	"fmt"
	"golang.conradwood.net/apis/email"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/weblogin/common"
	"golang.conradwood.net/weblogin/web"
	"net/url"
	"time"
)

func CreateRegisterEmail(ctx context.Context, cr *pb.RegisterState) (*pb.Email, error) {

	return create_email(cr)
}

func create_email(cr *pb.RegisterState) (*pb.Email, error) {
	host := cr.Host
	if host == "" {
		host = web.SSOHost()
	}
	ps, err := make_proto(cr)
	if err != nil {
		return nil, err
	}
	ps = url.QueryEscape(ps)
	res := &pb.Email{
		Subject: fmt.Sprintf("Your registration at %s", host),
		Link:    fmt.Sprintf("https://%s/weblogin/register?v_reg=%s", host, ps),
	}
	txt := `
Your Emailaddress (%s) has recently been used to register at our website \"%s\". If this was not you, please accept our apologies and discard this email.

If this was you, please click on the link below to authorise us to proceed with your registration.
We will not process your registration any further until you have authorised us to do so by clicking on the link below.

%s
`
	res.Body = fmt.Sprintf(txt, cr.Email, host, res.Link)

	return res, nil
}

// turn it into a signed RegisterProto base64 encoded protobuf
func make_proto(cr *pb.RegisterState) (string, error) {
	cr.Created = uint32(time.Now().Unix())
	rs, err := utils.Marshal(cr)
	if err != nil {
		return "", err
	}
	b := common.SignString(rs)
	rp := &pb.RegisterProto{State: rs, Signature: b}
	rps, err := utils.Marshal(rp)
	if err != nil {
		return "", err
	}
	return rps, nil
}
func DecodeEmailLink(ctx context.Context, vreg string) (*pb.RegisterState, error) {
	return decode_email_link(vreg)
}
func decode_email_link(vreg string) (*pb.RegisterState, error) {
	if vreg == "" {
		return nil, errors.InvalidArgs(authremote.Context(), "missing authentication information", "no vreg received")
	}
	rp := &pb.RegisterProto{}
	err := utils.Unmarshal(vreg, rp)
	if err != nil {
		return nil, err
	}
	b := common.Verify([]byte(rp.State), rp.Signature)
	if !b {
		fmt.Printf("invalid link received (unverifiable) [%s]\n", vreg)
		return nil, errors.InvalidArgs(authremote.Context(), "invalid link", "link signature is broken")
	}
	rs := &pb.RegisterState{}
	err = utils.Unmarshal(rp.State, rs)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (rr *RegisterRequest) send_email(w *web.WebRequest) error {
	th := rr.ReferrerHost()
	if th == "" {
		panic("no referrer host for email")
	}
	rs := &pb.RegisterState{Host: th, Email: rr.Email, Created: uint32(time.Now().Unix()), Magic: rr.magic}
	e, err := create_email(rs)
	if err != nil {
		return err
	}
	ctx := authremote.Context()
	t := &email.TemplateEmailRequest{
		Sender:       "foo",
		Recipient:    rs.Email,
		TemplateName: "verify_email",
		Values: map[string]string{
			"email":   rs.Email,
			"host":    rs.Host,
			"link":    e.Link,
			"subject": fmt.Sprintf("Your Registration at %s", rs.Host),
		},
	}
	r, err := email.GetEmailServiceClient().SendTemplate(ctx, t)
	if err != nil {
		fmt.Printf("Failed to send email to %s: %s\n", t.Recipient, utils.ErrorString(err))
		return err
	}
	if !r.Success {
		fmt.Printf("failed attempt to send email to %s", t.Recipient)
		return fmt.Errorf("of an uncategorised error")
	}
	return nil
}
