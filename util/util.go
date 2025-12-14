package util

import "strings"

type RequestError struct {
	Error string `json:"error"`
}

func NewRequestError(err error) RequestError {
	rerr := RequestError{Error: "Unknown error"}
	if err != nil {
		rerr.Error = err.Error()
	}
	return rerr
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
