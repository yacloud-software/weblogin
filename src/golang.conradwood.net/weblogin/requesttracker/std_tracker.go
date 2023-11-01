package requesttracker

import (
	"context"
	"fmt"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/weblogin/opts"
	"google.golang.org/grpc"
	"time"
)

var (
	logger_chan = make(chan *logreq, 100)
	service_map = make(map[string]*logservice)
)

func init() {
	go log_chan()
}

// TODO: maybe implement a "cleaner" if use many different services
type logservice struct {
	tracker     *pb.TrackerServiceClient
	cn          *grpc.ClientConn
	servicename string
	last_used   time.Time
}
type logreq struct {
	req *pb.AuthActivityRequest
}

type StandardTracker struct {
}

func (st *StandardTracker) LogActivity(action pb.AuthAction, req *pb.AuthActivityRequest) {
	fmt.Printf("[standardtracker %v] %#v\n", action, req)
	req.Action = action
	logreq := &logreq{req: req}
	select {
	case logger_chan <- logreq:
		//
	case <-time.After(time.Duration(500) * time.Millisecond):
		fmt.Printf("Unable to log request - timeout\n")
	default:
		fmt.Printf("Unable to log request - no space\n")
	}
}
func log_chan() {
	for {
		lr := <-logger_chan
		ls := get_service()
		ctx := authremote.ContextWithTimeout(time.Duration(500) * time.Millisecond)
		_, err := ls.LogActivity(ctx, lr.req)
		if err != nil {
			fmt.Printf("FAILED TO LOG ACTIVITY: %s\n", utils.ErrorString(err))
		}

	}
}
func get_service() *logservice {
	sn := opts.UserJourneyTracker()
	res := service_map[sn]
	if res != nil {
		return res
	}
	res = &logservice{}
	service_map[sn] = res
	res.cn = client.Connect(sn)
	res.servicename = sn
	return res
}

func (ls *logservice) LogActivity(ctx context.Context, req *pb.AuthActivityRequest) (*common.Void, error) {
	ls.last_used = time.Now()
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/"+ls.servicename+"/LogActivity", req, out, ls.cn)
	if err != nil {
		return nil, err
	}
	return out, nil
}
