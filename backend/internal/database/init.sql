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
    user_id INT REFERENCES users(id),
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the tags table
CREATE TABLE IF NOT EXISTS tags (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);

-- Create the entries table
CREATE TABLE IF NOT EXISTS entries (
    id SERIAL PRIMARY KEY,
    notebook_id INT REFERENCES notebooks(id),
    tag_id INT REFERENCES tags(id),
    title VARCHAR NOT NULL,
    content TEXT,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    parent_entry_id INT REFERENCES entries(id),
    has_photo BOOLEAN DEFAULT FALSE,
    type VARCHAR
);

-- Create the photos table
CREATE TABLE IF NOT EXISTS photos (
    id SERIAL PRIMARY KEY,
    entry_id INT REFERENCES entries(id),
    image_data BYTEA,
    mime_type VARCHAR NOT NULL,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the settings table
CREATE TABLE IF NOT EXISTS settings (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    key VARCHAR NOT NULL,
    value VARCHAR NOT NULL
);

-- Add indexes for faster queries
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_notebooks_user_id ON notebooks(user_id);
CREATE INDEX IF NOT EXISTS idx_entries_notebook_id ON entries(notebook_id);
CREATE INDEX IF NOT EXISTS idx_entries_tag_id ON entries(tag_id);
CREATE INDEX IF NOT EXISTS idx_photos_entry_id ON photos(entry_id);
CREATE INDEX IF NOT EXISTS idx_settings_user_id ON settings(user_id);