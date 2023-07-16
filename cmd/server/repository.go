package server

import (
	"github.com/zakariawahyu/go-iris-news/modules/channel"
	_channelRepository "github.com/zakariawahyu/go-iris-news/modules/channel/repository"
	"github.com/zakariawahyu/go-iris-news/modules/content"
	_contentRepository "github.com/zakariawahyu/go-iris-news/modules/content/repository"
	"github.com/zakariawahyu/go-iris-news/modules/region"
	_regionRepo "github.com/zakariawahyu/go-iris-news/modules/region/repository"
	"github.com/zakariawahyu/go-iris-news/modules/sub_channel"
	_subChannelRepository "github.com/zakariawahyu/go-iris-news/modules/sub_channel/repository"
	"github.com/zakariawahyu/go-iris-news/pkg/db"
)

type Repository struct {
	ContentRepo    content.ContentRepository
	channelRepo    channel.ChannelRepository
	subChannelRepo sub_channel.SubChannelRepository
	regionRepo     region.RegionRepository
}

func NewRepository(conn *db.Conn) *Repository {
	return &Repository{
		ContentRepo:    _contentRepository.NewContentRepository(conn.Mysql),
		channelRepo:    _channelRepository.NewChannelRepository(conn.Mysql),
		subChannelRepo: _subChannelRepository.NewSubChannelRepository(conn.Mysql),
		regionRepo:     _regionRepo.NewRegionRepository(conn.Mysql),
	}
}
