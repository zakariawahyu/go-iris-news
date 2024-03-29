package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-iris-news/entity"
	"github.com/zakariawahyu/go-iris-news/modules/channel"
)

type channelRepository struct {
	DB *bun.DB
}

func NewChannelRepository(DB *bun.DB) channel.ChannelRepository {
	return &channelRepository{
		DB: DB,
	}
}

func (repo *channelRepository) GetAll(ctx context.Context) ([]*entity.Channel, error) {
	channel := []*entity.Channel{}

	if err := repo.DB.NewSelect().Model(&channel).Relation("Suplemens").Relation("SubChannels").Scan(ctx); err != nil {
		return nil, err
	}

	return channel, nil
}

func (repo *channelRepository) GetBySlugOrId(ctx context.Context, slug string) (*entity.Channel, error) {
	channel := &entity.Channel{}

	if err := repo.DB.NewSelect().Model(channel).Relation("Suplemens").Relation("SubChannels").Where("channel.slug = ?", slug).WhereOr("channel.id = ?", slug).Scan(ctx); err != nil {
		return nil, err
	}

	return channel, nil
}
