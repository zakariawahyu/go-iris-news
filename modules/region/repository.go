package region

import (
	"context"
	"github.com/zakariawahyu/go-iris-news/entity"
)

type RegionRepository interface {
	GetAll(ctx context.Context) ([]*entity.Region, error)
	GetBySlugOrId(ctx context.Context, slug string) (*entity.Region, error)
}
