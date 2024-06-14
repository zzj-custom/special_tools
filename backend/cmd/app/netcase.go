package app

import (
	"special_tools/backend/cmd/dto/netcase"
	"special_tools/backend/internal/response"
)

type Netcase struct{}

func (n Netcase) ParseInfo(req []*netcase.ParseInfoDTO) *response.Reply {
	m, err := netcase.ParseInfo(req)
	if err != nil {
		return response.FailReply(response.ParseMetaInfoFail)
	}
	return response.OkReply(m)
}

func (n Netcase) Process(req *netcase.ProcessDTO) *response.Reply {
	err := netcase.Process(req)
	if err != nil {
		return response.FailReply(response.ProcessFail)
	}
	return response.OkReply(nil)
}
