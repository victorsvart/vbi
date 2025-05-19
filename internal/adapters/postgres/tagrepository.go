package postgres

import (
	"context"
	"errors"
	"strings"

	"github.com/victorsvart/vbi/internal/core"
	"gorm.io/gorm"
)

type tagRepositoryImpl struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) core.TagRepository {
	return &tagRepositoryImpl{db}
}

func (tr *tagRepositoryImpl) GetAll(ctx context.Context) (t []core.Tag, err error) {
	tx := tr.db.WithContext(ctx)
	if err = tx.Find(&t).Error; err != nil {
		return t, err
	}

	return t, nil
}

func (tr *tagRepositoryImpl) UpdateTag(ctx context.Context, t core.Tag) (core.Tag, error) {
	tx := tr.db.WithContext(ctx)

	var currName string
	if err := tx.
		Where("id = ?", t.ID).
		Pluck("name", &currName).Error; err != nil {
		return core.Tag{}, err
	}

	newName := strings.TrimSpace(t.Name)
	if currName == newName {
		return core.Tag{}, errors.New("existing tag err: can't update a tag's name to itself")
	}

	if err := tx.
		Where("id = ?", t.ID).
		Update("name", newName).Error; err != nil {
		return core.Tag{}, err
	}

	t.Name = newName
	return t, nil
}
