package email

import (
	"crypto/tls"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/gomail.v2"
	"log/slog"
	"os"
	"regexp"
	"sync"
)

const (
	emailRegex = "^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$"
)

var (
	email     *Email
	emailOnce sync.Once
)

func NewEmail(cfg *Config) *Email {
	emailOnce.Do(func() {
		email = new(Email)
		email.cfg = cfg
	})
	return email
}

func (e *Email) ValidateEmail(email []string) bool {
	if len(email) == 0 {
		slog.With("email", email).Error("邮箱不能为空")
		return false
	}

	// 编译正则表达式
	re := regexp.MustCompile(emailRegex)
	for _, s := range email {
		if !re.MatchString(s) {
			slog.With("email", s).Error("邮箱格式不正确")
			return false
		}
	}
	return true
}

func (e *Email) validate(str string, opts ...Option) bool {
	if !e.ValidateEmail([]string{str}) {
		return false
	}

	options := &Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.Cc != nil {
		if !e.ValidateEmail(options.Cc) {
			return false
		}
	}

	if options.Bcc != nil {
		if !e.ValidateEmail(options.Bcc) {
			return false
		}
	}

	if options.Attach != "" {
		if !e.fileExists(options.Attach) {
			return false
		}
	}

	if options.Video != "" {
		if !e.fileExists(options.Video) {
			return false
		}
	}

	if options.Image != "" {
		if !e.fileExists(options.Image) {
			return false
		}
	}

	e.Extend = options

	return true
}

// 检查文件是否存在
func (e *Email) fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			slog.With("filePath", filePath).Error("文件不存在")
			return false
		}
	}
	return true
}

func (e *Email) Send(to string, code int, opts ...Option) error {
	if !e.validate(to, opts...) {
		return errors.New("验证码发送失败")
	}

	var (
		m   = gomail.NewMessage()
		err error
	)

	// 设置请求header
	m = e.SetHeader(m, to)
	// 设置body
	m, err = e.SetBody(m, code)
	if err != nil {
		return errors.Wrap(err, "设置body失败")
	}
	// 设置附件
	m = e.Attach(m)
	// 发送邮件
	d := gomail.NewDialer(
		e.cfg.Host,
		e.cfg.Port,
		e.cfg.UserName,
		e.cfg.Password,
	)

	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	if err := d.DialAndSend(m); err != nil {
		return errors.Wrap(err, "发送邮件失败")
	}
	return nil
}

func (e *Email) SetHeader(m *gomail.Message, to string) *gomail.Message {
	m.SetHeader("From", e.cfg.UserName)
	//m.SetHeader("From", "alias"+"<zzjlovetl>") // 增加发件人别名

	m.SetHeader("To", to) // 收件人，可以多个收件人，但必须使用相同的 SMTP 连接
	if e.Extend.Cc != nil {
		m.SetHeader("Cc", e.Extend.Cc...) // 抄送，可以多个
	}

	if e.Extend.Bcc != nil {
		m.SetHeader("Bcc", e.Extend.Bcc...) // 密送，可以多个
	}

	if e.Extend.Subject != "" {
		m.SetHeader("Subject", e.Extend.Subject) // 邮件主题
	}
	return m
}

func (e *Email) SetBody(m *gomail.Message, code int) (*gomail.Message, error) {
	f, err := os.ReadFile("./template.html")
	if err != nil {
		slog.With("err", err).Error("读取模板文件失败")
		return m, errors.Wrapf(err, "读取模板文件失败")
	}

	var (
		account = e.Extend.Account
		web     = e.Extend.Web
	)

	if e.Extend.Account == "" {
		account = "系统用户"
	}

	if e.Extend.Web == "" {
		web = "zzjlovetl"
	}

	m.SetBody("text/html", fmt.Sprintf(string(f), account, code, web))
	return m, nil
}

func (e *Email) Attach(m *gomail.Message) *gomail.Message {
	if e.Extend.Attach != "" {
		m.Attach(e.Extend.Attach) // 附件文件，可以是文件，照片，视频等等
	}

	if e.Extend.Video != "" {
		m.Attach(e.Extend.Video) // 视频
	}

	if e.Extend.Image != "" {
		m.Attach(e.Extend.Image) // 图片
	}
	return m
}
