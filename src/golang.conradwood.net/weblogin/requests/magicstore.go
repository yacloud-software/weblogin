package requests

import (
    "golang.conradwood.net/go-easyops/authremote"
	"context"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/weblogin/common"
	"time"
)

// add a 'magic' to submitted parameters
func generateMagicIfNecessary(cr *Request) error {
	if cr.req.Submitted == nil {
		cr.req.Submitted = make(map[string]string)
	}
	if len(cr.req.Submitted[common.WEBLOGIN_STATE]) >= 10 {
		return nil
	}
	magic, _, err := createState(cr)
	if err != nil {
		return err
	}
	cr.req.Submitted[common.WEBLOGIN_STATE] = magic
	return nil
}

func (cr *Request) putMagic(magic string, state *pb.State) error {
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
func (cr *Request) getState(ctx context.Context) (*pb.State, error) {
	magic := cr.req.Submitted[common.WEBLOGIN_STATE]
	return cr.getMagic(ctx, magic)
}
func (cr *Request) getMagic(ctx context.Context, magic string) (*pb.State, error) {
	return common.ParseMagic(ctx, magic)
}
