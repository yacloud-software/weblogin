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
	"sync"
	"time"
)

var (
	logger_chan = make(chan *logreq, 100)
	service_map = make(map[string]*logservice)
	retries     []*logreq
	retry_lock  sync.Mutex
)

func init() {
	go log_chan()
	go log_chan_retry()
}

// TODO: maybe implement a "cleaner" if use many different services
type logservice struct {
	tracker     *pb.TrackerServiceClient
	cn          *grpc.ClientConn
	servicename string
	last_used   time.Time
}
type logreq struct {
	req          *pb.AuthActivityRequest
	last_attempt time.Time
	done         bool
	counter      int
}

type StandardTracker struct {
}

func (st *StandardTracker) LogActivity(action pb.AuthAction, req *pb.AuthActivityRequest) {
	//	fmt.Printf("[standardtracker %v] %#v\n", action, req)
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
		lr.last_attempt = time.Now()
		ls := get_service()
		ctx := authremote.ContextWithTimeout(time.Duration(500) * time.Millisecond)
		_, err := ls.LogActivity(ctx, lr.req)
		if err != nil {
			fmt.Printf("FAILED TO LOG ACTIVITY (to service \"%s\"): %s\n", ls.servicename, utils.ErrorString(err))
			retry_lock.Lock()
			retries = append(retries, lr)
			retry_lock.Unlock()
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

func log_chan_retry() {
	for {
		time.Sleep(time.Duration(3) * time.Second)
		retry_lock.Lock()
		rs := retries
		retry_lock.Unlock()
		for _, r := range rs {
			if time.Since(r.last_attempt) < time.Duration(4)*time.Second {
				continue
			}
			r.counter++
			r.last_attempt = time.Now()
			ls := get_service()
			ctx := authremote.ContextWithTimeout(time.Duration(1500) * time.Millisecond)
			_, err := ls.LogActivity(ctx, r.req)
			if err == nil {
				r.done = true

			}

		}

		// now remove all those that are "done" or broken from our retry queue
		retry_lock.Lock()
		var rsn []*logreq
		for _, r := range retries {
			if r.done || r.counter > 1000 {
				continue
			}
			rsn = append(rsn, r)
		}
		retries = rsn
		retry_lock.Unlock()

	}
}
