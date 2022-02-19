package common

import (
	"bytes"
	"fmt"
	"golang.conradwood.net/go-easyops/utils"
	"html/template"
	"strings"
)

type Template_data interface {
}
type RequestTemplate struct {
	Request              *Request
	Data                 Template_data
	StateQuery           template.HTMLAttr
	Heading              string
	SiteKey              string
	RegistrationEnabled  bool
	Msg                  string
	Weblogin_state_name  string
	Weblogin_state_value string
	Email                string // to pre-fill email fields
}

func (r *Request) Render(page string, data Template_data) ([]byte, error) {
	rt := &RequestTemplate{
		Request:              r,
		RegistrationEnabled:  *allow_registrations,
		Data:                 data,
		StateQuery:           template.HTMLAttr("?" + WEBLOGIN_STATE + "=" + r.ref),
		Weblogin_state_name:  WEBLOGIN_STATE,
		Weblogin_state_value: r.ref,
	}

	templ := ""
	b, err := utils.ReadFile("templates/v3/header.html")
	if err != nil {
		return nil, err
	}
	templ = templ + string(b)

	b, err = utils.ReadFile("templates/v3/" + page + ".html")
	if err != nil {
		return nil, err
	}
	templ = templ + string(b)

	b, err = utils.ReadFile("templates/v3/footer.html")
	if err != nil {
		return nil, err
	}
	templ = templ + string(b)

	t, err := template.New(page).Parse(templ)
	if err != nil {
		printTemplate(templ)
		return nil, err
	}
	buf := &bytes.Buffer{}
	err = t.Execute(buf, rt)
	if err != nil {
		printTemplate(templ)
		return nil, err
	}

	return buf.Bytes(), nil
}

func printTemplate(s string) {
	for i, line := range strings.Split(s, "\n") {
		fmt.Printf("%02d. %s\n", (i + 1), line)
	}
}
