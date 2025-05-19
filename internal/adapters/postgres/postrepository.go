package postgres

import (
	"context"

	"github.com/victorsvart/vbi/internal/core"
	"gorm.io/gorm"
)

type postRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) core.PostRepository {
	return &postRepositoryImpl{db}
}

func (pr *postRepositoryImpl) GetAll(ctx context.Context) ([]core.Post, error) {
	posts := make([]core.Post, 0)
	tx := pr.db.WithContext(ctx)
	if err := tx.Find(&posts).Error; err != nil {
		return posts, err
	}

	return posts, nil
}

func (pr *postRepositoryImpl) Get(ctx context.Context, id uint) (core.Post, error) {
	var p core.Post
	tx := pr.db.WithContext(ctx)
	if err := tx.Find(&p, id).Error; err != nil {
		return p, err
	}

	return p, nil
}

func (pr *postRepositoryImpl) GetByTag(ctx context.Context, tagID uint) (p []core.Post, err error) {
	tx := pr.db.WithContext(ctx)
	if err = tx.
		Joins("JOIN post_tags ON post_tags.post_id = posts.id").
		Where("post_tags.tag_id = ? ", tagID).
		Preload("Tags").
		Find(&p).Error; err != nil {
		return p, err
	}

	return p, nil
}

func (pr *postRepositoryImpl) Create(ctx context.Context, p *core.Post) error {
	tx := pr.db.WithContext(ctx)
	if err := tx.Create(&p).Error; err != nil {
		return err
	}

	return nil
}

func (pr *postRepositoryImpl) Update(ctx context.Context, p *core.Post) error {
	tx := pr.db.WithContext(ctx)
	if err := tx.Save(&p).Error; err != nil {
		return err
	}

	return nil
}
