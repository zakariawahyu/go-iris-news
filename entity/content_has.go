package entity

type ContentHasTag struct {
	ID        int64 `db:"id,primary" json:"id"`
	TagID     int64 `db:"tag_id" json:"tag_id"`
	Tag       *Tag  `ref:"tag_id" fk:"id" json:"tag"`
	ContentID int64 `db:"content_id" json:"content_id"`
}

type ContentHasTopic struct {
	ID        int64  `db:"id,primary" json:"id"`
	TopicID   int64  `bun:"tag_id" json:"topic_id"`
	Topic     *Topic `ref:"topic_id" fk:"id" json:"topic"`
	ContentID int64  `db:"content_id" json:"content_id" json:"content_id"`
}

type ContentHasReporter struct {
	ID         int64     `db:"id,primary" json:"id"`
	ReporterID int64     `db:"reporter_id" json:"reporter_id"`
	Reporter   *Reporter `ref:"reporter_id" fk:"id" json:"reporter"`
	ContentID  int64     `db:"content_id" json:"content_id"`
}
