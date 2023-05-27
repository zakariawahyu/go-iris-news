package repository

import (
	"github.com/zakariawahyu/go-iris-news/entity"
	"golang.org/x/net/context"
)

type ContentRepository interface {
	GetBySlug(ctx context.Context, slug string) (entity.Content, error)
}
