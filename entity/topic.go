package entity

type Topic struct {
	ID   int64  `db:"id,primary" json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (Topic) Table() string {
	return "topics"
}
