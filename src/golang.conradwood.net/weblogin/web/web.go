package web

import (
	"flag"
)

var (
	allow_register     = flag.Bool("allow_register", false, "allow registrations of new users")
	template_path      = flag.String("template_path", "templates", "Path to the HTML templates")
	ssohost            = flag.String("ssohost", "sso.yacloud.eu", "the host to use as sso domain")
	captcha_key        = flag.String("captcha_key", "", "google recaptcha key")
	captcha_secret_key = flag.String("captcha_server_key", "", "google recaptcha secret key")
)

func AllowRegister() bool {
	return *allow_register
}
func CaptchaKey() string {
	return *captcha_key
}
func CaptchaSecretKey() string {
	return *captcha_secret_key
}
func SSOHost() string {
	return *ssohost
}
func TemplatePath() string {
	return *template_path
}
