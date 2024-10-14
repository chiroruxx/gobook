package params

import (
	"net/http"
	"net/url"
	"testing"
)

func TestUnpack(t *testing.T) {
	type args struct {
		req *http.Request
		ptr interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "mail/success",
			args: args{
				req: &http.Request{
					Form: url.Values{
						"mail": []string{"test@example.com"},
					},
				},
				ptr: &struct {
					Mail string `http:"mail" validate:"mail"`
				}{},
			},
			wantErr: false,
		},
		{
			name: "mail/fail",
			args: args{
				req: &http.Request{
					Form: url.Values{
						"mail": []string{"abc"},
					},
				},
				ptr: &struct {
					Mail string `http:"mail" validate:"mail"`
				}{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Unpack(tt.args.req, tt.args.ptr); (err != nil) != tt.wantErr {
				t.Errorf("Unpack() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
