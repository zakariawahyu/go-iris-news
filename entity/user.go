package entity

type User struct {
	ID        int64  `db:"id,primary" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}

type UserResponse struct {
	ID        int64  `db:"id,primary" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}

func (User) Table() string {
	return "users"
}

func (UserResponse) Table() string {
	return "users"
}
