package activitylog

import (
	"fmt"
	slack "golang.conradwood.net/apis/slackgateway"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
)

func send_notification(format string, args ...interface{}) {
	text := fmt.Sprintf(format, args...)
	p := &slack.PostRequest{UserID: "7", Text: text}
	go func(pr *slack.PostRequest) {
		ctx := authremote.Context()
		_, err := slack.GetSlackGatewayClient().Post(ctx, pr)
		if err != nil {
			fmt.Printf("Failed to slack \"%s\" to \"%s\": %s", pr.Text, pr.UserID, utils.ErrorString(err))
		}
	}(p)
}
