package common

import (
	"bytes"
	"fmt"
	"golang.conradwood.net/go-easyops/utils"
	"html/template"
)

type Template_data interface {
}
type RequestTemplate struct {
	Request             *Request
	Data                Template_data
	StateQuery          template.HTMLAttr
	Heading             string
	SiteKey             string
	RegistrationEnabled bool
}

func (r *Request) Render(page string, data Template_data) ([]byte, error) {
	rt := &RequestTemplate{
		Request:    r,
		Data:       data,
		StateQuery: template.HTMLAttr("?" + WEBLOGIN_STATE + "=" + r.ref),
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
		fmt.Printf("1. Template:\n%s\n", templ)
		return nil, err
	}
	buf := &bytes.Buffer{}
	err = t.Execute(buf, rt)
	if err != nil {
		fmt.Printf("2. Template:\n%s\n", templ)
		return nil, err
	}

	return buf.Bytes(), nil
}
