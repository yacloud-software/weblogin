package requests

import (
	"context"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/weblogin/common"
	"golang.conradwood.net/weblogin/requesttracker"
	"time"
)

// add a 'magic' to submitted parameters
func generateMagicIfNecessary(cr *requesttracker.Request) error {
	if cr.Request().Submitted == nil {
		cr.Request().Submitted = make(map[string]string)
	}
	if len(cr.Request().Submitted[common.WEBLOGIN_STATE]) >= 10 {
		return nil
	}
	magic, _, err := createState(cr)
	if err != nil {
		return err
	}
	cr.Request().Submitted[common.WEBLOGIN_STATE] = magic
	return nil
}

func putMagic(cr *requesttracker.Request, magic string, state *pb.State) error {
	ctx := authremote.Context()
	state_string, err := utils.Marshal(state)
	if err != nil {
		return err
	}
	if state_string == "" {
		panic("state_string is empty. cannot create a state without ANY information")
	}
	exp := time.Now().Add(time.Duration(180) * time.Second)
	err = client.PutWithIDAndExpiry(ctx, common.MAGIC_PREFIX+magic, []byte(state_string), exp)
	if err != nil {
		cr.Debugf("Failed to put magic (%s): %s\n", magic, err)
	} else {
		cr.Debugf("Stored magic %s\n", magic)
	}
	return err
}
func getState(ctx context.Context, cr *requesttracker.Request) (*pb.State, error) {
	magic := cr.Request().Submitted[common.WEBLOGIN_STATE]
	return getMagic(ctx, cr, magic)
}
func getMagic(ctx context.Context, cr *requesttracker.Request, magic string) (*pb.State, error) {
	return common.ParseMagic(ctx, magic)
}
