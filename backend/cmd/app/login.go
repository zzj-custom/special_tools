package app

import (
	"special_tools/backend/cmd/dto/login"
	"special_tools/backend/internal/response"
)

type Login struct{}

func (a *Login) Code(req *login.CodeDTO) *response.Reply {
	return login.Code(req)
}
