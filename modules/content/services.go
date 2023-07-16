package content

import (
	"context"
	"github.com/zakariawahyu/go-iris-news/entity"
)

type ContentServices interface {
	GetContentAllRow(ctx context.Context, types string, key string, limit int, offset int) (contents []entity.ContentRowResponse)
}
