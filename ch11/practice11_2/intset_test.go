package practice11_2

import "testing"

func TestIntSet_Has(t *testing.T) {
	type fields struct {
		Words []uint64
	}
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"has",
			fields{Words: []uint64{
				2, // 10
			}},
			args{
				1,
			},
			true,
		},
		{
			"has over 64",
			fields{Words: []uint64{
				0, // 0
				1, // 1
			}},
			args{
				64,
			},
			true,
		},
		{
			"not has",
			fields{Words: []uint64{
				2, // 10
			}},
			args{
				0,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{
				Words: tt.fields.Words,
			}
			if got := s.Has(tt.args.x); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_Add(t *testing.T) {
	type fields struct {
		Words []uint64
	}
	type args struct {
		x int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantState []uint64
	}{
		{
			"add to blank",
			fields{
				Words: []uint64{},
			},
			args{
				x: 1,
			},
			[]uint64{
				2, // 10
			},
		},
		{
			"add same",
			fields{
				Words: []uint64{
					2, // 10
				},
			},
			args{
				x: 1,
			},
			[]uint64{
				2, // 10
			},
		},
		{
			"add another",
			fields{
				Words: []uint64{
					2, // 10
				},
			},
			args{
				x: 0,
			},
			[]uint64{
				3, // 11
			},
		},
		{
			"add over 64",
			fields{
				Words: []uint64{
					2, // 10
				},
			},
			args{
				x: 64,
			},
			[]uint64{
				2, // 10
				1, // 01
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{
				Words: tt.fields.Words,
			}
			s.Add(tt.args.x)

			if !assertEqualsWords(tt.wantState, s.Words) {
				t.Errorf("State afterAdd() is %v, want %v", s.Words, tt.wantState)
			}
		})
	}
}

func assertEqualsWords(want, got []uint64) bool {
	if len(want) != len(got) {
		return false
	}

	for i := range got {
		if want[i] != got[i] {
			return false
		}
	}

	return true
}

func TestIntSet_UnionWith(t *testing.T) {
	type fields struct {
		Words []uint64
	}
	type args struct {
		t *IntSet
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantState *IntSet
	}{
		{
			"{0,1} + {2,3} = {0,1,2,3}",
			fields{
				Words: []uint64{
					3, // 0011
				},
			},
			args{
				&IntSet{
					Words: []uint64{
						12, // 1100
					}},
			},
			&IntSet{
				Words: []uint64{
					15, // 1111
				},
			},
		},
		{
			"{0,1} + {1,2} = {0,1,2}",
			fields{
				Words: []uint64{
					3, // 011
				},
			},
			args{
				&IntSet{
					Words: []uint64{
						6, // 110
					}},
			},
			&IntSet{
				Words: []uint64{
					7, // 111
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{
				Words: tt.fields.Words,
			}
			s.UnionWith(tt.args.t)

			if !assertEqualsWords(tt.wantState.Words, s.Words) {
				t.Errorf("State after UnionWith() is %v, want %v", s.Words, tt.wantState)
			}
		})
	}
}

func TestIntSet_String(t *testing.T) {
	type fields struct {
		Words []uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"empty",
			fields{
				Words: []uint64{},
			},
			"{}",
		},
		{
			"has single element",
			fields{
				Words: []uint64{
					2, // 10
				},
			},
			"{1}",
		},
		{
			"has some elements",
			fields{
				Words: []uint64{
					3, // 11
				},
			},
			"{0 1}",
		},
		{
			"has over 64",
			fields{
				Words: []uint64{
					0,
					1,
				},
			},
			"{64}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{
				Words: tt.fields.Words,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
