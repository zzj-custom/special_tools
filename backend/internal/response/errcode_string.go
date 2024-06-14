// Code generated by "stringer -type=ErrCode -linecomment"; DO NOT EDIT.

package response

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Ok-0]
	_ = x[QueryParamsError-100002]
	_ = x[ValidateCodeFail-100001]
	_ = x[SendCodeFail-100002]
	_ = x[CodeRepeat-100003]
	_ = x[AcquiredVideoList-200001]
	_ = x[DownloadError-200002]
	_ = x[AcquiredImagesFailed-300001]
	_ = x[ParseMetaInfoFail-400000]
	_ = x[ProcessFail-400001]
}

const (
	_ErrCode_name_0 = "ok"
	_ErrCode_name_1 = "验证验证码错误请求参数错误，请检查后重试！验证码已发送，请稍后重试"
	_ErrCode_name_2 = "获取视频列表数据失败下载失败"
	_ErrCode_name_3 = "获取必应图片失败"
	_ErrCode_name_4 = "解析音乐信息失败网易云转换失败"
)

var (
	_ErrCode_index_1 = [...]uint8{0, 21, 63, 99}
	_ErrCode_index_2 = [...]uint8{0, 30, 42}
	_ErrCode_index_4 = [...]uint8{0, 24, 45}
)

func (i ErrCode) String() string {
	switch {
	case i == 0:
		return _ErrCode_name_0
	case 100001 <= i && i <= 100003:
		i -= 100001
		return _ErrCode_name_1[_ErrCode_index_1[i]:_ErrCode_index_1[i+1]]
	case 200001 <= i && i <= 200002:
		i -= 200001
		return _ErrCode_name_2[_ErrCode_index_2[i]:_ErrCode_index_2[i+1]]
	case i == 300001:
		return _ErrCode_name_3
	case 400000 <= i && i <= 400001:
		i -= 400000
		return _ErrCode_name_4[_ErrCode_index_4[i]:_ErrCode_index_4[i+1]]
	default:
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
