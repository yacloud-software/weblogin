package common

import (
	"context"
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/utils"
	"time"
)

const (
	MAGIC_PREFIX   = "WEBLOGIN3_"
	WEBLOGIN_STATE = "weblogin3_state_yacloud"
)

var (
	allow_registrations = flag.Bool("v3_allow_registrations", true, "if true, allow registration of new users")
)

func Debugf(format string, args ...interface{}) {
	fmt.Printf("[common] "+format, args...)
}
func SaveMagic(ctx context.Context, magic string, state *pb.V3State) error {
	if len(magic) < 10 {
		Debugf("invalid magic %s: length < 10 (%d)\n", magic, len(magic))
		lm := fmt.Sprintf("invalid magic \"%s\": length < 10 (%d)\n", magic, len(magic))
		return errors.InvalidArgs(ctx, "this url is invalid or has expired", lm)
	}
	b, err := utils.MarshalBytes(state)
	if err != nil {
		return err
	}
	err = client.PutWithIDAndExpiry(ctx, MAGIC_PREFIX+magic, b, time.Now().Add(time.Duration(30)*time.Minute))
	if err != nil {
		return err
	}
	return nil
}
func ParseMagic(ctx context.Context, magic string) (*pb.V3State, error) {
	if len(magic) < 10 {
		Debugf("invalid magic %s: length < 10 (%d)\n", magic, len(magic))
		lm := fmt.Sprintf("invalid magic \"%s\": length < 10 (%d)\n", magic, len(magic))
		return nil, errors.InvalidArgs(ctx, "this url is invalid or has expired", lm)
	}

	o, err := client.Get(ctx, MAGIC_PREFIX+magic)
	if err != nil {
		Debugf("error getting magic \"%s\": %s\n", magic, err)
		return nil, err
	}
	st := &pb.V3State{}
	err = utils.UnmarshalBytes(o, st)
	if err != nil {
		Debugf("error unmarshalling magic \"%s\": %s\n", magic, err)
		return nil, errors.InvalidArgs(ctx, "invalid state", "[getmagic] invalid state for valid magic (%s)", magic)
	}
	return st, nil

}

func SSOHost() string {
	return "l.yacloud.eu"
}
