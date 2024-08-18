package storage

import (
	"database/sql"
)

type TagStorage struct {
	Conn *sql.DB
}

func NewTagStorage(conn *sql.DB) *TagStorage {
	return &TagStorage{Conn: conn}
}

type Tag struct {
	ID int `json:"id"`
	Label string `json:"label"`
	Color string `json:"color"`
	Value string `json:"value"`
}

func (s *TagStorage) GetAllTags() ([]Tag, error) {
	rows, err := s.Conn.Query(`SELECT id, label, color, value
		FROM tags`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		var tag Tag
		err := rows.Scan(&tag.ID, &tag.Label, &tag.Color, &tag.Value)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}