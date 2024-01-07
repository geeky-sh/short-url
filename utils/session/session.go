package session

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const SESSION_KEY = "session_id"

type sessionInfo struct {
	userID    uint
	username  string
	expiresAt time.Time
}

type Store struct {
	sessions map[string]sessionInfo
}

func Init() *Store {
	return &Store{sessions: make(map[string]sessionInfo)}
}

func (r *Store) Create(userID uint, username string) string {
	uuid := uuid.NewString()
	r.sessions[uuid] = sessionInfo{userID: userID, username: username, expiresAt: time.Now().Add(72 * time.Hour)}
	return uuid
}

func (r *Store) GetID(uid string) (uint, error) {
	si, ok := r.sessions[uid]
	if !ok {
		return 0, errors.New("session does not exist")
	}
	if si.expiresAt.Before(time.Now()) {
		return 0, errors.New("session has expired")
	}
	return si.userID, nil
}

func (r *Store) ListAll() {
	fmt.Println(len(r.sessions))
	for _, s := range r.sessions {
		fmt.Printf("%d %s\n", s.userID, s.username)
	}
}
