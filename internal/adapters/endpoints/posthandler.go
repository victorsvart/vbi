package endpoints

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/victorsvart/vbi/internal/core"
	"github.com/victorsvart/vbi/pkg/response"
)

var (
	ErrIdIsRequired = errors.New("id is required")
)

type PostHandler struct {
	serv core.PostService
}

func NewPostHandler(api chi.Router, serv core.PostService) {
	handler := PostHandler{serv}
	api.Route("/post", func(r chi.Router) {
		r.Get("/", handler.GetAll)
		r.Get("/{id}", handler.Get)
		r.Post("/", handler.Get)
		r.Put("/", handler.Update)
	})
}

func (ph *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	posts, err := ph.serv.GetAll(r.Context())
	if err != nil {
		response.Send(w, http.StatusInternalServerError, err)
		return
	}

	response.Send(w, http.StatusOK, posts)
}

func (ph *PostHandler) Get(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")
	if idPath == "" {
		response.Send(w, http.StatusBadRequest, ErrIdIsRequired)
		return
	}

	id, err := strconv.ParseUint(idPath, 10, 64)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, err)
		return
	}

	post, err := ph.serv.Get(r.Context(), uint(id))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, err)
		return
	}

	response.Send(w, http.StatusOK, post)
}

func (ph *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var i core.PostInput
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		response.Send(w, http.StatusBadRequest, err)
		return
	}

	post, err := ph.serv.Create(r.Context(), i)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, err)
		return
	}

	response.Send(w, http.StatusCreated, post)
}

func (ph *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	var i core.PostInput
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		response.Send(w, http.StatusBadRequest, err)
		return
	}

	post, err := ph.serv.Update(r.Context(), i)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, err)
		return
	}

	response.Send(w, http.StatusOK, post)
}
