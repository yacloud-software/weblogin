package requests

import (
	"golang.conradwood.net/go-easyops/prometheus"
)

var (
	//	userProfileClient userProfile.UserProfileHandlerServiceClient
	sucCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "weblogin_successful_logins",
			Help: "V=1 UNIT=hz DESC=successfull logins",
		},
		[]string{},
	)
	failCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "weblogin_failed_logins",
			Help: "V=1 UNIT=hz DESC=failed logins",
		},
		[]string{},
	)
	rsucCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "weblogin_successful_password_resets",
			Help: "V=1 UNIT=hz DESC=successfull password resets",
		},
		[]string{},
	)
	rfailCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "weblogin_failed_password_resets",
			Help: "V=1 UNIT=hz DESC=failed password rests",
		},
		[]string{},
	)
	templateRenderCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "weblogin_template_render_counter",
			Help: "V=1 UNIT=hz DESC=failed to render a page",
		},
		[]string{"name", "counter"},
	)
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "weblogin_serve_request_counter",
			Help: "V=1 UNIT=hz DESC=request counter",
		},
		[]string{"counter"},
	)
)

func init() {
	prometheus.MustRegister(templateRenderCounter, requestCounter)
}
