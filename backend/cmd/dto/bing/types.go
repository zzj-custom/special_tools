package bing

type ListResponse struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	Copyright     string `json:"copyright"`
	CopyrightLink string `json:"copyrightLink"`
	Url           string `json:"url"`
	Start         string `json:"start"`
	End           string `json:"end"`
	Location      string `json:"location"`
	ClickCount    int    `json:"clickCount"`
	DownloadCount int    `json:"downloadCount"`
}
