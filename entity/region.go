package entity

type Region struct {
	ID   int64  `db:"id,primary" json:"id"`
	Name string `db:"name" json:"name"`
	Slug string `db:"slug" json:"slug"`
}

type RegionResponse struct {
	ID   int64  `db:"id,primary" json:"id"`
	Name string `db:"name" json:"name"`
	Slug string `db:"slug" json:"slug"`
}

func (u Region) Table() string {
	return "regions"
}

func (u RegionResponse) Table() string {
	return "regions"
}
