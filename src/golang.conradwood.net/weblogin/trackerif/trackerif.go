package trackerif

import (
	pb "golang.conradwood.net/apis/weblogin"
)

// these functions cannot throw errors, they are in the critical path
// and must not interfere with the user path
type TrackerIF interface {
	LogActivity(action pb.AuthAction, req *pb.AuthActivityRequest)
}
