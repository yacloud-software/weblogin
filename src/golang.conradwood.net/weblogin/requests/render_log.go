package requests

import (
	"fmt"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/weblogin/requesttracker"
)

func renderlog(cr *requesttracker.Request) (*pb.WebloginResponse, error) {
	res := NewWebloginResponse()
	page := cr.GetPara("page")
	button := cr.GetPara("button")
	update := cr.GetPara("update")
	// TODO: increase a metric of sort, for now just log it to verify it is compatible with html/js code
	fmt.Printf("[renderlog] %s %s %s\n", page, button, update)
	return res, nil
}
