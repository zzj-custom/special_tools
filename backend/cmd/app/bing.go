package app

import (
	"special_tools/backend/cmd/dto/bing"
	"special_tools/backend/internal/response"
)

type Bing struct{}

func (a *Bing) Images() *response.Reply {
	return bing.List()
}
