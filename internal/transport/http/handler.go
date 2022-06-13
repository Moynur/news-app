//go:generate mockgen -package=handler -source=handler.go -destination=./handler_mock.go Handler
package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/moynur/gateway/internal/models"
	"github.com/moynur/gateway/internal/service"
)

type loadArticlesReq struct {
	Cursor   int
	Category string
	Provider string
}

type Article struct {
	Title    string
	Summary  string
	ImageRef string
}

type loadArticlesResp struct {
	NextCursor int
	Articles   []Article
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
	log.Println("something")
	w.Header().Set("Content-Type", "application/json")
	NewDecoder := json.NewDecoder(r.Body)
	NewDecoder.DisallowUnknownFields()

	var newAuthRequest loadArticlesReq
	err := NewDecoder.Decode(&newAuthRequest)
	if err != nil {
		log.Println("\n Unable to decode req", err)
		errorBadRequest(w, "unable to decode request")
		return
	}
	request := transformAuth(newAuthRequest)

	resp, err := h.service.GetArticles(request)
	if err != nil {
		log.Println("cant auth", err)
		errorUnknownFailure(w, "failed to auth")
		return
	}
	log.Println(resp)
	err = json.NewEncoder(w).Encode("handlerResp")
	if err != nil {
		log.Println("failure to write resp", err)
		errorUnknownFailure(w, "unknown failure")
		return
	}
}

func transformAuth(req loadArticlesReq) models.GetArticlesRequest {
	return models.GetArticlesRequest{
		Cursor:   req.Cursor,
		Category: req.Category,
		Provider: req.Provider,
	}
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
