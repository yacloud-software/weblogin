package web

import (
	"bytes"
	"fmt"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/weblogin/common"
	"html/template"
)

func (w *WebRequest) Render(templateFile string, data common.Template_data) ([]byte, error) {
	s1, err := utils.ReadFile(TemplatePath() + "/header.html")
	if err != nil {
		return nil, err
	}
	s2, err := utils.ReadFile(TemplatePath() + "/" + templateFile + ".html")
	if err != nil {
		return nil, err
	}
	s3, err := utils.ReadFile(TemplatePath() + "/footer.html")
	if err != nil {
		return nil, err
	}
	s1 = append(s1, s2...)
	s1 = append(s1, s3...)
	t := template.New(templateFile)
	t.Funcs(template.FuncMap{
		"username":      w.Username,
		"StateQuery":    data.StateQuery,
		"Heading":       data.Heading,
		"ReferrerHost":  data.ReferrerHost,
		"TargetURL":     data.TargetURL,
		"GetQueryValue": data.GetQueryValue,
	})
	_, err = t.Parse(string(s1))
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	err = t.Execute(buf, data)
	if err != nil {
		fmt.Printf("Failed to execute template: %s\n", err)
		return nil, err
	}
	return buf.Bytes(), nil
}
func (w *WebRequest) Username() string {
	return ""
}
