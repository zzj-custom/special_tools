package bing

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"reflect"
	"special_tools/backend/cmd/bootstrapper"
	"special_tools/backend/internal/response"
	"special_tools/backend/pkg/iMysql"
	"testing"
)

func TestMain(m *testing.M) {
	cfg := bootstrapper.Bootstrap("../../../../build/config.toml", func(in fsnotify.Event) {})
	_, _ = iMysql.NewClient(cfg.Database)
	os.Exit(m.Run())
}

func TestList(t *testing.T) {
	tests := []struct {
		name string
		want *response.Reply
	}{
		// TODO: Add test cases.
		{
			name: "test",
			want: response.OkReply([]*ListResponse{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := List()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}
