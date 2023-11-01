package opts

import (
	"flag"
)

var (
	debug = flag.Bool("debug", false, "debug mode")
)

func IsDebug() bool {
	return *debug
}
