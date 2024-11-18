package repositories

import (
	"context"
	"os"
	"vesgoapp/src/models"

	"github.com/jmoiron/sqlx"
)

type LoginRepository interface {
	Search_like_nickname(ctx context.Context, nickname string) (*models.Usuario, error)
}

type loginRepository struct {
	db *sqlx.DB
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
