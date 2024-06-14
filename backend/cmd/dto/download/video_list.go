package download

import (
	"github.com/sirupsen/logrus"
	"special_tools/backend/internal/downloader"
	"special_tools/backend/internal/response"
	"special_tools/backend/pkg/extractor"
)

func (d *Downloader) Videos() ([]*extractor.Data, error) {
	options := extractor.Options{
		Playlist: d.PlayList,
		Items:    d.Items,
	}
	return extractor.Dispatch(d.U, options)
}

func (d *Downloader) VideoList() *response.Reply {
	// TODO 根据传参构建options
	videos, err := d.Videos()
	if err != nil {
		logrus.WithField("url", d.U).WithError(err).Error("get video list error")
		return response.FailReply(response.AcquiredVideoList)
	}
	return response.OkReply(videos)
}

func (d *Downloader) Download() *response.Reply {
	videos, err := d.Videos()
	if err != nil {
		logrus.WithField("url", d.U).WithError(err).Error("get video list error")
		return response.FailReply(response.AcquiredVideoList)
	}
	if len(videos) == 0 {
		return response.FailReply(response.AcquiredVideoList)
	}

	for _, item := range videos {
		down := downloader.NewDownloader(&downloader.Options{
			OutputPath: d.OutPath,
		})

		if down.Loader(item) != nil {
			logrus.WithField("url", d.U).WithError(err).Error("download video error")
		}
	}
	return response.OkReply(videos)
}
