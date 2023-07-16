package repository

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zakariawahyu/go-iris-news/entity"
	"github.com/zakariawahyu/go-iris-news/modules/content"
)

type contentRepository struct {
	DB *bun.DB
}

func NewContentRepository(Conn *bun.DB) content.ContentRepository {
	return &contentRepository{Conn}
}

func (repo *contentRepository) GetAllRow(ctx context.Context, types string, key string, limit int, offset int) ([]*entity.ContentRowResponse, error) {
	content := []*entity.ContentRowResponse{}

	if err := repo.DB.NewSelect().Model(&content).
		Where("content_row_response.is_active = ?", true).
		Apply(func(q *bun.SelectQuery) *bun.SelectQuery { //Relation Function
			if types == "channel" || types == "subchannel" {
				q = q.Relation("Channel").Relation("SubChannel")
			} else if types == "region" {
				q = q.Relation("Region").Relation("SubPhotos")
			} else { //home
				q = q.Relation("Channel").Relation("SubChannel").Relation("Region").Relation("SubPhotos")
			}
			return q
		}).
		Apply(func(q *bun.SelectQuery) *bun.SelectQuery { //Where Function
			if types == "channel" {
				q = q.Where("headline_type NOT IN (?)", bun.In([]int64{1, 2})).Where("type = ?", "channel").Where("type_id = ?", key)
			} else if types == "subchannel" {
				q = q.Where("headline_type = 0").Where("type = ?", "channel").Where("type_child_id = ?", key)
			} else if types == "region" {
				q = q.Where("headline_type NOT IN (?)", bun.In([]int{1, 2}))
				if key == "1189" {
					q = q.Where("type = ?", "region")
				} else {
					q = q.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type = ?", "region").Where("type_id = ?", key)
					}).WhereGroup(" OR ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type = ?", "photo").Where("type_id = ?", key).Where("type_child_id is null")
					}).WhereGroup(" OR ", func(q *bun.SelectQuery) *bun.SelectQuery {
						return q.Where("type = ?", "video").Where("type_id = ?", key).Where("type_child_id is null")
					})
				}
			} else { //home
				q = q.Where("headline_type != 1").Where("is_national = 1")
			}
			return q
		}).
		WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("ads_position is null").WhereOr("ads_position = 0")
		}).
		Order("published_date desc").
		Limit(limit).
		Offset(offset).
		Scan(ctx); err != nil {
		return nil, err
	}

	return content, nil
}
