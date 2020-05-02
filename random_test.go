package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandom(t *testing.T) {
	type args struct {
		length  int
		letters string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must returns random string with length greater-than letters",
			args: args{
				length:  10,
				letters: "QW@#as34",
			},
		},
		{
			name: "must returns random string with length less-than letters",
			args: args{
				length:  5,
				letters: "QW@#as34",
			},
		},
		{
			name: "must returns random string with length equal letters",
			args: args{
				length:  8,
				letters: "QW@#as34",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rnd := Random(tt.args.length, tt.args.letters)
			assert.Equal(t, len(rnd), tt.args.length)
			for _, c := range rnd {
				assert.Contains(t, tt.args.letters, string(c))
			}
		})
	}
}
