package requests

import (
	"bytes"
	"context"
	"fmt"
	"golang.conradwood.net/apis/themes"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/weblogin/common"
	"golang.conradwood.net/weblogin/web"
	"html/template"
)

type extra_data struct {
	td  common.Template_data
	ctx context.Context
}

func (e *extra_data) Heading() string {
	t, err := themes.GetThemesClient().GetHeaderText(e.ctx, &themes.HostThemeRequest{Host: e.td.ReferrerHost()})
	if err != nil {
		fmt.Printf("unable to get heading: %s\n", utils.ErrorString(err))
		return ""
	}
	return t.Text
}

func (cr *Request) renderTemplate(l common.Template_data, templateFile string) ([]byte, error) {
	cr.Debugf("Rendering template %s\n", templateFile)
	tfname := web.TemplatePath() + "/" + templateFile + ".html"
	t := template.New(templateFile)
	e := &extra_data{td: l, ctx: cr.ctx}
	t.Funcs(template.FuncMap{
		"username":     l.Username,
		"StateQuery":   l.StateQuery,
		"Heading":      e.Heading,
		"ReferrerHost": l.ReferrerHost,
		"TargetURL":    l.TargetURL,
	})
	b, err := readTemplateFile(tfname)
	if err != nil {
		cr.Printf("Unable to load template \"%s\"\n", tfname)
		return nil, err
	}

	_, err = t.Parse(string(b))
	if err != nil {
		fmt.Printf("Unable to load template \"%s\"\n", tfname)
		return nil, err
	}

	buf := &bytes.Buffer{}
	err = t.Execute(buf, l)
	if err != nil {
		cr.Printf("Failed to execute template: %s\n", err)
		return nil, err
	}
	return buf.Bytes(), nil
}

func readTemplateFile(templateFile string) (string, error) {
	s1, err := utils.ReadFile(web.TemplatePath() + "/header.html")
	if err != nil {
		return "", err
	}
	s2, err := utils.ReadFile(templateFile)
	if err != nil {
		return "", err
	}
	s3, err := utils.ReadFile(web.TemplatePath() + "/footer.html")
	if err != nil {
		return "", err
	}
	return string(s1) + string(s2) + string(s3), nil
}
