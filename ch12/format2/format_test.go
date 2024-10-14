package format2

import (
	"reflect"
	"testing"
)

func Test_formatComplex(t *testing.T) {
	type args struct {
		v      reflect.Value
		indent int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"slice",
			args{
				v: reflect.ValueOf([]int{
					1, 2, 3,
				}),
				indent: 0,
			},
			`slice [
  0: 1
  1: 2
  2: 3
]`,
		},
		{
			"slice nest",
			args{
				v: reflect.ValueOf([][]int{
					{1, 2},
					{3, 4},
				}),
				indent: 0,
			},
			`slice [
  0: slice [
    0: 1
    1: 2
  ]
  1: slice [
    0: 3
    1: 4
  ]
]`,
		},
		{
			"map",
			args{
				v: reflect.ValueOf(map[string]int{
					"a": 0,
					"b": 1,
				}),
				indent: 0,
			},
			`map {
  "a": 0
  "b": 1
}`,
		},
		{
			"struct",
			args{
				v: reflect.ValueOf(struct {
					name  string
					value string
				}{
					name:  "Name",
					value: "Value",
				}),
				indent: 0,
			},
			`struct {
  name: "Name"
  value: "Value"
}`,
		},
		{
			"pointer",
			args{
				v:      reflect.ValueOf(ptr(true)),
				indent: 0,
			},
			`*bool: true`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatComplex(tt.args.v, tt.args.indent); got != tt.want {
				t.Errorf("formatComplex() = \n%v\n, want \n%v\n", got, tt.want)
			}
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
