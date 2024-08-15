package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type SessionManager struct {
	Rdb *redis.Client
	Conn *sql.DB
}

func NewSessionManager(rdb *redis.Client, conn *sql.DB) *SessionManager {
	return &SessionManager{
		Rdb: rdb,
		Conn: conn,
	}
}

type UserSession struct {
	Id int `json:"id"`
	DisplayName string `json:"display_name"`
	Email string `json:"email"`
}

type User struct {
	Id int `db:"id"`
	DisplayName string `db:"display_name"`
	Email string `db:"email"`
	PasswordHash string `db:"password_hash"`
}

func (s *SessionManager) CreateSession(user *UserSession) (string, error) {
	sessionId := uuid.NewString()
	data, _ := json.Marshal(user)

	if err := s.Rdb.Set(context.Background(), sessionId, string(data), 24*time.Hour).Err(); err != nil {
		return "", err
	}
	return sessionId, nil
}

func (s *SessionManager) SignIn(email, password string) (*UserSession, error) {
	// Check if the user exists
	var user User
	if err := s.Conn.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user); err != nil {
		return nil, err
	}

	// Check if the password matches
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, err
	}

	// Create a new session for the user
	_, err := s.CreateSession(&UserSession{
		Id: user.Id,
		DisplayName: user.DisplayName,
		Email: user.Email,
	})

	if err != nil {
		return nil, err
	}

	return &UserSession{
		Id: user.Id,
		DisplayName: user.DisplayName,
		Email: user.Email,
	}, nil
}

func (s *SessionManager) SignOut(sessionId string) error {
	return s.Rdb.Del(context.Background(), sessionId).Err()
}

func (s *SessionManager) GetSession(sessionId string) (*UserSession, error) {
	data, err := s.Rdb.Get(context.Background(), sessionId).Result()
	if err != nil {
		return nil, err
	}

	var user UserSession
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return nil, err
	}

	return &user, nil
}