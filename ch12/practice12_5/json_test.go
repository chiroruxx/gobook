package practice12_5

import (
	"encoding/json"
	"errors"
	"math"
	"reflect"
	"testing"
)

func TestMarshal(t *testing.T) {
	testInf := Test{}

	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
		Float           float64
		inf             Inf
	}

	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Strerling Handen",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Normin.)",
			"Best Adapt Screenplay (Normin.)",
			"Best Director (Normin.)",
			"Best Picture (Normin.)",
		},
		Float: math.Pi,
		inf:   testInf,
	}

	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"json",
			args{
				v: strangelove,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.args.v)
			want, wantErr := json.Marshal(tt.args.v)
			if !errors.Is(err, wantErr) {
				t.Errorf("Marshal() error = %v, wantErr %v", err, wantErr)
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Marshal()\ngot  = %s\nwant = %s", got, want)
			}
		})
	}
}
