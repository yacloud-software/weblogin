package requests

import (
	"fmt"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/weblogin/requesttracker"
)

func needSessionPage(cr *requesttracker.Request) (*pb.WebloginResponse, error) {
	return nil, fmt.Errorf("need session not implemented")
}
