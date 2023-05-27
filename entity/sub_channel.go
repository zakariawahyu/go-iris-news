package entity

type SubChannel struct {
	ID   int64  `db:"id,primary" json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type SubChannelResponse struct {
	ID   int64  `db:"id,primary" json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (SubChannel) Table() string {
	return "sub_channels"
}

func (SubChannelResponse) Table() string {
	return "sub_channels"
}
