package requests

import (
	"context"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/themes"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/weblogin/requesttracker"
	"strings"
)

var (
	debug_themes = flag.Bool("debug_themes", false, "debug themes stuff")
)

func serveThemes(ctx context.Context, cr *requesttracker.Request) (*pb.WebloginResponse, error) {
	req := cr.Request()
	host := req.Host
	state, err := getState(ctx, cr)
	if err != nil {
		fmt.Printf("[themes] Error getting state: %s\n", err)
	}
	if err == nil && state != nil {
		host = state.TriggerHost
	}
	if *debug_themes {
		fmt.Printf("[themes] Coming from host \"%s\"\n", host)
	}
	path := strings.TrimSuffix(req.Path, "?")
	if strings.HasSuffix(path, "stylesheet.css") {
		cs, err := themes.GetThemesClient().GetCSS(ctx, &themes.HostThemeRequest{Host: host})
		if err != nil {
			return nil, err
		}
		return &pb.WebloginResponse{Body: cs.Data, MimeType: "text/css"}, nil
	}
	if strings.HasSuffix(path, "logo.png") {
		cs, err := themes.GetThemesClient().GetLogo(ctx, &themes.HostThemeRequest{Host: host})
		if err != nil {
			return nil, err
		}
		return &pb.WebloginResponse{Body: cs.Data, MimeType: cs.MimeType}, nil
	}
	if strings.HasSuffix(path, "favicon.ico") {
		cs, err := themes.GetThemesClient().GetFavIcon(ctx, &themes.HostThemeRequest{Host: host})
		if err != nil {
			return nil, err
		}
		return &pb.WebloginResponse{Body: cs.Data, MimeType: cs.MimeType}, nil
	}
	return nil, nil
}
