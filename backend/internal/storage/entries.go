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
	NotebookID int
	Title      string
	Content    string
	Role       *string
	HasPhoto   bool
	TagId *int
}

type NewChildEntry struct {
	NotebookID int
	Title      string
	Content    string
	Role       *string
	ParentEntryId *int
	TagId *int
}

type NewPhoto struct {
	EntryId int
	AuthorId int
	ImageData []byte
	MimeType string
	UploadedAt time.Time
}

type Entry struct {
	ID        int `json:"id"`
	NotebookID int `json:"notebook_id"`
	AuthorId int `json:"author_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Role      string `json:"role"`
	Timestamp time.Time `json:"timestamp"`
	ParentEntryId *int `json:"parent_entry_id"`
	HasPhoto bool `json:"has_photo"`
	TagId *int `json:"tag_id"`
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
	Role      string `json:"role"`
	Timestamp time.Time `json:"timestamp"`
	ParentEntryId *int `json:"parent_entry_id"`
	TagId *int `json:"tag_id"`
}

func (s *EntryStorage) CreateEntry(userId int, entry *NewEntry) (*Entry, error) {
    var entryId int
    var newEntry Entry

    // Prepare the SQL query for insertion
    var insertQuery string
    var insertArgs []interface{}

    if *entry.TagId == 0 {
        insertQuery = `INSERT INTO entries (notebook_id, author_id, title, content, role)
                        VALUES ($1, $2, $3, $4, $5) RETURNING id`
        insertArgs = []interface{}{entry.NotebookID, userId, entry.Title, entry.Content, entry.Role}
    } else {
        insertQuery = `INSERT INTO entries (notebook_id, author_id, title, content, tag_id, role)
                        VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
        insertArgs = []interface{}{entry.NotebookID, userId, entry.Title, entry.Content, *entry.TagId, entry.Role}
    }

    // Execute the insertion query and get the new entry ID
    err := s.Conn.QueryRow(insertQuery, insertArgs...).Scan(&entryId)
    if err != nil {
        return nil, err
    }

    // Retrieve the newly created entry from the database
    err = s.Conn.QueryRow(`SELECT id, notebook_id, author_id, title, content, timestamp, has_photo, tag_id, role
                           FROM entries
                           WHERE id = $1`, entryId).Scan(
        &newEntry.ID,
        &newEntry.NotebookID,
        &newEntry.AuthorId,
        &newEntry.Title,
        &newEntry.Content,
        &newEntry.Timestamp,
        &newEntry.HasPhoto,
        &newEntry.TagId,
		&newEntry.Role,
    )
    if err != nil {
        return nil, err
    }

    return &newEntry, nil
}


func (s *EntryStorage) CreateChildEntry(userId int, entry *NewChildEntry) (*ChildrenEntry, error) {
    var entryId int
    var newEntry ChildrenEntry


    // Prepare the SQL query for insertion
    var insertQuery string
    var insertArgs []interface{}

    if entry.Role == nil {
        insertQuery = `INSERT INTO entries (notebook_id, author_id, title, content, parent_entry_id, tag_id)
                        VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
        insertArgs = []interface{}{entry.NotebookID, userId, entry.Title, entry.Content, entry.ParentEntryId, *entry.TagId}
    } else {
        insertQuery = `INSERT INTO entries (notebook_id, author_id, title, content, tag_id, parent_entry_id, role)
                        VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
        insertArgs = []interface{}{entry.NotebookID, userId, entry.Title, entry.Content, *entry.TagId, *entry.ParentEntryId, *entry.Role}
    }

    // Execute the insertion query and get the new entry ID
    err := s.Conn.QueryRow(insertQuery, insertArgs...).Scan(&entryId)
    if err != nil {
        return nil, err
    }

    // Retrieve the newly created entry from the database
    err = s.Conn.QueryRow(`SELECT id, notebook_id, author_id, title, content, timestamp, tag_id, parent_entry_id, role
                           FROM entries
                           WHERE id = $1`, entryId).Scan(
        &newEntry.ID,
        &newEntry.NotebookID,
        &newEntry.AuthorId,
        &newEntry.Title,
        &newEntry.Content,
        &newEntry.Timestamp,
        &newEntry.TagId,
        &newEntry.ParentEntryId,
		&newEntry.Role,
    )
    if err != nil {
        return nil, err
    }

    return &newEntry, nil
}

