package services

import (
	"context"
	"github.com/zakariawahyu/go-iris-news/entity"
	"github.com/zakariawahyu/go-iris-news/modules/channel"
	"github.com/zakariawahyu/go-iris-news/modules/content"
	"github.com/zakariawahyu/go-iris-news/modules/region"
	"github.com/zakariawahyu/go-iris-news/modules/sub_channel"
	"github.com/zakariawahyu/go-iris-news/pkg/exception"
	"github.com/zakariawahyu/go-iris-news/pkg/helpers"
	"strconv"
	"time"
)

type contentServices struct {
	contentRepo    content.ContentRepository
	channelRepo    channel.ChannelRepository
	subChannelRepo sub_channel.SubChannelRepository
	regionRepo     region.RegionRepository
	contextTimeout time.Duration
}

func NewContentServices(repo content.ContentRepository, channelRepo channel.ChannelRepository, subChannelRepo sub_channel.SubChannelRepository, regionRepo region.RegionRepository, timeout time.Duration) content.ContentServices {
	return &contentServices{
		contentRepo:    repo,
		channelRepo:    channelRepo,
		subChannelRepo: subChannelRepo,
		regionRepo:     regionRepo,
		contextTimeout: timeout,
	}
}

func (serv *contentServices) GetContentAllRow(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	var id string

	if types != "" {
		if types == "channel" {
			channel, err := serv.channelRepo.GetBySlugOrId(c, key)
			exception.PanicIfNeeded(err)

			id = strconv.FormatInt(channel.ID, 10)
		} else if types == "region" {
			region, err := serv.regionRepo.GetBySlugOrId(c, key)
			exception.PanicIfNeeded(err)

			id = strconv.FormatInt(region.ID, 10)
		} else if types == "subchannel" {
			subChannel, err := serv.subChannelRepo.GetBySlugOrId(c, key)
			exception.PanicIfNeeded(err)

			id = strconv.FormatInt(subChannel.ID, 10)
		} else if types != "channel" || types != "subchannel" || types != "region" {
			panic(helpers.ErrNotFound)
		}
	}

	res, err := serv.contentRepo.GetAllRow(c, types, id, limit, offset)
	exception.PanicIfNeeded(err)

	contents = entity.NewContentRowArrayResponse(res)

	return contents
}
