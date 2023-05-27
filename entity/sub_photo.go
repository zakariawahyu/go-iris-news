package entity

type SubPhoto struct {
	ID        int64  `db:"id,primary" json:"id"`
	ContentID int64  `json:"content_id"`
	Content   string `json:"content"`
	Image     string `json:"image"`
}
