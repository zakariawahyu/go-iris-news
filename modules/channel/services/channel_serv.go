package services

import (
	"context"
	"github.com/zakariawahyu/go-iris-news/entity"
	"github.com/zakariawahyu/go-iris-news/modules/channel"
	"github.com/zakariawahyu/go-iris-news/pkg/exception"
	"time"
)

type channelServices struct {
	channelRepo    channel.ChannelRepository
	contextTimeout time.Duration
}

func NewChannelServices(channelRepo channel.ChannelRepository, timeout time.Duration) channel.ChannelServices {
	return &channelServices{
		channelRepo:    channelRepo,
		contextTimeout: timeout,
	}
}

func (serv *channelServices) GetAllChannel(ctx context.Context) (channels []entity.ChannelResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	res, err := serv.channelRepo.GetAll(c)
	exception.PanicIfNeeded(err)

	for _, channel := range res {
		channels = append(channels, entity.NewChannelResponse(channel))
	}

	return channels
}

func (serv *channelServices) GetChannelBySlugOrId(ctx context.Context, slug string) entity.ChannelResponse {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	channel, err := serv.channelRepo.GetBySlugOrId(c, slug)
	exception.PanicIfNeeded(err)

	return entity.NewChannelResponse(channel)
}
