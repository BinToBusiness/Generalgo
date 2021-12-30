package tests

import (
	"testing"

	generalgo "github.com/Microwatts/generalgo"
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
