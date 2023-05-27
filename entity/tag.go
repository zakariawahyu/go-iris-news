package entity

type Tag struct {
	ID   int64  `db:"id,primary" json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (Tag) Table() string {
	return "tags"
}
