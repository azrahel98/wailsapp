package models

import "github.com/golang-jwt/jwt/v5"

type Usuario struct {
	ID       int    `db:"id"`
	Nombre   string `db:"nombre"`
	Password string `db:"pass"`
	Nick     string `db:"nickname"`
	Role     int    `db:"nivel"`
}

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
