package response

type ErrCode int

//go:generate stringer -type=ErrCode -linecomment

const (
	Ok               ErrCode = 0      // ok
	QueryParamsError ErrCode = 100002 // 请求参数错误，请检查后重试！
)

// 登录
const (
	ValidateCodeFail ErrCode = 100001 // 验证验证码错误
	SendCodeFail     ErrCode = 100002 // 发送验证码错误
	CodeRepeat       ErrCode = 100003 // 验证码已发送，请稍后重试
)

// 视频下载
const (
	AcquiredVideoList ErrCode = 200001 // 获取视频列表数据失败
	DownloadError     ErrCode = 200002 // 下载失败
)

// 必应

const (
	AcquiredImagesFailed ErrCode = 300001 // 获取必应图片失败
)

// 网易云转换
const (
	ParseMetaInfoFail ErrCode = 400000 // 解析音乐信息失败
	ProcessFail       ErrCode = 400001 // 网易云转换失败
)
