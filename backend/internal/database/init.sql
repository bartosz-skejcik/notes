-- Create the users table
CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    display_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create the entries table
CREATE TABLE entries (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date DATE NOT NULL,
    type VARCHAR(50) NOT NULL
);

-- Create the entry_relations table
CREATE TABLE entry_relations (
    id UUID PRIMARY KEY,
    parent_entry_id UUID REFERENCES entries(id),
    child_entry_id UUID REFERENCES entries(id)
);

-- Create the tags table
CREATE TABLE tags (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    user_id UUID REFERENCES users(id)
);

-- Create the entry_tags table
CREATE TABLE entry_tags (
    id UUID PRIMARY KEY,
    entry_id UUID REFERENCES entries(id),
    tag_id UUID REFERENCES tags(id)
);

-- Add indexes for faster querying
CREATE INDEX idx_users_email ON users(email);

CREATE INDEX idx_entries_user_id ON entries(user_id);
CREATE INDEX idx_entries_date ON entries(date);
CREATE INDEX idx_entries_type ON entries(type);

CREATE INDEX idx_entry_relations_parent_entry_id ON entry_relations(parent_entry_id);
CREATE INDEX idx_entry_relations_child_entry_id ON entry_relations(child_entry_id);

CREATE INDEX idx_tags_user_id ON tags(user_id);
CREATE INDEX idx_tags_name ON tags(name);

CREATE INDEX idx_entry_tags_entry_id ON entry_tags(entry_id);
CREATE INDEX idx_entry_tags_tag_id ON entry_tags(tag_id);