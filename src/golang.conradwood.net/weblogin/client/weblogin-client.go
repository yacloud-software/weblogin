package main

import (
	"flag"
	"fmt"
	"golang.conradwood.net/apis/weblogin"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
)

// static variables for flag parser
var (
	wlc   weblogin.WebloginClient
	host  = flag.String("host", "", "host we are simulating")
	email = flag.String("email", "", "emailaddress we are simulating")
)

func main() {
	flag.Parse()
	wlc = weblogin.GetWebloginClient()
	ctx := authremote.Context()
	rs := &pb.RegisterState{
		Host:  *host,
		Email: *email,
	}
	em, err := wlc.CreateRegisterEmail(ctx, rs)
	utils.Bail("failed reg email", err)
	fmt.Printf("%s\n\n%s\n", em.Subject, em.Body)
	fmt.Printf("Done\n")
}
