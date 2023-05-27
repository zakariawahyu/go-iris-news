package server

import (
	_contentRepository "github.com/zakariawahyu/go-iris-news/modules/content/repository"
	"github.com/zakariawahyu/go-iris-news/utils"
)

type Repository struct {
	ContentRepo _contentRepository.ContentRepository
}

func NewRepository(conn *utils.Conn) *Repository {
	return &Repository{
		ContentRepo: _contentRepository.NewContentRepository(conn.Mysql),
	}
}
