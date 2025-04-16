package entity

type Users struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	// tambahkan kolom lain sesuai kebutuhan
}

func (Users) TableName() string {
	return "users"
}