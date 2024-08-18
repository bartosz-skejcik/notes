-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR UNIQUE NOT NULL,
    password_hash VARCHAR NOT NULL,
    display_name VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the notebooks table
CREATE TABLE IF NOT EXISTS notebooks (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the tags table
CREATE TABLE IF NOT EXISTS tags (
    id SERIAL PRIMARY KEY,
    label VARCHAR NOT NULL,
    value VARCHAR NOT NULL,
    color VARCHAR NOT NULL
);

-- Create the entries table
CREATE TABLE IF NOT EXISTS entries (
    id SERIAL PRIMARY KEY,
    notebook_id INT REFERENCES notebooks(id) ON DELETE CASCADE,
    author_id INT REFERENCES users(id) ON DELETE SET NULL,
    tag_id INT REFERENCES tags(id) ON DELETE SET NULL DEFAULT 1,
    title VARCHAR NOT NULL,
    content TEXT DEFAULT '',
    role VARCHAR NOT NULL DEFAULT 'user',
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    parent_entry_id INT REFERENCES entries(id) ON DELETE CASCADE,
    has_photo BOOLEAN DEFAULT FALSE
);

-- Create the photos table
CREATE TABLE IF NOT EXISTS photos (
    id SERIAL PRIMARY KEY,
    entry_id INT REFERENCES entries(id) ON DELETE CASCADE,
    author_id INT REFERENCES users(id) ON DELETE SET NULL,
    image_data BYTEA NOT NULL,
    mime_type VARCHAR NOT NULL,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the settings table
CREATE TABLE IF NOT EXISTS settings (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    key VARCHAR NOT NULL,
    value VARCHAR NOT NULL
);

-- Add indexes for faster queries
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_notebooks_user_id ON notebooks(user_id);
CREATE INDEX IF NOT EXISTS idx_entries_notebook_id ON entries(notebook_id);
CREATE INDEX IF NOT EXISTS idx_entries_user_id ON entries(author_id);
CREATE INDEX IF NOT EXISTS idx_entries_tag_id ON entries(tag_id);
CREATE INDEX IF NOT EXISTS idx_photos_entry_id ON photos(entry_id);
CREATE INDEX IF NOT EXISTS idx_photos_author_id ON photos(author_id);
CREATE INDEX IF NOT EXISTS idx_settings_user_id ON settings(user_id);

-- Populate tags table with default values
INSERT INTO tags (label, value, color) VALUES ('None', 'none', '--muted');
INSERT INTO tags (label, value, color) VALUES ('Highlight', 'highlight', '--orange');
INSERT INTO tags (label, value, color) VALUES ('Do later', 'do-later', '--green');
INSERT INTO tags (label, value, color) VALUES ('New idea', 'new-idea', '--blue');
