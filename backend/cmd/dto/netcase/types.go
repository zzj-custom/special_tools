package netcase

type ParseInfoDTO struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type ConvertResponse struct {
	Name          string          `json:"name"`          // 文件名称
	MusicID       int             `json:"musicId"`       // 音乐id
	MusicName     string          `json:"musicName"`     // 音乐名称
	Artist        [][]interface{} `json:"artist"`        // [[string,int],] 艺术家
	AlbumID       int             `json:"albumId"`       // 专辑id
	Album         string          `json:"album"`         // 专辑
	AlbumPicDocID interface{}     `json:"albumPicDocId"` // string or int 专辑图片id
	AlbumPic      string          `json:"albumPic"`      // 专辑图片
	BitRate       int             `json:"bitrate"`       // 音乐比特率
	Mp3DocID      string          `json:"mp3DocId"`      // 音乐文件id
	Duration      int             `json:"duration"`      // 音乐时长
	MvID          int             `json:"mvId"`          // 视频id
	Alias         []string        `json:"alias"`         // 音乐别名
	TransNames    []interface{}   `json:"transNames"`    // string or int 翻译名
	Format        string          `json:"format"`        // 音乐格式
}

type ProcessDTO struct {
	Path    []string `json:"path"`
	OutPath string   `json:"outPath"`
}
