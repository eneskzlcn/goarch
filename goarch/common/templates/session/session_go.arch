package session

import (
	"errors"
	"net/http"
)

type Session interface {
	Exists(r *http.Request, key string) bool
	Get(r *http.Request, key string) any
	Enable(handler http.Handler) http.Handler
	GetString(r *http.Request, key string) string
	PopError(r *http.Request, key string) error
	Pop(r *http.Request, key string) any
	Put(r *http.Request, key string, data any)
	Remove(r *http.Request, key string)
}

func validateSessionKey(s string) error {
	if len(s) != 32 {
		return errors.New("length of session key must be in length 32")
	}
	return nil
}
