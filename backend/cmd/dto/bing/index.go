package bing

import (
	"special_tools/backend/internal/repository"
	"special_tools/backend/internal/response"
)

func List() *response.Reply {
	repo := repository.NewBingImagesRepository()
	result, err := repo.GetAll()
	if err != nil {
		return response.FailReply(response.AcquiredImagesFailed)
	}

	if len(result) == 0 {
		return response.OkReply(make([]*ListResponse, 0))
	}

	resp := make([]*ListResponse, 0, len(result))
	for _, item := range result {
		resp = append(resp, &ListResponse{
			Id:            item.Id,
			Name:          item.Name,
			Copyright:     item.Copyright,
			CopyrightLink: item.CopyrightLink,
			Url:           item.Url,
			Start:         item.Start,
			End:           item.End,
			Location:      item.Location,
			ClickCount:    item.ClickCount,
			DownloadCount: item.DownloadCount,
		})
	}

	return response.OkReply(result)
}
