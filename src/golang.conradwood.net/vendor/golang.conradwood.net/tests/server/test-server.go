package main

import (
	"context"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/echoservice"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/tokens"
	"golang.conradwood.net/go-easyops/utils"
	"google.golang.org/grpc"
	"time"
)

var (
	port      = flag.Int("port", 4106, "The grpc server port")
	ping      = flag.Bool("ping", false, "ping continously")
	ping_once = flag.Bool("ping_once", false, "ping once")
	ctr       = 0
)

// create a simple standard server
type echoServer struct {
}

func main() {
	flag.Parse()
	fmt.Printf("GO-EASYOPS Echo test server/client\n")
	if *ping || *ping_once {
		c := pb.GetEchoClient()
		for {
			now := time.Now()
			ctx := authremote.Context()
			ctx = tokens.ContextWithToken()
			ctx = authremote.Context()
			ctx = authremote.Context()
			u := auth.GetUser(ctx)
			fmt.Printf("   pinging as %s\n", auth.Description(u))
			_, err := c.Ping(ctx, &common.Void{})
			if err != nil {
				fmt.Printf("Error :%s\n", utils.ErrorString(err))
			}
			dur := time.Since(now).Milliseconds()
			fmt.Printf("%d Pinged (%d milliseconds)\n", ctr, dur)
			ctr++
			if !*ping {
				return
			}
			time.Sleep(time.Duration(300) * time.Millisecond)
		}
	}
	p := *port
	p = p + utils.RandomInt(50)
	sd := server.NewServerDef()
	sd.Port = p
	sd.Register = server.Register(
		func(g *grpc.Server) error {
			pb.RegisterEchoServiceServer(g, &echoServer{})
			return nil
		},
	)
	err := server.ServerStartup(sd)
	//	err := create.NewEchoServiceServer(&echoServer{}, p)
	utils.Bail("Unable to start server", err)
}

func (e *echoServer) Ping(ctx context.Context, req *common.Void) (*pb.PingResponse, error) {
	u := auth.GetUser(ctx)
	fmt.Printf("    %d Pinged by %s\n", ctr, auth.Description(u))
	ctr++
	i := utils.RandomInt(10)
	if i > 3 {
		return nil, errors.Unavailable(ctx, "Ping()")
	}
	return &pb.PingResponse{}, nil
}