func (s *EntryStorage) UpdateChildEntry(userId int, entryId int, childEntryId int, entry *NewChildEntry) (*ChildrenEntry, error) {
    var newEntry ChildrenEntry


    // Prepare the SQL query for insertion
    var insertQuery string
    var insertArgs []interface{}

	insertQuery = `UPDATE entries
					SET notebook_id = $1, author_id = $2, title = $3, content = $4, role = $5, parent_entry_id = $6, tag_id = $7
					WHERE id = $8
					RETURNING id`
	insertArgs = []interface{}{entry.NotebookID, userId, entry.Title, entry.Content, entry.Role, entry.ParentEntryId, entry.TagId, childEntryId}


    // Execute the insertion query and get the new entry ID
    err := s.Conn.QueryRow(insertQuery, insertArgs...).Scan(&entryId)
    if err != nil {
        return nil, err
    }

    // Retrieve the newly created entry from the database
    err = s.Conn.QueryRow(`SELECT id, notebook_id, author_id, title, content, timestamp, tag_id, parent_entry_id, role
                           FROM entries
                           WHERE id = $1`, entryId).Scan(
        &newEntry.ID,
        &newEntry.NotebookID,
        &newEntry.AuthorId,
        &newEntry.Title,
        &newEntry.Content,
        &newEntry.Timestamp,
        &newEntry.TagId,
        &newEntry.ParentEntryId,
		&newEntry.Role,
    )
    if err != nil {
        return nil, err
    }

    return &newEntry, nil
}

func (s *EntryStorage) UpdateEntry(userId int, entryId int, entry *NewEntry) (*Entry, error) {
    var newEntry Entry

    // Prepare the SQL query for insertion
    var insertQuery string
    var insertArgs []interface{}

	insertQuery = `UPDATE entries
					SET notebook_id = $1, author_id = $2, title = $3, content = $4, tag_id = $5, role = $6
					WHERE id = $7
					RETURNING id`
	insertArgs = []interface{}{entry.NotebookID, userId, entry.Title, entry.Content, *entry.TagId, entry.Role, entryId}


    // Execute the insertion query and get the new entry ID
    err := s.Conn.QueryRow(insertQuery, insertArgs...).Scan(&entryId)
    if err != nil {
        return nil, err
    }

    // Retrieve the newly created entry from the database
    err = s.Conn.QueryRow(`SELECT id, notebook_id, author_id, title, content, timestamp, tag_id, parent_entry_id, role
                           FROM entries
                           WHERE id = $1`, entryId).Scan(
        &newEntry.ID,
        &newEntry.NotebookID,
        &newEntry.AuthorId,
        &newEntry.Title,
        &newEntry.Content,
        &newEntry.Timestamp,
        &newEntry.TagId,
        &newEntry.ParentEntryId,
		&newEntry.Role,
    )
    if err != nil {
        return nil, err
    }

    return &newEntry, nil
}


func (s *EntryStorage) GetEntriesForNotebook(userId int, notebookId int) ([]Entry, error) {
	rows, err := s.Conn.Query(`SELECT id, notebook_id, author_id, title, content, timestamp, has_photo, tag_id, parent_entry_id, role
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
		err := rows.Scan(&entry.ID, &entry.NotebookID, &entry.AuthorId, &entry.Title, &entry.Content, &entry.Timestamp, &entry.HasPhoto, &entry.TagId, &entry.ParentEntryId, &entry.Role)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}


func (s *EntryStorage) GetEntryById(userId int, notebookId int, entryId int) (Entry, error) {
	row := s.Conn.QueryRow(`SELECT id, notebook_id, author_id, title, content, timestamp, has_photo, tag_id, role
		FROM entries
		WHERE id = $1
		AND notebook_id = $2
		AND author_id = $3`,
		entryId,
		userId,
		notebookId,
	)

	var entry Entry
	err := row.Scan(&entry.ID, &entry.NotebookID, &entry.AuthorId, &entry.Title, &entry.Content, &entry.Timestamp, &entry.HasPhoto, &entry.TagId, &entry.Role)
	if err != nil {
		return entry, err
	}

	return entry, nil
}

func (s *EntryStorage) GetEntryChildren(userId int, notebookId int, entryId int) ([]ChildrenEntry, error) {
	rows, err := s.Conn.Query(`SELECT id, notebook_id, author_id, title, content, timestamp, tag_id, parent_entry_id, role
		FROM entries
		WHERE parent_entry_id = $1
		AND notebook_id = $2
		AND author_id = $3
		ORDER BY timestamp ASC`,
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
		err := rows.Scan(&entry.ID, &entry.NotebookID, &entry.AuthorId, &entry.Title, &entry.Content, &entry.Timestamp, &entry.TagId, &entry.ParentEntryId, &entry.Role)
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