package session

import (
	"errors"
	"github.com/eneskzlcn/softdare/internal/config"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/golangcollege/sessions"
	"net/http"
)

type CollegeSessionAdapter struct {
	logger     logger.Logger
	session    *sessions.Session
	sessionKey []byte
}

func NewCollegeSessionAdapter(logger logger.Logger, config config.Session) *CollegeSessionAdapter {
	if logger == nil {
		return nil
	}
	if err := validateSessionKey(config.Key); err != nil {
		logger.Errorf("Err validating session key", err.Error())
		return nil
	}
	sessionKeyByte := []byte(config.Key)
	session := sessions.New(sessionKeyByte)
	return &CollegeSessionAdapter{session: session, logger: logger}
}

func (s *CollegeSessionAdapter) Exists(r *http.Request, key string) bool {
	exists := s.session.Exists(r, key)
	return exists
}

func (s *CollegeSessionAdapter) Get(r *http.Request, key string) any {
	return s.session.Get(r, key)
}

func (s *CollegeSessionAdapter) Put(r *http.Request, key string, data interface{}) {
	s.session.Put(r, key, data)
}

func (s *CollegeSessionAdapter) Remove(r *http.Request, key string) {
	s.session.Remove(r, key)
}

func (s *CollegeSessionAdapter) GetString(r *http.Request, key string) string {
	return s.session.GetString(r, key)
}

func (s *CollegeSessionAdapter) PopString(r *http.Request, key string) string {
	return s.session.PopString(r, key)
}

func (s *CollegeSessionAdapter) Pop(r *http.Request, key string) any {
	return s.session.Pop(r, key)
}

/*PopError extracts a string that known as oops from session and converts it to oops*/
func (s *CollegeSessionAdapter) PopError(r *http.Request, key string) error {
	str := s.PopString(r, key)
	if str != "" {
		return errors.New(str)
	}
	return nil
}
func (s *CollegeSessionAdapter) Enable(handler http.Handler) http.Handler {
	s.logger.Debug("session enabled for given handler")
	return s.session.Enable(handler)
}
