package models

type GetArticlesRequest struct {
	Cursor   int
	Category string
	Provider string
}

type GetArticlesResponse struct {
	NextCursor int
	Articles   []Article
}

type Article struct {
	Title    string
	Summary  string
	ImageRef string
}
