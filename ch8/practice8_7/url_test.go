package practice8_7

import (
	"net/url"
	"testing"
)

func Test_newUrl(t *testing.T) {
	type args struct {
		urlString  string
		requestUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "abs",
			args: args{
				"https://example.com",
				"",
			},
			want:    "https://example.com",
			wantErr: false,
		},
		{
			name: "rel top",
			args: args{
				"/",
				"https://example.com",
			},
			want:    "https://example.com",
			wantErr: false,
		},
		{
			name: "rel top 2",
			args: args{
				"/abc",
				"https://example.com",
			},
			want:    "https://example.com/abc",
			wantErr: false,
		},
		{
			name: "rel top 3",
			args: args{
				"/abc",
				"https://example.com/def",
			},
			want:    "https://example.com/abc",
			wantErr: false,
		},
		{
			name: "rel 1",
			args: args{
				"abc",
				"https://example.com",
			},
			want:    "https://example.com/abc",
			wantErr: false,
		},
		{
			name: "rel 2",
			args: args{
				"def",
				"https://example.com/abc",
			},
			want:    "https://example.com/abc/def",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqPURL, err := url.Parse(tt.args.requestUrl)
			if err != nil {
				t.Errorf("Cannot parse request url: %v", err)
			}

			reqURL := URL{value: *reqPURL}

			got, err := newURL(tt.args.urlString, &reqURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("newURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want != got.String() {
				t.Errorf("newURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
