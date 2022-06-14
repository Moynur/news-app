CREATE TABLE IF NOT EXISTS news_articles
(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    link TEXT NOT NULL UNIQUE,
    thumbnail TEXT NOT NULL,
    category TEXT,
    created_at timestamp DEFAULT current_timestamp
)