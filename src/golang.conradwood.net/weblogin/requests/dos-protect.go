package requests

import (
	"flag"
	"fmt"
	"golang.conradwood.net/apis/antidos"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/cache"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/weblogin/requesttracker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	//	"net"
	"sync"
	"time"
)

var (
	ips        = cache.New("anti-dos", time.Duration(5)*time.Minute, 1000)
	enable_dos = flag.Bool("enable_dos_protection", true, "does a basic invalid login detection perip")
)

type ipcache struct {
	total      int
	lock       sync.Mutex
	URLCounter map[string]int
}

// return error if ip makes too many dodgy requests
func IsDosing(cr *requesttracker.Request) error {
	if !*enable_dos {
		return nil
	}
	peer_ip_string := cr.IP()
	if peer_ip_string == "" {
		fmt.Printf("antidos: no ip address!!!\n")
		return nil
	}
	u := auth.GetUser(cr.Context())
	if u != nil {
		return nil
	}
	u = cr.GetUser()
	if u != nil {
		return nil
	}

	var ipc *ipcache
	o := ips.Get(peer_ip_string)
	if o == nil {
		return nil
	}
	ipc = o.(*ipcache)
	url := cr.Request().Host + cr.Request().Path
	if ipc.isPeerOverLimit() {
		fmt.Printf("Blocked peer %s (accessing \"%s\")\n", peer_ip_string, url)
		ctx := authremote.Context()
		_, err := antidos.GetAntiDOSClient().IPFailure(ctx, &antidos.IPFailureRequest{Botiness: 1, Message: "peer over limit", IP: peer_ip_string})
		if err != nil {
			fmt.Printf("Failed to antidos: %s\n", utils.ErrorString(err))
		}
		return status.Error(codes.ResourceExhausted, "you reached your limit of accesses. please try later")
	}
	if ipc.isURLOverLimit(url) {
		fmt.Printf("Blocked peer %s on url %s\n", peer_ip_string, url)
		return status.Error(codes.ResourceExhausted, "you reached your limit of accesses. please try later")
	}
	return nil
}
func (i *ipcache) isURLOverLimit(url string) bool {
	uc := i.URLCounter[url]
	if uc > 20 {
		return true
	}
	return false
}
func (i *ipcache) isPeerOverLimit() bool {
	if i.total > 30 {
		return true
	}
	return false
}

// call if a url is called. TODO: implement optimistic locking
func CountURL(cr *requesttracker.Request) {
	var ipc *ipcache
	o := ips.Get(cr.IP())
	if o == nil {
		ipc = &ipcache{
			URLCounter: make(map[string]int),
		}
		ips.Put(cr.IP(), ipc)
	} else {
		ipc = o.(*ipcache)
	}
	url := cr.Request().Host + cr.Request().Path
	ipc.Count(url)
}

func (i *ipcache) Count(url string) {
	i.lock.Lock()
	o := i.URLCounter[url]
	o++
	i.URLCounter[url] = o
	i.total++
	i.lock.Unlock()
}
