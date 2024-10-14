package practice12_11

import (
	"net/url"
	"reflect"
	"testing"
)

func TestPack(t *testing.T) {
	type args struct {
		given *url.URL
		ptr   any
	}
	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		{
			"string",
			args{
				&url.URL{},
				struct {
					typ string `http:"type"`
				}{
					"a",
				},
			},
			&url.URL{
				RawQuery: "type=a",
			},
			false,
		},
		{
			"int",
			args{
				&url.URL{},
				struct {
					typ int `http:"type"`
				}{
					32,
				},
			},
			&url.URL{
				RawQuery: "type=32",
			},
			false,
		},
		{
			"float",
			args{
				&url.URL{},
				struct {
					typ float64 `http:"type"`
				}{
					12.5,
				},
			},
			&url.URL{
				RawQuery: "type=12.5",
			},
			false,
		},
		{
			"bool",
			args{
				&url.URL{},
				struct {
					typ bool `http:"type"`
				}{
					false,
				},
			},
			&url.URL{
				RawQuery: "type=false",
			},
			false,
		},
		{
			"ptr",
			args{
				&url.URL{},
				struct {
					typ *string `http:"type"`
				}{
					ptr("abc"),
				},
			},
			&url.URL{
				RawQuery: "type=abc",
			},
			false,
		},
		{
			"multiple",
			args{
				&url.URL{},
				struct {
					typ string `http:"type"`
					val int    `http:"value"`
				}{
					"abc",
					123,
				},
			},
			&url.URL{
				RawQuery: "type=abc&value=123",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Pack(tt.args.given, tt.args.ptr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pack() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
