package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-iris-news/entity"
	"github.com/zakariawahyu/go-iris-news/modules/sub_channel"
)

type subChannelRepository struct {
	DB *bun.DB
}

func NewSubChannelRepository(DB *bun.DB) sub_channel.SubChannelRepository {
	return &subChannelRepository{
		DB: DB,
	}
}
func (repo *subChannelRepository) GetAll(ctx context.Context) ([]*entity.SubChannel, error) {
	subChannel := []*entity.SubChannel{}

	if err := repo.DB.NewSelect().Model(&subChannel).Relation("Channel").Scan(ctx); err != nil {
		return nil, err
	}

	return subChannel, nil
}

func (repo *subChannelRepository) GetBySlugOrId(ctx context.Context, slug string) (*entity.SubChannel, error) {
	subChannel := &entity.SubChannel{}

	if err := repo.DB.NewSelect().Model(subChannel).Relation("Channel").Where("sub_channel.slug = ? ", slug).WhereOr("sub_channel.id = ?", slug).Scan(ctx); err != nil {
		return nil, err
	}

	return subChannel, nil
}
