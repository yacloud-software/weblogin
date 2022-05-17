package activitylog

import (
	"context"
	pb "golang.conradwood.net/apis/weblogin"
	//	"golang.conradwood.net/weblogin/db"
	"fmt"
	"strings"
	"time"
)

type Logger struct {
	Email       string
	UserID      string
	IP          string
	TriggerHost string
	BrowserID   string
	UserAgent   string
}

func (l *Logger) Log(ctx context.Context, message string) {
	if strings.Contains(l.IP, "2001:ba8:1f1:f0bd::2") {
		// do not slack me on my own ones
		return
	}
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
	s := fmt.Sprintf("[weblogin,email=%s%s] %s", al.Email, ip, al.LogMessage)
	send_notification("%s", s)
	s = s + fmt.Sprintf(", useragent=%s, deviceid=%s", l.UserAgent, l.BrowserID)
	fmt.Println(s)
}
