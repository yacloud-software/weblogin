package requests

import (
	"bytes"
	"fmt"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/weblogin/web"
	"text/template"
)

type template_data interface {
	Username() string
}

func (cr *Request) renderTemplate(l template_data, templateFile string) ([]byte, error) {
	cr.Debugf("Rendering template %s\n", templateFile)

	tfname := web.TemplatePath() + "/" + templateFile + ".html"
	t := template.New(templateFile)
	t.Funcs(template.FuncMap{
		"username": l.Username,
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
