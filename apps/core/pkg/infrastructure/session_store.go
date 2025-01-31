package infrastructure

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/piksar-eu/webapp/apps/core/pkg/web"
)

func NewInMemorySessionStore() web.SessionStore {
	return &inMemorySessionStore{
		data: make(map[string]web.Session),
	}
}

type inMemorySessionStore struct {
	data map[string]web.Session
	mu   sync.RWMutex
}

func (s *inMemorySessionStore) Get(sessionID string) (*web.Session, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	session, exists := s.data[sessionID]
	if !exists {
		return nil, errors.New("session not found")
	}

	return &session, nil
}

func (s *inMemorySessionStore) Save(session *web.Session) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[session.Id] = *session

	return nil
}

func NewPgSessionStore(db *sql.DB) web.SessionStore {
	return &pgSessionStore{
		db: db,
	}
}

type pgSessionStore struct {
	db *sql.DB
}

func (s *pgSessionStore) Get(sessionID string) (*web.Session, error) {
	rows, err := s.db.Query("SELECT created_at, expires_at, data FROM core__sessions WHERE id = $1 LIMIT 1", sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("session not found")
	}

	var createdAt time.Time
	var expiresAt time.Time
	var dataJSON string

	err = rows.Scan(&createdAt, &expiresAt, &dataJSON)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(dataJSON), &data)
	if err != nil {
		return nil, errors.New("can not unmarshal session data")
	}

	return &web.Session{
		Id:        sessionID,
		CreatedAt: createdAt,
		ExpiresAt: expiresAt,
		Data:      data,
	}, nil
}

func (s *pgSessionStore) Save(session *web.Session) error {
	dataJSON, err := json.Marshal(session.Data)
	if err != nil {
		return fmt.Errorf("session data can not be marschal")
	}

	query := `
		INSERT INTO core__sessions (id, created_at, expires_at, data)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (id) DO UPDATE
		SET created_at = EXCLUDED.created_at,
			expires_at = EXCLUDED.expires_at,
			data = EXCLUDED.data;
	`
	_, err = s.db.Exec(query, session.Id, session.CreatedAt, session.ExpiresAt, dataJSON)
	if err != nil {
		return fmt.Errorf("can not save session")
	}

	return nil
}

func NewCachedSessionStore(s web.SessionStore) web.SessionStore {
	cacheSize := 1000
	return &cachedSessionStore{
		sessionStore: s,
		data:         make(map[string]web.Session, cacheSize),
		order:        make([]string, 0, cacheSize),
		cacheSize:    cacheSize,
	}
}

type cachedSessionStore struct {
	sessionStore web.SessionStore
	data         map[string]web.Session
	order        []string
	cacheSize    int
	mu           sync.RWMutex
}

func (s *cachedSessionStore) Get(sessionID string) (*web.Session, error) {
	s.mu.RLock()
	if session, exists := s.data[sessionID]; exists {
		s.mu.RUnlock()
		return &session, nil
	}
	s.mu.RUnlock()

	session, err := s.sessionStore.Get(sessionID)
	if err != nil {
		return nil, err
	}

	s.addToCache(*session)
	return session, nil
}

func (s *cachedSessionStore) Save(session *web.Session) error {
	if err := s.sessionStore.Save(session); err != nil {
		return err
	}
	s.addToCache(*session)

	return nil
}

func (s *cachedSessionStore) addToCache(session web.Session) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[session.Id]; exists {
		s.data[session.Id] = session

		// Move the session ID to the end to mark it as recently used
		for i, id := range s.order {
			if id == session.Id {
				s.order = append(s.order[:i], s.order[i+1:]...)
				break
			}
		}
		s.order = append(s.order, session.Id)
		return
	}

	// If cache is full, remove the oldest session
	if len(s.data) >= s.cacheSize {
		oldestID := s.order[0]
		s.order = s.order[1:]
		delete(s.data, oldestID)
	}

	s.data[session.Id] = session
	s.order = append(s.order, session.Id)
}
