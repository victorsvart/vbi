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
	ErrIdIsRequired    = errors.New("id is required")
	ErrTagIdIsRequired = errors.New("tag id is required")
)

type PostHandler struct {
	serv core.PostService
}

func NewPostHandler(api chi.Router, serv core.PostService) {
	handler := PostHandler{serv}
	api.Route("/post", func(r chi.Router) {
		r.Get("/", handler.GetAll)
		r.Get("/{id}", handler.Get)
		r.Get("/{tagId}", handler.GetByTag)
		r.Post("/", handler.Get)
		r.Put("/", handler.Update)
	})
}

func (ph *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	posts, err := ph.serv.GetAll(r.Context())
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendOk(w, posts)
}

func (ph *PostHandler) Get(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")
	if idPath == "" {
		response.SendBadRequest(w, ErrIdIsRequired)
		return
	}

	id, err := strconv.ParseUint(idPath, 10, 64)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	post, err := ph.serv.Get(r.Context(), uint(id))
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendOk(w, post)
}

func (ph *PostHandler) GetByTag(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")
	if idPath == "" {
		response.SendBadRequest(w, ErrTagIdIsRequired)
		return
	}

	tagID, err := strconv.ParseUint(idPath, 10, 64)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	p, err := ph.serv.GetByTag(r.Context(), uint(tagID))
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendOk(w, p)
}

func (ph *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var i core.PostInput
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		response.SendBadRequest(w, err)
		return
	}

	post, err := ph.serv.Create(r.Context(), i)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendCreated(w, post)
}

func (ph *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	var i core.PostInput
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		response.SendBadRequest(w, err)
		return
	}

	post, err := ph.serv.Update(r.Context(), i)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}
	response.SendOk(w, post)
}
