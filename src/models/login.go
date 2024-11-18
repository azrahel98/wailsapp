package models

type Usuario struct {
	ID       int    `db:"id"`
	Nombre   string `db:"nombre"`
	Password string `db:"pass"`
	Nick     string `db:"nickname"`
	Role     int    `db:"nivel"`
}
