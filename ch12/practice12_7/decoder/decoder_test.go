package decoder

import (
	"io"
	"strings"
	"testing"
)

func TestDecoder_Decode(t *testing.T) {
	type person struct {
		Name string
		Age  int
	}

	type typeType uint
	const (
		typeInt typeType = iota
		typeString
		typeStruct
		typeList
		typeMap
	)

	type fields struct {
		r io.Reader
	}
	tests := []struct {
		name     string
		fields   fields
		typeType typeType
		want     any
		wantErr  bool
	}{
		{
			name: "number",
			fields: fields{
				r: strings.NewReader("123"),
			},
			typeType: typeInt,
			want:     123,
			wantErr:  false,
		},
		{
			name: "string",
			fields: fields{
				r: strings.NewReader(`"abc"`),
			},
			typeType: typeString,
			want:     "abc",
			wantErr:  false,
		},
		{
			name: "struct",
			fields: fields{
				r: strings.NewReader(`((Name "John") (Age 18))`),
			},
			typeType: typeStruct,
			want: person{
				Name: "John",
				Age:  18,
			},
			wantErr: false,
		},
		{
			name: "list",
			fields: fields{
				r: strings.NewReader(`(1 2 3)`),
			},
			typeType: typeList,
			want:     []int{1, 2, 3},
			wantErr:  false,
		},
		{
			name: "map",
			fields: fields{
				r: strings.NewReader(`((Name "John") (Address "Tokyo"))`),
			},
			typeType: typeMap,
			want: map[string]string{
				"Name":    "John",
				"Address": "Tokyo",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Decoder{
				r: tt.fields.r,
			}
			switch tt.typeType {
			case typeInt:
				var n int
				if err := d.Decode(&n); (err != nil) != tt.wantErr {
					t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				}
				if tt.want != n {
					t.Errorf("Decode() addr = %v, want %v", n, tt.want)
				}
			case typeString:
				var s string
				if err := d.Decode(&s); (err != nil) != tt.wantErr {
					t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				}
				if tt.want != s {
					t.Errorf("Decode() addr = %v, want %v", s, tt.want)
				}
			case typeStruct:
				var s person
				if err := d.Decode(&s); (err != nil) != tt.wantErr {
					t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				w := tt.want.(person)
				if w.Name != s.Name {
					t.Errorf("Decode() name = %v, want %v", s.Name, w.Name)
				}
				if w.Age != s.Age {
					t.Errorf("Decode() age = %v, want %v", s.Age, w.Age)
				}
			case typeList:
				l := make([]int, 3)
				if err := d.Decode(&l); (err != nil) != tt.wantErr {
					t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				wl := tt.want.([]int)
				if len(wl) != len(l) {
					t.Errorf("Decode() len = %v, want %v", len(l), len(wl))
					return
				}
				for i, v := range wl {
					if v != l[i] {
						t.Errorf("Decode() %d, v = %v, want %v", i, l[i], v)
					}
				}
			case typeMap:
				m := make(map[string]string)
				if err := d.Decode(&m); (err != nil) != tt.wantErr {
					t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				w := tt.want.(map[string]string)
				if w["Name"] != m["Name"] {
					t.Errorf("Decode() name = %v, want %v", m["Name"], w["Name"])
				}
				if w["Age"] != m["Age"] {
					t.Errorf("Decode() age = %v, want %v", m["Age"], w["Age"])
				}
			default:
				panic("unhandled typeType")
			}
		})
	}
}
