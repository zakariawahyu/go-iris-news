package entity

type Reporter struct {
	ID    int64  `id:"id,primary" json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Image string `json:"image"`
}

func (Reporter) Table() string {
	return "reporters"
}
