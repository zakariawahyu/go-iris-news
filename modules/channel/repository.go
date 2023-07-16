package channel

import (
	"context"
	"github.com/zakariawahyu/go-iris-news/entity"
)

type ChannelRepository interface {
	GetAll(ctx context.Context) ([]*entity.Channel, error)
	GetBySlugOrId(ctx context.Context, slug string) (*entity.Channel, error)
}
