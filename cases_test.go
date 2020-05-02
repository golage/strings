package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToSnake(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "must returns snake case from camel case",
			arg:  "postgresHost",
			want: "postgres_host",
		},
		{
			name: "must returns snake case from capital case",
			arg:  "PostgresHost",
			want: "postgres_host",
		},
		{
			name: "must returns snake case from snake case",
			arg:  "postgres_host",
			want: "postgres_host",
		},
		{
			name: "must returns snake case from string with special chars and spaces",
			arg:  "Postgres \t@% \n &*/.~ Host",
			want: "postgres_host",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, ToSnake(tt.arg), tt.want)
		})
	}
}
