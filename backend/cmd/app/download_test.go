package app

import (
	"context"
	"reflect"
	"special_tools"
	"special_tools/backend/internal/response"
	"special_tools/backend/pkg/extractor"
	"special_tools/backend/pkg/extractor/bilibili"
	"testing"
)

func init() {
	extractor.NewManagerDownloader().Register("bilibili", &bilibili.Bili{})
	extractor.NewManagerDownloader().Register("b23", &bilibili.Bili{})
}

func TestApp_VideoList(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	type args struct {
		u        string
		playList bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *response.Reply
	}{
		{
			name: "测试视频列表",
			fields: fields{
				ctx: context.TODO(),
			},
			args: args{
				u:        "https://www.bilibili.com/video/BV1dS4y1y7vd",
				playList: true,
			},
			want: nil,
		},
		{
			name: "测试视频单个",
			fields: fields{
				ctx: context.TODO(),
			},
			args: args{
				u:        "https://www.bilibili.com/video/BV1dS4y1y7vd",
				playList: false,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &main.App{
				ctx: tt.fields.ctx,
			}
			if got := a.VideoList(tt.args.u, tt.args.playList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VideoList() = %v, want %v", got, tt.want)
			}
		})
	}
}
