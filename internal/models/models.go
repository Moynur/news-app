package models

type GetArticlesRequest struct {
	Cursor   int
	Category string
	Provider string
	Title    string
}

type GetArticlesResponse struct {
	NextCursor int
	Articles   []Article
}

type Article struct {
	ID       int
	Title    string
	Summary  string
	ImageRef string
	Link     string
}
