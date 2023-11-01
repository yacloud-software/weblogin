package trackerif

import (
	"context"
	"golang.conradwood.net/apis/auth"
)

// these functions cannot throw errors, they are in the critical path
// and must not interfere with the user path
type TrackerIF interface {
	UnspecifiedRequest(ctx context.Context, ti TrackerInfo)
	LoginRendered(ctx context.Context, ti TrackerInfo)
	EmailVerified(ctx context.Context, ti TrackerInfo)
	RegistrationSubmitted(ctx context.Context, ti TrackerInfo)
}

type TrackerInfo interface {
	GetIP() string
	GetPreviousUser() *auth.User
	GetCurrentUser() *auth.User
	GetAuthEmail() string
	GetURL() string
	GetError() error
}
