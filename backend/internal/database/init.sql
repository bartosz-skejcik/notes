-- Create the 'users' table if it doesn't exist
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR PRIMARY KEY,
    email VARCHAR UNIQUE NOT NULL,
    password_hash VARCHAR NOT NULL,
    display_name VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the 'entries' table if it doesn't exist
CREATE TABLE IF NOT EXISTS entries (
    id VARCHAR PRIMARY KEY,
    user_id VARCHAR REFERENCES users(id),
    title VARCHAR NOT NULL,
    content TEXT,  -- Storing markdown as text
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    parent_entry_id VARCHAR REFERENCES entries(id) ON DELETE SET NULL,
    has_photo BOOLEAN DEFAULT FALSE,
    type VARCHAR NOT NULL
);

-- Create the 'photos' table if it doesn't exist
CREATE TABLE IF NOT EXISTS photos (
    id VARCHAR PRIMARY KEY,
    entry_id VARCHAR REFERENCES entries(id) ON DELETE CASCADE,
    image_data BYTEA NOT NULL,
    mime_type VARCHAR NOT NULL,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the 'tags' table if it doesn't exist
CREATE TABLE IF NOT EXISTS tags (
    id VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL,
    user_id VARCHAR REFERENCES users(id)
);

-- Create the 'entry_tags' table if it doesn't exist
CREATE TABLE IF NOT EXISTS entry_tags (
    id VARCHAR PRIMARY KEY,
    entry_id VARCHAR REFERENCES entries(id) ON DELETE CASCADE,
    tag_id VARCHAR REFERENCES tags(id) ON DELETE CASCADE
);

-- Create the 'settings' table if it doesn't exist
CREATE TABLE IF NOT EXISTS settings (
    id VARCHAR PRIMARY KEY,
    user_id VARCHAR REFERENCES users(id),
    key VARCHAR NOT NULL,
    value VARCHAR NOT NULL
);

-- Indexes for faster querying

-- Index for users.email (already unique)
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- Index for entries.user_id
CREATE INDEX IF NOT EXISTS idx_entries_user_id ON entries(user_id);

-- Index for entries.timestamp
CREATE INDEX IF NOT EXISTS idx_entries_timestamp ON entries(timestamp);

-- Index for entries.parent_entry_id
CREATE INDEX IF NOT EXISTS idx_entries_parent_entry_id ON entries(parent_entry_id);

-- Index for photos.entry_id
CREATE INDEX IF NOT EXISTS idx_photos_entry_id ON photos(entry_id);

-- Index for tags.user_id
CREATE INDEX IF NOT EXISTS idx_tags_user_id ON tags(user_id);

-- Index for entry_tags.entry_id
CREATE INDEX IF NOT EXISTS idx_entry_tags_entry_id ON entry_tags(entry_id);

-- Index for entry_tags.tag_id
CREATE INDEX IF NOT EXISTS idx_entry_tags_tag_id ON entry_tags(tag_id);

-- Index for settings.user_id
CREATE INDEX IF NOT EXISTS idx_settings_user_id ON settings(user_id);

-- Index for settings.key to optimize key-value lookups
CREATE INDEX IF NOT EXISTS idx_settings_key ON settings(key);
