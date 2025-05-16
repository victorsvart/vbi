package core

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	hypertext     string
	title         string
	subtitle      string
	headerImage   string
	viewCount     uint64
	allowComments bool
	active        bool
	Comments      []Comment
}

type PostInput struct {
	ID          *uint  `json:"id"`
	Hypertext   string `json:"hypertext"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	HeaderImage string `json:"headerImage"`
}

func (pi PostInput) ToPost() Post {
	id := uint(0)
	if pi.ID != nil {
		id = *pi.ID
	}

	post := Post{
		Model:       gorm.Model{ID: id},
		hypertext:   pi.Hypertext,
		title:       pi.Title,
		subtitle:    pi.Subtitle,
		headerImage: pi.HeaderImage,
	}

	return post
}

func NewPost(i PostInput) Post {
	return Post{
		hypertext:   i.Hypertext,
		title:       i.Title,
		subtitle:    i.Subtitle,
		headerImage: i.HeaderImage,
	}
}

func (p *Post) SetActive() {
	p.active = true
}

func (p *Post) SetUnactive() {
	p.active = true
}

func (p *Post) AllowComments() {
	p.allowComments = true
}

func (p *Post) BlockComments() {
	p.allowComments = false
}
