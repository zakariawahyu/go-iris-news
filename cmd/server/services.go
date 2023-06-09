package server

import (
	_contentServices "github.com/zakariawahyu/go-iris-news/modules/content/services"
	"time"
)

type Services struct {
	contentServices _contentServices.ContentServices
}

func NewServices(repo *Repository, timeoutContext time.Duration) *Services {
	return &Services{
		contentServices: _contentServices.NewContentServices(repo.ContentRepo, timeoutContext),
	}
}
