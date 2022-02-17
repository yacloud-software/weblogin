package requests

import (
	"context"
	"fmt"
	"golang.conradwood.net/apis/themes"
	pb "golang.conradwood.net/apis/weblogin"
	"strings"
)

func serveThemes(ctx context.Context, cr *Request) (*pb.WebloginResponse, error) {
	req := cr.req
	host := req.Host
	state, err := cr.getState(ctx)
	if err != nil {
		fmt.Printf("[themes] Error getting state: %s\n", err)
	}
	if err == nil && state != nil {
		host = state.TriggerHost
	}

	fmt.Printf("[themes] Coming from host \"%s\"\n", host)
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
