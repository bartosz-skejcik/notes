package storage

import (
	"database/sql"
)

type CategoryStorage struct {
	Conn *sql.DB
}

func NewCategoryStorage(conn *sql.DB) *CategoryStorage {
	return &CategoryStorage{
		Conn: conn,
	}
}

type NewCategory struct {
	NotebookId int `json:"notebook_id"`
	UserId int `json:"user_id"`
	Name string `json:"name"`
	Color string `json:"color"`
}

type Category struct {
	Id int `json:"id"`
	NotebookId int `json:"notebookId"`
	Name string `json:"name"`
	Color string `json:"color"`
}

func (s *CategoryStorage) CreateCategory(category *NewCategory) (int, error) {
	var categoryId int

	if err := s.Conn.QueryRow("INSERT INTO categories (notebook_id, user_id, name, color) VALUES ($1, $2, $3, $4) RETURNING id", category.NotebookId, category.UserId, category.Name, category.Color).Scan(&categoryId); err != nil {
		return 0, err
	}

	return categoryId, nil
}

func (s *CategoryStorage) GetCategories(notebookId int, userId int) ([]Category, error) {
	rows, err := s.Conn.Query("SELECT id, notebook_id, name, color FROM categories WHERE notebook_id = $1 AND user_id = $2", notebookId, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]Category, 0)

	for rows.Next() {
		var category Category
		err = rows.Scan(&category.Id, &category.NotebookId, &category.Name, &category.Color)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, rows.Err()
}