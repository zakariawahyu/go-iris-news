package server

import (
	"github.com/zakariawahyu/go-iris-news/modules/channel"
	_channelServices "github.com/zakariawahyu/go-iris-news/modules/channel/services"
	"github.com/zakariawahyu/go-iris-news/modules/content"
	_contentServices "github.com/zakariawahyu/go-iris-news/modules/content/services"
	"github.com/zakariawahyu/go-iris-news/modules/region"
	_regionServices "github.com/zakariawahyu/go-iris-news/modules/region/services"
	"github.com/zakariawahyu/go-iris-news/modules/sub_channel"
	_subChannelServices "github.com/zakariawahyu/go-iris-news/modules/sub_channel/services"
	"time"
)

type Services struct {
	contentServices    content.ContentServices
	channelServices    channel.ChannelServices
	subChannelServices sub_channel.SubChannelServices
	regionServices     region.RegionServices
}

func NewServices(repo *Repository, timeoutContext time.Duration) *Services {
	return &Services{
		contentServices:    _contentServices.NewContentServices(repo.ContentRepo, repo.channelRepo, repo.subChannelRepo, repo.regionRepo, timeoutContext),
		channelServices:    _channelServices.NewChannelServices(repo.channelRepo, timeoutContext),
		subChannelServices: _subChannelServices.NewSubChannelServices(repo.subChannelRepo, timeoutContext),
		regionServices:     _regionServices.NewRegionServices(repo.regionRepo, timeoutContext),
	}
}
