package activitylog

import (
	"context"
	pb "golang.conradwood.net/apis/weblogin"
	//	"golang.conradwood.net/weblogin/db"
	"fmt"
	"time"
)

type Logger struct {
	Email       string
	UserID      string
	IP          string
	TriggerHost string
}

func (l *Logger) Log(ctx context.Context, message string) {
	al := &pb.ActivityLog{
		IP:          l.IP,
		UserID:      l.UserID,
		Email:       l.Email,
		TriggerHost: l.TriggerHost,
		Occured:     uint32(time.Now().Unix()),
		LogMessage:  message,
	}
	//	db.DefaultDBActivityLog().Save(ctx, al)
	ip := ""
	if al.IP != "" {
		ip = fmt.Sprintf(", ip=%s", al.IP)
	}
	send_notification("[weblogin,email=%s%s] message=%s", al.Email, ip, al.LogMessage)

}
