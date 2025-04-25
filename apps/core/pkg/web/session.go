package web

import (
	"context"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/piksar-eu/webapp/apps/core/pkg/shared"
)

var sessionStoreOnce sync.Once
var sessionStore SessionStore

func SessionMiddleware(store SessionStore) func(next http.Handler) http.Handler {

	sessionStoreOnce.Do(func() {
		sessionStore = store
	})

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
				session = createSession()
				store.Save(session)

				http.SetCookie(w, &http.Cookie{
					Name:    "session_id",
					Value:   session.Id,
					Expires: session.ExpiresAt,
					Path:    "/",
					Domain:  os.Getenv("SESSION_COOKIE_DOMAIN"),
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
		ExpiresAt: createdAt.Add(60 * 24 * time.Hour),
	}
}

type Session struct {
	Id        string
	CreatedAt time.Time
	ExpiresAt time.Time
	Data      *sync.Map
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
		s.session.Data = &sync.Map{}
	}
	s.session.Data.Store(key, val)
	sessionStore.Save(s.session)
}

func (s *SessionContext) Get(key string) interface{} {
	if s.session.Data == nil {
		return nil
	}
	val, ok := s.session.Data.Load(key)
	if !ok {
		return nil
	}
	return val
}

func (s *SessionContext) Del(keys ...string) {
	if s.session.Data == nil {
		return
	}

	keyExists := false
	for _, key := range keys {
		if _, ok := s.session.Data.Load(key); ok {
			s.session.Data.Delete(key)
			keyExists = true
		}
	}

	if keyExists {
		sessionStore.Save(s.session)
	}
}

func SessionCtx(r *http.Request) *SessionContext {
	ctx := r.Context()

	sessCtx, ok := ctx.Value(SessionKey).(*SessionContext)
	if !ok {
		return nil
	}

	return sessCtx
}

func GetSessionUser(r *http.Request) *shared.SessionUser {
	sessCtx := SessionCtx(r)

	if sessCtx == nil {
		return nil
	}

	user := sessCtx.Get("user")

	sessionUser := &shared.SessionUser{}
	if err := shared.MapToStruct(user, &sessionUser); err != nil {
		return nil
	}

	return sessionUser
}
