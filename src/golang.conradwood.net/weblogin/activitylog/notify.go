package activitylog

import (
	"fmt"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
	"golang.yacloud.eu/apis/messaging"
)

func send_notification(format string, args ...interface{}) {
	text := fmt.Sprintf(format, args...)
	p := &messaging.TopicMessageRequest{Topic: "weblogin", Message: text}
	go func(pr *messaging.TopicMessageRequest) {
		ctx := authremote.Context()
		_, err := messaging.GetMessagingClient().MessageToTopic(ctx, pr)
		if err != nil {
			fmt.Printf("Failed to slack \"%s\" to \"%s\": %s", pr.Message, "none", utils.ErrorString(err))
		}
	}(p)
}
