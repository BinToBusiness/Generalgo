package tests

import (
	"testing"

	"github.com/Microwatts/general"
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
			args: args{"Matheus Saraiva", general.PersonNameValidRegex},
			want: true,
		},
		{
			name: "FailMatch",
			args: args{"Matheus Saraiv4", general.PersonNameValidRegex},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := general.StringMatch(tt.args.value, tt.args.pattern); got != tt.want {
				t.Errorf("StringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
