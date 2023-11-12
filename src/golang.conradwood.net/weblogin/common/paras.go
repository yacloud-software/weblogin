package common

import (
	"flag"
	"time"
)

var (
	auth_cookie_lifetime = flag.Duration("auth_cookie_lifetime", time.Duration(43200)*time.Second, "auth cookie lifetime")
	auth_token_lifetime  = flag.Duration("auth_token_lifetime", time.Duration(43200)*time.Second, "auth token lifetime")
	sess_cookie_lifetime = flag.Duration("sess_cookie_lifetime", time.Duration(48)*time.Hour, "session cookie lifetime")
)

func AuthCookieLifetime() time.Duration {
	return *auth_cookie_lifetime
}
func SessionCookieLifetime() time.Duration {
	return *sess_cookie_lifetime
}
func AuthTokenLifetime() time.Duration {
	return *auth_token_lifetime
}
