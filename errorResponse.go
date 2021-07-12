package shippo

import (
	"fmt"
	"strings"
)

type errorResponse map[string][]string

func (e errorResponse) Error() string {
	var errors []string
	for key, group := range e {
		joined := strings.Join(group, ", ")
		err := fmt.Sprintf("%s error: %s", key, joined)
		errors = append(errors, err)
	}

	switch len(errors) {
	case 0:
		return ""
	case 1:
		return errors[0]

	default:
		return strings.Join(errors, "\n")
	}
}
