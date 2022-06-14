//go:generate mockgen -package=handler -source=handler.go -destination=./handler_mock.go Handler
package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/moynur/news-app/internal/models"
	"github.com/moynur/news-app/internal/service"
)

type LoadArticlesReq struct {
	Cursor   int    `json:"cursor,omitempty"`
	Category string `json:"category" json:"category,omitempty"`
	Provider string `json:"provider,omitempty"`
	Title    string `json:"title,omitempty"`
}

type Article struct {
	Title    string `json:"title,omitempty"`
	Summary  string `json:"summary,omitempty"`
	ImageRef string `json:"image_ref,omitempty"`
	Link     string `json:"link,omitempty"`
}

type LoadArticlesResp struct {
	NextCursor int       `json:"next_cursor,omitempty"`
	Articles   []Article `json:"articles,omitempty"`
}

type Handler struct {
	service service.Service
}

func NewHandler(svc service.Service) (*Handler, error) {
	return &Handler{
		service: svc,
	}, nil
}

func (h *Handler) ApplyRoutes(r *mux.Router) {
	r.HandleFunc("/loadArticles", h.LoadArticles).Methods(http.MethodGet)
}

func (h *Handler) LoadArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	NewDecoder := json.NewDecoder(r.Body)
	NewDecoder.DisallowUnknownFields()

	var newAuthRequest LoadArticlesReq
	var response LoadArticlesResp
	err := NewDecoder.Decode(&newAuthRequest)
	if err != nil {
		log.Println("\n Unable to decode req", err)
		errorBadRequest(w, "unable to decode request")
		return
	}
	resp, err := h.service.GetArticles(mapRequest(newAuthRequest))
	if err != nil {
		switch err {
		case service.ErrNotFound:
			errorNotFound(w, "no articles found")
		default:
			errorUnknownFailure(w, "failed to fetch articles")
		}
		return
	}
	log.Println(resp)
	response.NextCursor = resp.NextCursor
	for _, article := range resp.Articles {
		response.Articles = append(response.Articles, Article{
			Title:    article.Title,
			Summary:  article.Summary,
			ImageRef: article.ImageRef,
			Link:     article.Link,
		})
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("failure to write resp", err)
		errorUnknownFailure(w, "unknown failure")
		return
	}
}

func (h *Handler) GetArticle(w http.ResponseWriter, r *http.Request) {
	// would be responsible for taking a single article by URL/ID and fetching from db
	// then converting the service into a HTML template to consume
	//h.service.Getpege() for example
}

func mapRequest(req LoadArticlesReq) models.GetArticlesRequest {
	return models.GetArticlesRequest{
		Cursor:   req.Cursor,
		Category: req.Category,
		Provider: req.Provider,
		Title:    req.Title,
	}
}

func errorNotFound(w http.ResponseWriter, message string) {
	writeError(w, message, http.StatusNotFound)
}

func errorUnknownFailure(w http.ResponseWriter, message string) {
	writeError(w, message, http.StatusInternalServerError)
}

func errorBadRequest(w http.ResponseWriter, message string) {
	writeError(w, message, http.StatusBadRequest)
}

func errorUnprocessable(w http.ResponseWriter, message string) {
	writeError(w, message, http.StatusUnprocessableEntity)
}

func writeError(w http.ResponseWriter, message string, code int) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)

	if encoderErr := enc.Encode(code); err != nil {
		err = encoderErr
	}
}
