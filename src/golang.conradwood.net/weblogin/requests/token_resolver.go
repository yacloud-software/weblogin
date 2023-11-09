package requests

import (
	"context"
	"fmt"
	au "golang.conradwood.net/apis/auth"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/common"
	"golang.conradwood.net/go-easyops/utils"
	"golang.yacloud.eu/apis/sessionmanager"
)

// if token is valid and resolves to a valid user, return it
func UserByToken(ctx context.Context, token string) (*au.User, error) {
	st := &sessionmanager.SessionToken{Token: token}
	sv, err := sessionmanager.GetSessionManagerClient().VerifySession(ctx, st)
	if err != nil {
		fmt.Printf("sessionmanager failure: %s\n", utils.ErrorString(err))
	} else {
		if sv.IsSessionToken {
			if !sv.IsValid {
				return nil, nil
			}
			us := common.VerifySignedUser(sv.User)
			if us != nil && !us.ServiceAccount { // don't check emailverified, it might not be yet (in weblogin)
				return us, nil
			}
		}
	}
	apr := &au.AuthenticateTokenRequest{Token: token}
	u, err := authremote.GetAuthClient().GetByToken(ctx, apr)
	if err != nil {
		return nil, err
	}
	if !u.Valid {
		return nil, nil
	}
	return u.User, nil
}
