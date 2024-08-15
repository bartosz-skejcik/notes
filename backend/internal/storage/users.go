package storage

import (
	"database/sql"
	"fmt"
)

type UserStorage struct {
	Conn *sql.DB
}

func NewUserStorage(conn *sql.DB) *UserStorage {
	return &UserStorage{Conn: conn}
}

type NewUser struct {
	DisplayName string `json:"display_name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (s *UserStorage) CreateUser(user *NewUser) (int, error) {
    var userId int

    if err := s.Conn.QueryRow(
        "INSERT INTO users (display_name, email, password_hash) VALUES ($1, $2, $3) RETURNING id",
        user.DisplayName, user.Email, user.Password,
    ).Scan(&userId); err != nil {
        return 0, fmt.Errorf("failed to create user: %v", err)
    }

    return userId, nil
}

