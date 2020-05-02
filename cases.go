package strings

import (
	"regexp"
	"strings"
)

var (
	snakeMatchSpaces   = regexp.MustCompile("[^a-zA-Z0-9_]")
	snakeMatchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	snakeMatchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

// ToSnake returns s with all Unicode letters mapped to their snake case.
func ToSnake(s string) string {
	snake := snakeMatchSpaces.ReplaceAllString(s, "")
	snake = snakeMatchFirstCap.ReplaceAllString(snake, "${1}_${2}")
	snake = snakeMatchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
