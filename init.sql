CREATE TABLE IF NOT EXISTS news_articles
(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    Link TEXT NOT NULL UNIQUE,
    Thumbnail TEXT NOT NULL,
    created_at timestamp DEFAULT current_timestamp
)