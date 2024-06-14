package login

import (
	"github.com/go-playground/validator/v10"
	"log/slog"
	"special_tools/backend/cmd/bootstrapper/config"
	"special_tools/backend/internal/code"
	"special_tools/backend/internal/response"
	"special_tools/backend/internal/utils"
	"special_tools/backend/pkg/email"
)

func Code(req *CodeDTO) *response.Reply {
	if err := validator.New().Struct(req); err != nil {
		return response.FailReply(response.QueryParamsError)
	}

	// 验证码是否存在
	c := utils.GenerateRandomNumber(6)
	nc := code.NewCode(c)
	ok, err := nc.Validate()
	if err != nil {
		return response.FailReply(response.ValidateCodeFail)
	}

	if ok {
		return response.FailReply(response.CodeRepeat)
	}

	go func() {
		// 发送验证码
		e := email.NewEmail(config.Get().Email)
		if err = e.Send(
			req.To,
			c,
			email.WithOptionsAccount("zzj"),
			email.WithOptionsWeb("zzjlovetl"),
		); err != nil {
			slog.With(
				slog.String("to", req.To),
				slog.Int("code", c),
			).With("err", err).Error("send code error")
			return
		}

		// redis保存验证码
		if err = nc.Set(); err != nil {
			slog.With(
				slog.Int("code", c),
			).With("err", err).Error("set code error")
			return
		}
	}()

	return response.OkReply("123456")
}
