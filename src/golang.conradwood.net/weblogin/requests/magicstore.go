package requests

import (
	"context"
	"fmt"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/tokens"
	"golang.conradwood.net/go-easyops/utils"
	"time"
)

const (
	prefix = "WEBLOGIN_"
)

// add a 'magic' to submitted parameters
func generateMagicIfNecessary(cr *Request) error {
	if cr.req.Submitted == nil {
		cr.req.Submitted = make(map[string]string)
	}
	if len(cr.req.Submitted[WEBLOGIN_STATE]) >= 10 {
		return nil
	}
	magic, _, err := createState(cr)
	if err != nil {
		return err
	}
	cr.req.Submitted[WEBLOGIN_STATE] = magic
	return nil
}

func (cr *Request) putMagic(magic string, state *pb.State) error {
	ctx := tokens.ContextWithToken()
	state_string, err := utils.Marshal(state)
	if err != nil {
		return err
	}
	if state_string == "" {
		panic("state_string is empty. cannot create a state without ANY information")
	}
	exp := time.Now().Add(time.Duration(180) * time.Second)
	err = client.PutWithIDAndExpiry(ctx, prefix+magic, []byte(state_string), exp)
	if err != nil {
		cr.Debugf("Failed to put magic (%s): %s\n", magic, err)
	} else {
		cr.Debugf("Stored magic %s\n", magic)
	}
	return err
}
func (cr *Request) getState(ctx context.Context) (*pb.State, error) {
	magic := cr.req.Submitted[WEBLOGIN_STATE]
	return cr.getMagic(ctx, magic)
}
func (cr *Request) getMagic(ctx context.Context, magic string) (*pb.State, error) {
	if len(magic) < 10 {
		cr.Debugf("invalid magic %s: length < 10 (%d)\n", magic, len(magic))
		lm := fmt.Sprintf("invalid magic %s: length < 10 (%d)\n", magic, len(magic))
		return nil, errors.InvalidArgs(ctx, "this url is invalid or has expired", lm)
	}

	o, err := client.Get(ctx, prefix+magic)
	if err != nil {
		cr.Debugf("error getting magic %s: %s\n", magic, err)
		return nil, err
	}
	st := &pb.State{}
	err = utils.Unmarshal(string(o), st)
	if err != nil {
		cr.Debugf("error unmarshalling magic %s: %s\n", magic, err)
		return nil, errors.InvalidArgs(ctx, "invalid state", "[getmagic] invalid state for valid magic (%s)", magic)
	}
	return st, nil
}
