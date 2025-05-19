package core

import (
	"context"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Hypertext     string
	Title         string
	Summary       string
	HeaderImage   string
	ViewCount     uint64
	AllowComments bool
	Active        bool
	Comments      []Comment
	Tags          []Tag `gorm:"many2many:post_tags;"`
}

type PostRepository interface {
	GetAll(context.Context) ([]Post, error)
	Get(context.Context, uint) (Post, error)
	GetByTag(context.Context, uint) (p []Post, err error)
	Create(context.Context, *Post) error
	Update(context.Context, *Post) error
}

type PostService interface {
	GetAll(context.Context) ([]Post, error)
	Get(context.Context, uint) (Post, error)
	GetByTag(context.Context, uint) (p []Post, err error)
	Create(context.Context, PostInput) (Post, error)
	Update(context.Context, PostInput) (Post, error)
}

type PostInput struct {
	ID          *uint  `json:"id"`
	Hypertext   string `json:"hypertext"`
	Title       string `json:"title"`
	Summary     string `json:"summary"`
	HeaderImage string `json:"headerImage"`
}

func (pi PostInput) ToPost() Post {
	id := uint(0)
	if pi.ID != nil {
		id = *pi.ID
	}

	post := Post{
		Model:       gorm.Model{ID: id},
		Hypertext:   pi.Hypertext,
		Title:       pi.Title,
		Summary:     pi.Summary,
		HeaderImage: pi.HeaderImage,
	}

	return post
}

func NewPost(i PostInput) Post {
	return Post{
		Hypertext:   i.Hypertext,
		Title:       i.Title,
		Summary:     i.Summary,
		HeaderImage: i.HeaderImage,
	}
}

func (p *Post) SetActive() {
	p.Active = true
}

func (p *Post) SetUnactive() {
	p.Active = true
}

func (p *Post) EnableComments() {
	p.AllowComments = true
}

func (p *Post) BlockComments() {
	p.AllowComments = false
}
