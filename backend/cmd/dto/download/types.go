package download

type Downloader struct {
	U        string `json:"u"`        // 下载地址
	OutPath  string `json:"outPath"`  // 保存路径
	PlayList bool   `json:"playList"` // 是否是播放列表
	Items    string `json:"items"`    // 下载范围，定义例如：1,5,6,8-10
}
