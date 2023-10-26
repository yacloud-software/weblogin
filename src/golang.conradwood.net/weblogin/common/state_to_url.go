package common

import (
	"fmt"
	pb "golang.conradwood.net/apis/weblogin"
	"net/url"
)

func State2URL(state *pb.State) string {
	res := "https://" + state.TriggerHost + "/" + state.TriggerPath
	if state.TriggerQuery != "" {
		res = res + "?" + state.TriggerQuery
	}
	return res
}
func State2URLValues(state *pb.State) map[string]string {
	res := make(map[string]string)
	surl := State2URL(state)
	u, err := url.Parse(surl)
	if err != nil {
		fmt.Printf("[error] url invalid (%s): %s\n", surl, err)
		return res
	}
	values := u.Query()
	for k, v := range values {
		if len(v) == 0 {
			continue
		}
		res[k] = v[0]
	}
	return res
}
