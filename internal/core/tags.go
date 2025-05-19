package core

import (
	"context"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name string `json:"name"`
}

type TagInput struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (ti *TagInput) ToTag() Tag {
	return Tag{
		Model: gorm.Model{ID: ti.Id},
		Name:  ti.Name,
	}
}

type TagRepository interface {
	GetAll(context.Context) (t []Tag, err error)
	UpdateTag(context.Context, Tag) (Tag, error)
}

type TagService interface {
	GetAll(context.Context) (t []Tag, err error)
	UpdateTag(context.Context, TagInput) (Tag, error)
}
