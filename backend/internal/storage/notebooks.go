package storage

import (
	"database/sql"
	"time"
)

type NotebookStorage struct {
	Conn *sql.DB
}

func NewNotebookStorage(conn *sql.DB) *NotebookStorage {
	return &NotebookStorage{Conn: conn}
}

type NewNotebook struct {
	UserId int `json:"user_id"`
	Name string `json:"name"`
	CreatedAt time.Time
}

type Notebook struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *NotebookStorage) CreateNotebook(notebook *NewNotebook) (int, error) {
	var notebookId int

	notebook.CreatedAt = time.Now()

	if err := s.Conn.QueryRow(
		"INSERT INTO notebooks (user_id, name, created_at) VALUES ($1, $2, $3) RETURNING id",
		notebook.UserId, notebook.Name, notebook.CreatedAt,
	).Scan(&notebookId); err != nil {
		return 0, err
	}

	return notebookId, nil
}

func (s *NotebookStorage) GetNotebooks(userId int) ([]Notebook, error) {
	rows, err := s.Conn.Query(
		"SELECT id, user_id, name, created_at FROM notebooks WHERE user_id = $1",
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notebooks := make([]Notebook, 0)

	for rows.Next() {
		var notebook Notebook
		if err := rows.Scan(
			&notebook.Id,
			&notebook.UserId,
			&notebook.Name,
			&notebook.CreatedAt,
		); err != nil {
			return nil, err
		}
		notebooks = append(notebooks, notebook)
	}

	return notebooks, nil
}

func (s *NotebookStorage) GetNotebookById(notebookId int) (Notebook, error) {
	notebook := Notebook{}

	if err := s.Conn.QueryRow(
		"SELECT id, user_id, name, created_at FROM notebooks WHERE id = $1",
		notebookId,
	).Scan(
		&notebook.Id,
		&notebook.UserId,
		&notebook.Name,
		&notebook.CreatedAt,
	); err != nil {
		return Notebook{}, err
	}

	return notebook, nil
}