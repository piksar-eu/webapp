package web

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var sessionStore SessionStore

func SessionMiddleware(store SessionStore) func(next http.Handler) http.Handler {

	sessionStore = store

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "OPTIONS" {
				next.ServeHTTP(w, r)
				return
			}

			var session *Session

			cookie, err := r.Cookie("session_id")
			if err == nil {
				session = getSession(cookie.Value)
			}

			if session == nil {
				session := createSession()
				store.Save(session)

				http.SetCookie(w, &http.Cookie{
					Name:    "session_id",
					Value:   session.Id,
					Expires: session.ExpiresAt,
					Path:    "/",
				})
			}

			ctx := context.WithValue(r.Context(), SessionKey, &SessionContext{session: session})
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func getSession(sessionID string) *Session {
	session, err := sessionStore.Get(sessionID)

	if err == nil && session.ExpiresAt.After(time.Now()) {
		return session
	}

	return nil
}

func createSession() *Session {
	createdAt := time.Now()

	return &Session{
		Id:        uuid.New().String(),
		CreatedAt: createdAt,
		ExpiresAt: createdAt.Add(24 * time.Hour),
	}
}

type Session struct {
	Id        string
	CreatedAt time.Time
	ExpiresAt time.Time
	Data      map[string]interface{}
}

type SessionStore interface {
	Get(string) (*Session, error)
	Save(*Session) error
}

type contextKey string

const SessionKey contextKey = "session"

type SessionContext struct {
	session *Session
}

func (s *SessionContext) Id() string {
	return s.session.Id
}

func (s *SessionContext) Add(key string, val interface{}) {
	if s.session.Data == nil {
		s.session.Data = make(map[string]interface{})
	}

	s.session.Data[key] = val

	sessionStore.Save(s.session)
}

func (s *SessionContext) Get(key string) interface{} {
	val, ok := s.session.Data[key]

	if !ok {
		return nil
	}

	return val
}

func (s *SessionContext) Del(key string) {
	_, ok := s.session.Data[key]

	if !ok {
		return
	}

	delete(s.session.Data, key)
}

func SessionCtx(r *http.Request) *SessionContext {
	ctx := r.Context()

	sessCtx, ok := ctx.Value(SessionKey).(*SessionContext)
	if !ok {
		return nil
	}

	return sessCtx
}
