package requests

import (
	"context"
	pb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/utils"
	"mime"
	"path/filepath"
)

func (r *Request) ServeAsset(ctx context.Context, assetname string) (*pb.WebloginResponse, error) {
	fname := "templates/v2/assets/" + utils.MakeSafeFilename(assetname)
	f, err := utils.FindFileInWorkingDir(fname)
	//f, err := utils.FindFile(fname)
	if err != nil {
		r.Debugf("asset \"%s\" not found\n", fname)
		return nil, errors.NotFound(ctx, "not found", "asset not found: %s", assetname)
	}
	b, err := utils.ReadFile(f)
	if err != nil {
		return nil, err
	}
	mt := mime.TypeByExtension(filepath.Ext(assetname))
	wr := &pb.WebloginResponse{
		Body:     b,
		MimeType: mt,
	}
	return wr, nil
}
