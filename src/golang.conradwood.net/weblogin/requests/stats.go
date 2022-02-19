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
)
