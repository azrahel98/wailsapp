package repositories

import (
	"context"
	"fmt"
	"os"
	"time"
	"vesgoapp/src/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
)

type LoginRepository interface {
	Create_Token(userID, role int, username, secretKey string) (string, error)
	Search_like_nickname(ctx context.Context, nickname string) (*models.Usuario, error)
}

type loginRepository struct {
	db *sqlx.DB
}

func (l *loginRepository) Create_Token(userID, role int, username, secretKey string) (string, error) {
	claims := &models.Claims{
		UserID:   uint(userID),
		Username: username,
		Role:     fmt.Sprintf("%d", role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(48 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "mves",
			Subject:   fmt.Sprintf("%d", userID),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (l *loginRepository) Search_like_nickname(ctx context.Context, nickname string) (*models.Usuario, error) {
	var user models.Usuario

	key := os.Getenv("DBKEY")
	err := l.db.GetContext(ctx, &user, "select id,nombre,cast(aes_decrypt(pass,?) as char) pass,nickname,nivel from usuario where nickname = ?", key, nickname)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateLoginRepo(db *sqlx.DB) LoginRepository {
	return &loginRepository{db}
}
