package shippo

import (
	"encoding/json"
	"fmt"
	"strings"
)

type errorResponse map[string]errorsList

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

type errorsList []string

func (e *errorsList) UnmarshalJSON(bs []byte) (err error) {
	var ss []string
	if err = json.Unmarshal(bs, &ss); err == nil {
		*e = ss
		return
	}

	var str string
	if err = json.Unmarshal(bs, &str); err != nil {
		return
	}

	*e = errorsList{str}
	return
}
