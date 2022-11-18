package models

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username" form:"username" `
	Password string `db:"password" form:"password" `
}
