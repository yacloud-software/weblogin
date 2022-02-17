package requests

import (
	"context"
	"golang.conradwood.net/apis/themes"
	pb "golang.conradwood.net/apis/weblogin"
	"strings"
)

func serveThemes(ctx context.Context, req *pb.WebloginRequest) (*pb.WebloginResponse, error) {
	path := strings.TrimSuffix(req.Path, "?")
	if strings.HasSuffix(path, "stylesheet.css") {
		cs, err := themes.GetThemesClient().GetCSS(ctx, &themes.HostThemeRequest{Host: req.Host})
		if err != nil {
			return nil, err
		}
		return &pb.WebloginResponse{Body: cs.Data, MimeType: "text/css"}, nil
	}
	if strings.HasSuffix(path, "logo.png") {
		cs, err := themes.GetThemesClient().GetLogo(ctx, &themes.HostThemeRequest{Host: req.Host})
		if err != nil {
			return nil, err
		}
		return &pb.WebloginResponse{Body: cs.Data, MimeType: cs.MimeType}, nil
	}
	if strings.HasSuffix(path, "favicon.ico") {
		cs, err := themes.GetThemesClient().GetFavIcon(ctx, &themes.HostThemeRequest{Host: req.Host})
		if err != nil {
			return nil, err
		}
		return &pb.WebloginResponse{Body: cs.Data, MimeType: cs.MimeType}, nil
	}
	return nil, nil
}
