package channel

import (
	"context"
	"github.com/zakariawahyu/go-iris-news/entity"
)

type ChannelServices interface {
	GetAllChannel(ctx context.Context) (channels []entity.ChannelResponse)
	GetChannelBySlugOrId(ctx context.Context, slug string) entity.ChannelResponse
}
