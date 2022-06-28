package tests

import (
	"testing"

	"github.com/StackSistemas/generalgo"
)

func TestStringMatch(t *testing.T) {
	type args struct {
		value   string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "SucessMatch",
			args: args{"Matheus Saraiva", generalgo.PersonNameValidRegex},
			want: true,
		},
		{
			name: "FailMatch",
			args: args{"Matheus Saraiv4", generalgo.PersonNameValidRegex},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generalgo.StringMatch(tt.args.value, tt.args.pattern); got != tt.want {
				t.Errorf("StringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCnpjIsValid(t *testing.T) {
	type args struct {
		docnum string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "CnpjValid",
			args: args{"59541264000103"},
			want: true,
		},
		{
			name: "CnpjWithInvalidDigit",
			args: args{"59541264000144"},
			want: false,
		},
		{
			name: "CnpjWithInvalidLen",
			args: args{"595412640001"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generalgo.CnpjIsValid(tt.args.docnum); got != tt.want {
				t.Errorf("CnpjIsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
