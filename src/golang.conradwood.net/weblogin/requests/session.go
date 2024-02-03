package requests

import (
	"context"
	"fmt"
	"golang.conradwood.net/apis/h2gproxy"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/weblogin/requesttracker"
	as "golang.yacloud.eu/apis/session"
	"golang.yacloud.eu/apis/sessionmanager"
	"strings"
	"time"
)

const (
	SESSION_COOKIE_NAME = "Yei0neez1ohyohnith6iger6Oogexoo_sescook" // matches h2gproxy!

)

// this needs to check if there is a session (e.g. as cookie).
// if not, create a new one
// if there is, reuse it
// then, redirect browser back to where it came from (using the same method, payload etc)
// the last redirect needs to include some short-lived token, which h2gproxy can use to
// lookup the session
// (it is h2gproxy's job to then set the session cookie on the domain)
func needSessionPage(ctx context.Context, req *pb.WebloginRequest, cr *requesttracker.Request) (*pb.WebloginResponse, error) {
	idx := strings.LastIndex(req.Path, "/")
	if idx == -1 {
		return nil, fmt.Errorf("[needsession] no state in request")
	}

	u := auth.GetUser(ctx)
	uid := ""
	if u != nil {
		uid = u.ID
	}

	marshalled := req.Path[idx+1:]
	//	fmt.Printf("[needsession] Marshalled: \"%s\"\n", marshalled)
	state := &pb.SessionState{}
	err := utils.Unmarshal(marshalled, state)
	if err != nil {
		return nil, fmt.Errorf("[needsession] could not parse state: %w", err)
	}

	// builds the url it originally came from
	next_url := "https://" + state.TriggerHost + state.TriggerPath
	if state.TriggerQuery != "" {
		next_url = stripslash(next_url) + "?" + state.TriggerQuery
	}

	if state.NewSessionToken == "" {
		// builds the url to 	/weblogin/ on original host

		sessrequest := &sessionmanager.NewSessionRequest{}
		sr, err := sessionmanager.GetSessionManagerClient().NewSession(ctx, sessrequest)
		if err != nil {
			return nil, err
		}
		state.NewSessionToken = sr.Token

		s_state, err := utils.Marshal(state)
		if err != nil {
			return nil, err
		}
		next_url = "https://" + stripslash(state.TriggerHost) + "/weblogin/needsession/" + s_state
	}

	fmt.Printf("[needsession] Sending browser back to \"%s\"\n", next_url)

	session_expiry := time.Now().Add(time.Duration(30) * time.Minute)

	wr := &pb.WebloginResponse{
		Authenticated: false,
		User:          u,
		Token:         state.NewSessionToken,
		RedirectTo:    next_url,
		Cookies: []*h2gproxy.Cookie{
			&h2gproxy.Cookie{
				Name:   SESSION_COOKIE_NAME,
				Value:  state.NewSessionToken,
				Expiry: uint32(session_expiry.Unix()),
			},
		},
		ForceGetAfterRedirect: false,
		PeerIsDosing:          false,
		PeerIP:                req.Peer,
		HTTPCode:              0, // normal h2gproxy flow
		MimeType:              "application/binary",
		Session: &as.Session{
			SessionID: state.NewSessionToken,
			UserID:    uid,
		},
	}
	if u != nil {
		wr.Authenticated = true
	}
	return wr, nil
}

func stripslash(s string) string {
	for strings.HasSuffix(s, "/") {
		s = s[:len(s)-1]
	}
	return s
}
