package core

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	commenter string
	message   string
	active    bool
	postID    uint
}

type CommentInput struct {
	Commenter string `json:"commenter"`
	Message   string `json:"message"`
	PostID    uint   `json:"postID"`
}

func NewComment(c CommentInput) Comment {
	return Comment{
		commenter: c.Commenter,
		message:   c.Message,
	}
}

func (c *Comment) SetActive() {
	c.active = true
}

func (c *Comment) SetUnactive() {
	c.active = true
}
