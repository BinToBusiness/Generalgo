package tests

import (
	"testing"

	"github.com/StackSistemas/generalgo"
)

func TestStringMatch(t *testing.T) {
	type args struct {
		value    string
		patterns []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "SucessMatch",
			args: args{"Matheus Saraiva", []string{generalgo.PersonNameValidRegex}},
			want: true,
		},
		{
			name: "SucessTwoMatch",
			args: args{
				"Matheus Saraiva",
				[]string{generalgo.PersonNameValidRegex, `^.{1,40}$`},
			},
			want: true,
		},
		{
			name: "FailWithOnePattern",
			args: args{"Matheus Saraiv4", []string{generalgo.PersonNameValidRegex}},
			want: false,
		},
		{
			name: "FailTwoPattern",
			args: args{
				"Matheus Saraiva da Silva Pereira Gusmão de Souza Cavalcante Texeira de Freitas",
				[]string{generalgo.PersonNameValidRegex, `^.{1,40}$`},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generalgo.StringMatch(tt.args.value, tt.args.patterns...); got != tt.want {
				t.Errorf("StringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCpfOrCnpjIsValid(t *testing.T) {
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
			name: "CpfValid",
			args: args{"74715151200"},
			want: true,
		},
		{
			name: "CnpjWithInvalidDigit",
			args: args{"59541264000144"},
			want: false,
		},
		{
			name: "CpfWithInvalidDigit",
			args: args{"7475151209"},
			want: false,
		},
		{
			name: "CnpjWithInvalidLen",
			args: args{"595412640001"},
			want: false,
		},
		{
			name: "CpfWithInvalidLen",
			args: args{"747151512"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generalgo.CpfOrCnpjIsValid(tt.args.docnum); got != tt.want {
				t.Errorf("CpfOrCnpjIsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
