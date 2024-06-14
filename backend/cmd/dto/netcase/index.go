package netcase

import (
	"encoding/base64"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"special_tools/backend/pkg/ncm"
	"strings"
)

var (
	p = "/Users/Apple/Application/github/special_tools/file"
)

func ParseInfo(req []*ParseInfoDTO) ([]*ConvertResponse, error) {
	m := make([]*ConvertResponse, 0, len(req))
	for _, item := range req {
		logrus.WithFields(logrus.Fields{
			"name": item.Name,
		}).Info("file content")
		p, err := handleFile(item)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to handle file")
		}
		n := ncm.NewNcm(p, "")
		mi, err := n.ParseMateInfo()
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to parse mate info")
		}
		m = append(m, &ConvertResponse{
			Name:          item.Name,
			MusicID:       mi.MusicID,
			MusicName:     mi.MusicName,
			Artist:        mi.Artist,
			AlbumID:       mi.AlbumID,
			Album:         mi.Album,
			AlbumPicDocID: mi.AlbumPic,
			AlbumPic:      mi.AlbumPic,
			BitRate:       mi.BitRate,
			Mp3DocID:      mi.Mp3DocID,
			Duration:      mi.Duration,
			MvID:          mi.MvID,
			Alias:         mi.Alias,
			TransNames:    mi.TransNames,
			Format:        mi.Format,
		})
	}
	return m, nil
}

func handleFile(req *ParseInfoDTO) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(strings.Split(req.Content, ",")[1])
	if err != nil {
		return "", errors.Wrapf(err, "Failed to file content decode base64")
	}
	filePath := filepath.Join(p, req.Name)
	out, err := os.Create(filePath)
	if err != nil {
		return "", errors.Wrapf(err, "Failed to create file")
	}
	defer func() { _ = out.Close() }()

	_, err = out.Write(decoded)
	if err != nil {
		return "", errors.Wrapf(err, "Failed to write file")
	}

	logrus.WithFields(logrus.Fields{
		"path": filePath,
	}).Info("file path")

	return filePath, nil
}

func Process(req *ProcessDTO) error {
	for _, item := range req.Path {
		if err := ncm.NewNcm(filepath.Join(p, item), req.OutPath).Process(); err != nil {
			return errors.Wrapf(err, "Failed to process file[%s]", item)
		}
	}
	return nil
}
