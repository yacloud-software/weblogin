package common

import (
	"fmt"
)

type WError struct {
	msg    string
	Urgent bool
}

func UrgentErrorf(format string, args ...interface{}) *WError {
	return &WError{
		Urgent: true,
		msg:    fmt.Sprintf(format, args...),
	}
}
func Errorf(format string, args ...interface{}) *WError {
	return &WError{
		msg: fmt.Sprintf(format, args...),
	}
}

func (w *WError) Error() string {
	return w.msg
}
