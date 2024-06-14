package app

import (
	"github.com/go-playground/validator/v10"
	"special_tools/backend/cmd/dto/download"
	"special_tools/backend/internal/response"
)

type Download struct {
}

func (a *Download) VideoList(u string, playList bool) *response.Reply {
	if err := validator.New().Var(u, "required,url"); err != nil {
		return response.FailReply(response.QueryParamsError)
	}

	//TODO 验证是否是能下载的url地址
	d := &download.Downloader{
		U:        u,
		PlayList: playList,
	}
	return d.VideoList()
}

func (a *Download) Download(u string, outPath string, playList bool, items string) *response.Reply {
	if err := validator.New().Var(u, "required,url"); err != nil {
		return response.FailReply(response.QueryParamsError)
	}

	//TODO 验证是否是能下载的url地址
	d := &download.Downloader{
		U:        u,
		OutPath:  outPath,
		PlayList: playList,
		Items:    items,
	}
	return d.Download()
}
