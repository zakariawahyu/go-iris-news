package content

import (
	"context"
	"github.com/zakariawahyu/go-iris-news/entity"
)

type ContentRepository interface {
	GetAllRow(ctx context.Context, types string, key string, limit int, offset int) ([]*entity.ContentRowResponse, error)
}
