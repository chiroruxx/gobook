package parser

import (
	"reflect"
	"testing"

	"gobook/ch12/practice12_7/lexer"
)

func Test_parser_parse(t *testing.T) {
	type args struct {
		tokens []lexer.Token
	}
	tests := []struct {
		name    string
		args    args
		want    Node
		wantErr bool
	}{
		{
			"number",
			args{
				tokens: []lexer.Token{
					lexer.NewNumberToken([]byte{'1'}),
				},
			},
			&NumberNode{
				token: lexer.NewNumberToken([]byte{'1'}),
			},
			false,
		},
		{
			"string",
			args{
				tokens: []lexer.Token{
					&lexer.QuoteToken{},
					lexer.NewStringToken([]byte{'a'}),
					&lexer.QuoteToken{},
				},
			},
			&StringNode{
				tokens: []lexer.Token{
					lexer.NewStringToken([]byte{'a'}),
				},
			},
			false,
		},
		{
			"symbol",
			args{
				tokens: []lexer.Token{
					lexer.NewStringToken([]byte{'a'}),
				},
			},
			&SymbolNode{
				tokens: []lexer.Token{
					lexer.NewStringToken([]byte{'a'}),
				},
			},
			false,
		},
		{
			"number list",
			args{
				tokens: []lexer.Token{
					&lexer.ListStartToken{},
					lexer.NewNumberToken([]byte{'1'}),
					&lexer.ListSeparatorToken{},
					lexer.NewNumberToken([]byte{'2'}),
					&lexer.ListSeparatorToken{},
					lexer.NewNumberToken([]byte{'3'}),
					&lexer.ListEndToken{},
				},
			},
			&ListNode{
				nodes: []Node{
					&NumberNode{token: lexer.NewNumberToken([]byte{'1'})},
					&NumberNode{token: lexer.NewNumberToken([]byte{'2'})},
					&NumberNode{token: lexer.NewNumberToken([]byte{'3'})},
				},
			},
			false,
		},
		{
			"string list",
			args{
				tokens: []lexer.Token{
					&lexer.ListStartToken{},
					&lexer.QuoteToken{},
					lexer.NewStringToken([]byte{'a'}),
					&lexer.QuoteToken{},
					&lexer.ListSeparatorToken{},
					&lexer.QuoteToken{},
					lexer.NewStringToken([]byte{'b'}),
					&lexer.QuoteToken{},
					&lexer.ListSeparatorToken{},
					&lexer.QuoteToken{},
					lexer.NewStringToken([]byte{'c'}),
					&lexer.QuoteToken{},
					&lexer.ListEndToken{},
				},
			},
			&ListNode{
				nodes: []Node{
					&StringNode{tokens: []lexer.Token{
						lexer.NewStringToken([]byte{'a'}),
					}},
					&StringNode{tokens: []lexer.Token{
						lexer.NewStringToken([]byte{'b'}),
					}},
					&StringNode{tokens: []lexer.Token{
						lexer.NewStringToken([]byte{'c'}),
					}},
				},
			},
			false,
		},
		{
			"symbol list",
			args{
				tokens: []lexer.Token{
					&lexer.ListStartToken{},
					lexer.NewStringToken([]byte{'a'}),
					&lexer.ListSeparatorToken{},
					lexer.NewStringToken([]byte{'b'}),
					&lexer.ListSeparatorToken{},
					lexer.NewStringToken([]byte{'c'}),
					&lexer.ListEndToken{},
				},
			},
			&ListNode{
				nodes: []Node{
					&SymbolNode{tokens: []lexer.Token{
						lexer.NewStringToken([]byte{'a'}),
					}},
					&SymbolNode{tokens: []lexer.Token{
						lexer.NewStringToken([]byte{'b'}),
					}},
					&SymbolNode{tokens: []lexer.Token{
						lexer.NewStringToken([]byte{'c'}),
					}},
				},
			},
			false,
		},
		{
			"complex list",
			args{
				tokens: []lexer.Token{
					&lexer.ListStartToken{},
					lexer.NewStringToken([]byte("Name")),
					&lexer.ListSeparatorToken{},
					&lexer.QuoteToken{},
					lexer.NewStringToken([]byte("John")),
					&lexer.QuoteToken{},
					&lexer.ListEndToken{},
				},
			},
			&ListNode{
				nodes: []Node{
					&SymbolNode{tokens: []lexer.Token{
						lexer.NewStringToken([]byte("Name")),
					}},
					&StringNode{tokens: []lexer.Token{
						lexer.NewStringToken([]byte("John")),
					}},
				},
			},
			false,
		},
		{
			"nested list",
			args{
				tokens: []lexer.Token{
					&lexer.ListStartToken{},
					&lexer.ListStartToken{},
					lexer.NewNumberToken([]byte{'1'}),
					&lexer.ListEndToken{},
					&lexer.ListSeparatorToken{},
					&lexer.ListStartToken{},
					lexer.NewNumberToken([]byte{'2'}),
					&lexer.ListEndToken{},
					&lexer.ListSeparatorToken{},
					&lexer.ListStartToken{},
					lexer.NewNumberToken([]byte{'3'}),
					&lexer.ListEndToken{},
					&lexer.ListEndToken{},
				},
			},
			&ListNode{
				nodes: []Node{
					&ListNode{
						nodes: []Node{
							&NumberNode{
								token: lexer.NewNumberToken([]byte{'1'}),
							},
						},
					},
					&ListNode{
						nodes: []Node{
							&NumberNode{
								token: lexer.NewNumberToken([]byte{'2'}),
							},
						},
					},
					&ListNode{
						nodes: []Node{
							&NumberNode{
								token: lexer.NewNumberToken([]byte{'3'}),
							},
						},
					},
				},
			},
			false,
		},
		{
			`sample/((Name "John") (Age 18))`,
			args{
				tokens: []lexer.Token{
					&lexer.ListStartToken{},
					&lexer.ListStartToken{},
					lexer.NewStringToken([]byte("Name")),
					&lexer.ListSeparatorToken{},
					&lexer.QuoteToken{},
					lexer.NewStringToken([]byte("John")),
					&lexer.QuoteToken{},
					&lexer.ListEndToken{},
					&lexer.ListSeparatorToken{},
					&lexer.ListStartToken{},
					lexer.NewStringToken([]byte("Age")),
					&lexer.ListSeparatorToken{},
					lexer.NewNumberToken([]byte("18")),
					&lexer.ListEndToken{},
					&lexer.ListEndToken{},
				},
			},
			&ListNode{
				nodes: []Node{
					&ListNode{
						nodes: []Node{
							&SymbolNode{
								tokens: []lexer.Token{lexer.NewStringToken([]byte("Name"))},
							},
							&StringNode{
								tokens: []lexer.Token{lexer.NewStringToken([]byte("John"))},
							},
						},
					},
					&ListNode{
						nodes: []Node{
							&SymbolNode{
								tokens: []lexer.Token{lexer.NewStringToken([]byte("Age"))},
							},
							&NumberNode{
								token: lexer.NewNumberToken([]byte("18")),
							},
						},
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				current: &initialState{},
			}
			got, err := p.parse(tt.args.tokens)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
