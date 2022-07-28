package discogs

import (
	"encoding/json"
	"fmt"
	"io"
)

type knownError struct {
	Message string
}

type unknownError struct {
	Error string
}

// decodeError decode an error
// we expect a message field for known errors
// if decode fails, decode as an unknown error
func decodeError(r io.Reader, code int) error {
	var resp knownError

	err := json.NewDecoder(r).Decode(&resp)
	if err != nil {
		return decodeUnknownError(r, code)
	}

	return fmt.Errorf("status: %d, message: %s", code, resp.Message)
}

// decodeUnknownError decode an unknown error
// we expect an error field
// if decode fails, just return a generic error with the status code
func decodeUnknownError(r io.Reader, code int) error {
	var resp unknownError

	err := json.NewDecoder(r).Decode(&resp)
	if err != nil {
		return fmt.Errorf("status: %d, unknown error", code)
	}

	return fmt.Errorf("status: %d, error: %s", code, resp.Error)
}
