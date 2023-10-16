package common

import (
	pb "golang.conradwood.net/apis/weblogin"
)

func State2URL(state *pb.State) string {
	res := "https://" + state.TriggerHost + "/" + state.TriggerPath
	if state.TriggerQuery != "" {
		res = res + "?+state.TriggerQuery"
	}
	return res
}
