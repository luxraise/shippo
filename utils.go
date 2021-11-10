package shippo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

var unknownBytes = []byte(`"UNKNOWN"`)

func getRequestBody(request interface{}) (r io.Reader, err error) {
	if request == nil {
		return
	}

	var bs []byte
	if bs, err = json.Marshal(request); err != nil {
		return
	}

	r = bytes.NewReader(bs)
	return
}

func handleResponse(r io.Reader, value interface{}) (err error) {
	if value == nil {
		return
	}

	if err = json.NewDecoder(r).Decode(value); err != nil {
		return fmt.Errorf("error encountered while attempting to decode response as JSON: %v", err)
	}

	return
}

func handleError(r io.Reader) (err error) {
	var value errorResponse
	if err = handleResponse(r, &value); err != nil {
		return
	}

	return value
}

func isUnknown(bs []byte) bool {
	return bytes.Equal(unknownBytes, bs)
}

// NameToken represents a name/token pair
type NameToken struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}
