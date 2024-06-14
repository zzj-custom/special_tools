package code

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"special_tools/backend/pkg/iRedis"
	"strconv"
)

type Code struct {
	Code int
}

func NewCode(code int) *Code {
	return &Code{
		Code: code,
	}
}

func (receiver Code) key() string {
	return "code:" + strconv.Itoa(receiver.Code)
}

func (receiver Code) Validate() (bool, error) {
	pool, err := iRedis.Pool()
	if err != nil {
		return false, errors.Wrapf(err, "获取redis连接池失败")
	}
	conn := pool.Get()
	defer func() { _ = conn.Close() }()

	ok, err := redis.Bool(conn.Do("EXISTS", receiver.key()))
	if err != nil {
		return false, errors.Wrapf(err, "判断验证码是否存在失败")
	}
	return ok, nil
}

func (receiver Code) Set() error {
	pool, err := iRedis.Pool()
	if err != nil {
		return errors.Wrapf(err, "获取redis连接池失败")
	}
	conn := pool.Get()
	defer func() { _ = conn.Close() }()

	_, err = conn.Do("SETEX", receiver.key(), receiver.Code, 120)
	if err != nil {
		return errors.Wrapf(err, "设置验证码失败")
	}
	return nil
}
