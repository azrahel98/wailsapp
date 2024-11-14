package repositories

import (
	"context"
	"fmt"
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

	err := l.db.GetContext(ctx, &user, "select id,nombre,pass,nickname,lvl from usuario where nickname = ?", nickname)
	if err != nil {
		return nil, fmt.Errorf("Error")
	}
	return &user, nil
}

func CreateLoginRepo(db *sqlx.DB) LoginRepository {
	return &loginRepository{db}
}
