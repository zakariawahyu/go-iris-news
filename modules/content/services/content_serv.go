package services

import (
	"context"
	"github.com/zakariawahyu/go-iris-news/entity"
)

type ContentServices interface {
	GetBySlug(c context.Context, slug string) entity.ContentResponse
}
