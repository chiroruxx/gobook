package practice11_5

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{
				"a:b:c", ":",
			},
			3,
		},
	}

	for _, test := range tests {
		words := strings.Split(test.args.s, test.args.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.args.s, test.args.sep, got, test.want)
		}
	}
}
