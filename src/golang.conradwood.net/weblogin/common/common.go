package common

import (
	"context"
	"fmt"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/utils"
	"html/template"
)

const (
	MAGIC_PREFIX   = "WEBLOGIN_"
	WEBLOGIN_STATE = "weblogin_state_yacloud"
)

type Template_data interface {
	Username() string
	StateQuery() template.HTMLAttr
	ReferrerHost() string
	Heading() string
}

func Debugf(format string, args ...interface{}) {
	fmt.Printf("[common] "+format, args...)
}
func ParseMagic(ctx context.Context, magic string) (*pb.State, error) {
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
	st := &pb.State{}
	err = utils.Unmarshal(string(o), st)
	if err != nil {
		Debugf("error unmarshalling magic \"%s\": %s\n", magic, err)
		return nil, errors.InvalidArgs(ctx, "invalid state", "[getmagic] invalid state for valid magic (%s)", magic)
	}
	return st, nil

}
