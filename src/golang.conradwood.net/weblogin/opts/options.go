package opts

import (
	"flag"
)

var (
	debug              = flag.Bool("debug", false, "debug mode")
	userjourneyservice = flag.String("userjourneytracker", "userjourneytracker.UserJourneyTracker", "service tracking userjourneys")
)

func UserJourneyTracker() string {
	return *userjourneyservice
}
func IsDebug() bool {
	return *debug
}
