//////////////////////////////////////////////////////////////////////
//
// Given is a SessionManager that stores session information in
// memory. The SessionManager itself is working, however, since we
// keep on adding new sessions to the manager our program will
// eventually run out of memory.
//
// Your task is to implement a session cleaner routine that runs
// concurrently in the background and cleans every session that
// hasn't been updated for more than 5 seconds (of course usually
// session times are much longer).
//
// Note that we expect the session to be removed anytime between 5 and
// 7 seconds after the last update. Also, note that you have to be
// very careful in order to prevent race conditions.
//

package main

import (
	"errors"
	"log"
	"sync"
	"time"
)

const cleanerStep = 6 * time.Second

// SessionManager keeps track of all sessions from creation, updating
// to destroying.
type SessionManager struct {
	sessions sync.Map
}

// Session stores the session's data
type Session struct {
	Data       map[string]interface{}
	Expiration time.Time
}

// NewSessionManager creates a new sessionManager
func NewSessionManager() *SessionManager {
	return &SessionManager{}
}

// CreateSession creates a new session and returns the sessionID
func (m *SessionManager) CreateSession() (string, error) {
	sessionID, err := MakeSessionID()
	if err != nil {
		return "", err
	}

	m.sessions.Store(sessionID, Session{
		Data: make(map[string]interface{}),
	})

	return sessionID, nil
}

// ErrSessionNotFound returned when sessionID not listed in
// SessionManager
var ErrSessionNotFound = errors.New("SessionID does not exists")

// GetSessionData returns data related to session if sessionID is
// found, errors otherwise
func (m *SessionManager) GetSessionData(sessionID string) (map[string]interface{}, error) {
	session, ok := m.sessions.Load(sessionID)

	if !ok {
		return nil, ErrSessionNotFound
	}
	return session.(Session).Data, nil
}

// UpdateSessionData overwrites the old session data with the new one
func (m *SessionManager) UpdateSessionData(sessionID string, data map[string]interface{}) error {
	_, ok := m.sessions.Load(sessionID)
	if !ok {
		return ErrSessionNotFound
	}

	// Hint: you should renew expiry of the session here
	m.sessions.Store(sessionID, Session{
		Data:       data,
		Expiration: time.Now().Add(5 * time.Second),
	})

	updt <- time.Now()

	return nil
}

var updt = make(chan time.Time)

func main() {
	// Create new sessionManager and new session
	m := NewSessionManager()
	sID, err := m.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created new session with ID", sID)

	// Session Garbage collector
	go m.CleanOldTimer(updt)

	// Update session data
	data := make(map[string]interface{})
	data["website"] = "longhoang.de"
	updt <- time.Now()

	err = m.UpdateSessionData(sID, data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Update session data, set website to longhoang.de")

	// Retrieve data from manager again
	updatedData, err := m.GetSessionData(sID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Get session data:", updatedData)
}

// CleanOldTimer cleans the older sessions
func (m *SessionManager) CleanOldTimer(updated <-chan time.Time) {
	<-updated
	// from now it starts a countdown of 5 seconds?
	timer := time.NewTimer(cleanerStep)
	for {
		select {
		// new update has been made before timer end
		case <-updated:
			timer.Stop()
		case <-timer.C:
			m.CleanOld()
		}

		timer = time.NewTimer(cleanerStep)
	}
}

// CleanOld checks the sessions and clean the older ones
func (m *SessionManager) CleanOld() {
	m.sessions.Range(func(key, value interface{}) bool {
		if exp := value.(Session).Expiration; time.Now().After(exp) {
			m.sessions.Delete(key)
		}
		return true
	})
}
