package requesttracker

import (
	"context"
	"fmt"
	"golang.conradwood.net/weblogin/trackerif"
)

type StandardTracker struct {
}

func (st *StandardTracker) UnspecifiedRequest(ctx context.Context, ti trackerif.TrackerInfo) {
	fmt.Printf("[standardtracker unspecified] %#v\n", ti)
}
func (st *StandardTracker) LoginRendered(ctx context.Context, ti trackerif.TrackerInfo) {
	fmt.Printf("[standardtracker loginrendered] %#v\n", ti)

}
func (st *StandardTracker) EmailVerified(ctx context.Context, ti trackerif.TrackerInfo)         {}
func (st *StandardTracker) RegistrationSubmitted(ctx context.Context, ti trackerif.TrackerInfo) {}
