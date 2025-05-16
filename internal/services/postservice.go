package services

import (
	"context"

	"github.com/victorsvart/vbi/internal/adapters/postgres"
	"github.com/victorsvart/vbi/internal/core"
)

type PostService interface {
	GetAll(context.Context) ([]core.Post, error)
	Get(context.Context, uint) (core.Post, error)
	Create(context.Context, core.PostInput) (core.Post, error)
	Update(context.Context, core.PostInput) (core.Post, error)
}

type postServiceImpl struct {
	repo postgres.PostRepository
}

func NewPostService(repo postgres.PostRepository) PostService {
	return &postServiceImpl{repo}
}

func (ps *postServiceImpl) GetAll(ctx context.Context) ([]core.Post, error) {
	posts, err := ps.repo.GetAll(ctx)
	if err != nil {
		return posts, err
	}

	return posts, nil
}

func (ps *postServiceImpl) Get(ctx context.Context, id uint) (core.Post, error) {
	p, err := ps.repo.Get(ctx, id)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (ps *postServiceImpl) Create(ctx context.Context, input core.PostInput) (core.Post, error) {
	p := input.ToPost()
	if err := ps.repo.Create(ctx, &p); err != nil {
		return p, err
	}

	return p, nil
}

func (ps *postServiceImpl) Update(ctx context.Context, input core.PostInput) (core.Post, error) {
	p := input.ToPost()
	if err := ps.repo.Update(ctx, &p); err != nil {
		return p, err
	}

	return p, nil
}
