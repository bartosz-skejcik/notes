package storage

import (
	"database/sql"
	"time"
)

type EntryStorage struct {
	Conn *sql.DB
}

func NewEntryStorage(conn *sql.DB) *EntryStorage {
	return &EntryStorage{Conn: conn}
}

type NewEntry struct {
	NotebookID int `json:"notebook_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Timestamp  time.Time
	HasPhoto   bool `json:"has_photo"`
	TagId int `json:"tag_id"`
}

type NewChildEntry struct {
	NotebookID int `json:"notebook_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	ParentEntryId int `json:"parent_entry_id"`
	Timestamp  time.Time
	TagId int `json:"tag_id"`
}

type NewPhoto struct {
	EntryId int `json:"entry_id"`
	AuthorId int `json:"author_id"`
	ImageData []byte `json:"image_data"`
	MimeType string `json:"mime_type"`
	UploadedAt time.Time
}

type Entry struct {
	ID        int `json:"id"`
	NotebookID int `json:"notebook_id"`
	AuthorId int `json:"author_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Timestamp time.Time `json:"timestamp"`
	HasPhoto bool `json:"has_photo"`
	TagId int `json:"tag_id"`
}

type Photo struct {
	ID int `json:"id"`
	EntryId int `json:"entry_id"`
	AuthorId int `json:"author_id"`
	ImageData []byte `json:"image_data"`
	MimeType string `json:"mime_type"`
	UploadedAt time.Time `json:"uploaded_at"`
}

type ChildrenEntry struct {
	ID        int `json:"id"`
	NotebookID int `json:"notebook_id"`
	AuthorId int `json:"author_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Timestamp time.Time `json:"timestamp"`
	TagId int `json:"tag_id"`
}

func (s *EntryStorage) CreateEntry(userId int, entry *NewEntry) (int, error) {
	var entryId int

	entry.Timestamp = time.Now()

	err := s.Conn.QueryRow(`INSERT INTO entries (notebook_id, author_id, title, content, timestamp, tag_id)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		entry.NotebookID, userId, entry.Title, entry.Content, entry.Timestamp, entry.TagId).Scan(&entryId)

	if err != nil {
		return 0, err
	}

	return entryId, nil
}

func (s *EntryStorage) CreateChildEntry(userId int, entry *NewChildEntry) (int, error) {
	var entryId int

	entry.Timestamp = time.Now()

	err := s.Conn.QueryRow(`INSERT INTO entries (notebook_id, author_id, title, content, timestamp, tag_id, parent_entry_id)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		entry.NotebookID, userId, entry.Title, entry.Content, entry.Timestamp, entry.TagId, entry.ParentEntryId).Scan(&entryId)

	if err != nil {
		return 0, err
	}
	return entryId, nil
}

func (s *EntryStorage) GetEntriesForNotebook(userId int, notebookId int) ([]Entry, error) {
	rows, err := s.Conn.Query(`SELECT id, notebook_id, author_id, title, content, timestamp, has_photo, tag_id
		FROM entries
		WHERE notebook_id = $1
		AND author_id = $2
		ORDER BY timestamp DESC`,
		notebookId,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []Entry
	for rows.Next() {
		var entry Entry
		err := rows.Scan(&entry.ID, &entry.NotebookID, &entry.AuthorId, &entry.Title, &entry.Content, &entry.Timestamp, &entry.HasPhoto, &entry.TagId)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}


func (s *EntryStorage) GetEntryById(userId int, notebookId int, entryId int) (Entry, error) {
	row := s.Conn.QueryRow(`SELECT id, notebook_id, author_id, title, content, timestamp, has_photo, tag_id
		FROM entries
		WHERE id = $1
		AND notebook_id = $2
		AND author_id = $3`,
		entryId,
		userId,
		notebookId,
	)

	var entry Entry
	err := row.Scan(&entry.ID, &entry.NotebookID, &entry.AuthorId, &entry.Title, &entry.Content, &entry.Timestamp, &entry.HasPhoto, &entry.TagId)
	if err != nil {
		return entry, err
	}

	return entry, nil
}

func (s *EntryStorage) GetEntryChildren(userId int, notebookId int, entryId int) ([]ChildrenEntry, error) {
	rows, err := s.Conn.Query(`SELECT id, notebook_id, author_id, title, content, timestamp, tag_id
		FROM entries
		WHERE parent_entry_id = $1
		AND notebook_id = $2
		AND author_id = $3
		ORDER BY timestamp DESC`,
		entryId,
		notebookId,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []ChildrenEntry
	for rows.Next() {
		var entry ChildrenEntry
		err := rows.Scan(&entry.ID, &entry.NotebookID, &entry.AuthorId, &entry.Title, &entry.Content, &entry.Timestamp, &entry.TagId)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

func (s *EntryStorage) GetEntryPhotos(userId int, entryId int) ([]Photo, error) {
	rows, err := s.Conn.Query(`SELECT id, entry_id, author_id, image_data, mime_type, uploaded_at
		FROM photos
		WHERE entry_id = $1
		AND author_id = $2
		ORDER BY uploaded_at DESC`,
		entryId,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []Photo
	for rows.Next() {
		var photo Photo
		err := rows.Scan(&photo.ID, &photo.EntryId, &photo.AuthorId, &photo.ImageData, &photo.MimeType, &photo.UploadedAt)
		if err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}

	return photos, nil
}

func (s *EntryStorage) CreatePhoto(entryId int, photo *NewPhoto) (int, error) {
	var photoId int

	photo.UploadedAt = time.Now()

	err := s.Conn.QueryRow(`INSERT INTO photos (entry_id, author_id, image_data, mime_type, uploaded_at)
		VALUES ($1, $2, $3, $4, $5)`,
		entryId, photo.AuthorId, photo.ImageData, photo.MimeType, photo.UploadedAt).Scan(&photoId)

	if err != nil {
		return 0, err
	}

	return photoId, nil
}