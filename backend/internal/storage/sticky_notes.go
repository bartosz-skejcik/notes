package storage

import (
	"database/sql"
)

type StickyNoteStorage struct {
	Conn *sql.DB
}

func NewStickyNoteStorage(conn *sql.DB) *StickyNoteStorage {
	return &StickyNoteStorage{
		Conn: conn,
	}
}

type NewStickyNote struct {
	NotebookID int `json:"notebook_id"`
	CategoryID int `json:"category_id"`
	AuthorID int `json:"author_id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

type StickyNote struct {
	ID int `json:"id"`
	NotebookID int `json:"notebook_id"`
	CategoryID int `json:"category_id"`
	AuthorID int `json:"author_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func (s *StickyNoteStorage) CreateStickyNote(newStickyNote *NewStickyNote) (int, error) {
	var stickyNoteId int

	err := s.Conn.QueryRow(`INSERT INTO sticky_notes (notebook_id, category_id, author_id, title, content) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
	newStickyNote.NotebookID,
	newStickyNote.CategoryID,
	newStickyNote.AuthorID,
	newStickyNote.Title,
	newStickyNote.Content).Scan(&stickyNoteId)

	if err != nil {
		return 0, err
	}

	return stickyNoteId, nil
}

func (s *StickyNoteStorage) GetAllStickyNotesForUser(userId int) ([]StickyNote, error) {
	rows, err := s.Conn.Query(`SELECT id, notebook_id, category_id, author_id, title, content, created_at FROM sticky_notes WHERE author_id = $1`, userId)

	if err != nil { // check for errors
		return nil, err // return the error
	}
	defer rows.Close() // close the rows

	stickyNotes := make([]StickyNote, 0) // create a slice of type StickyNote

	for rows.Next() { // iterate over rows
		var stickyNote StickyNote // declare a variable of type StickyNote
		if err := rows.Scan(
			&stickyNote.ID,
			&stickyNote.NotebookID,
			&stickyNote.CategoryID,
			&stickyNote.AuthorID,
			&stickyNote.Title,
			&stickyNote.Content,
			&stickyNote.CreatedAt,
		); err != nil { // check for errors
			return nil, err // return the error
		}
		stickyNotes = append(stickyNotes, stickyNote) // append the stickyNote to the slice

	} // end of for loop

	return stickyNotes, nil // return the slice of stickyNotes
}

func (s *StickyNoteStorage) GetStickyNotesForUserCategory(userId int, categoryId int) ([]StickyNote, error) {
	rows, err := s.Conn.Query(`SELECT id, notebook_id, category_id, author_id, title, content, created_at FROM sticky_notes WHERE author_id = $1 AND category_id = $2`, userId, categoryId)

	if err != nil { // check for errors
		return nil, err // return the error
	}
	defer rows.Close() // close the rows

	stickyNotes := make([]StickyNote, 0) // create a slice of type StickyNote

	for rows.Next() { // iterate over rows
		var stickyNote StickyNote // declare a variable of type StickyNote
		if err := rows.Scan(
			&stickyNote.ID,
			&stickyNote.NotebookID,
			&stickyNote.CategoryID,
			&stickyNote.AuthorID,
			&stickyNote.Title,
			&stickyNote.Content,
			&stickyNote.CreatedAt,
		); err != nil { // check for errors
			return nil, err // return the error
		}
		stickyNotes = append(stickyNotes, stickyNote) // append the stickyNote to the slice

	} // end of for loop

	return stickyNotes, nil // return the slice of stickyNotes
}