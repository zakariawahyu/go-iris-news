package sub_channel

import (
	"context"
	"github.com/zakariawahyu/go-iris-news/entity"
)

type SubChannelRepository interface {
	GetAll(ctx context.Context) ([]*entity.SubChannel, error)
	GetBySlugOrId(ctx context.Context, slug string) (*entity.SubChannel, error)
}
