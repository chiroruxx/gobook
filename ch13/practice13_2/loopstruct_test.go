package practice13_2

import "testing"

func TestIsLoopStruct(t *testing.T) {
	type link struct {
		value string
		tail  *link
	}
	a, b, c, d, e := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}, &link{value: "d"}, &link{value: "e"}
	a.tail, b.tail, c.tail = b, a, c
	e.tail = d

	type args struct {
		x any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"loop",
			args{
				a,
			},
			true,
		},
		{
			"loop",
			args{
				b,
			},
			true,
		},
		{
			"loop",
			args{
				c,
			},
			true,
		},
		{
			"not loop",
			args{
				d,
			},
			false,
		},
		{
			"not loop",
			args{
				e,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLoopStruct(tt.args.x); got != tt.want {
				t.Errorf("IsLoopStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
