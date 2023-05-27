package repository

import (
	"context"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/zakariawahyu/go-iris-news/entity"
	"github.com/zakariawahyu/go-iris-news/helpers"
)

type ContentRepositoryImpl struct {
	Conn rel.Repository
}

func NewContentRepository(Conn rel.Repository) ContentRepository {
	return &ContentRepositoryImpl{Conn}
}

func (repo *ContentRepositoryImpl) GetBySlug(ctx context.Context, slug string) (entity.Content, error) {
	content := entity.Content{}

	query := rel.Select().Where(where.Eq("slug", slug)).Limit(1)

	err := repo.Conn.Find(ctx, &content, query)
	if err != nil {
		return content, helpers.ErrNotFound
	}

	if err := repo.Conn.Preload(ctx, &content, "user"); err != nil {
		return content, err
	}

	if err := repo.Conn.Preload(ctx, &content, "region"); err != nil {
		return content, err
	}

	if err := repo.Conn.Preload(ctx, &content, "channel"); err != nil {
		return content, err
	}

	if err := repo.Conn.Preload(ctx, &content, "sub_channel"); err != nil {
		return content, err
	}

	if err := repo.Conn.Preload(ctx, &content, "content_has_tag"); err != nil {
		return content, err
	}

	if err := repo.Conn.Preload(ctx, &content.ContentHasTag, "tag"); err != nil {
		return content, err
	}

	if err := repo.Conn.Preload(ctx, &content, "content_has_topic"); err != nil {
		return content, err
	}

	if err := repo.Conn.Preload(ctx, &content.ContentHasTopic, "topic"); err != nil {
		return content, err
	}

	if err := repo.Conn.Preload(ctx, &content, "content_has_reporter"); err != nil {
		return content, err
	}

	if err := repo.Conn.Preload(ctx, &content.ContentHasReporter, "reporter"); err != nil {
		return content, err
	}

	if err := repo.Conn.Preload(ctx, &content, "sub_photo"); err != nil {
		return content, err
	}

	return content, nil
}
