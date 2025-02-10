package services

import (
	"context"
	"fmt"
	"os"
	"vesgoapp/src/repositories"
)

type LoginService struct {
	repo repositories.LoginRepository
}

func NewLoginService(r repositories.LoginRepository) *LoginService {
	return &LoginService{repo: r}
}

func (s *LoginService) Login(nickname string, password string) (*string, error) {
	user, err := s.repo.Search_like_nickname(context.Background(), nickname)
	if err != nil {
		return nil, fmt.Errorf("el usuario no existe")
	}

	if user.Password != password {
		return nil, fmt.Errorf("la contraseña es incorrecta")
	}

	token, err := s.repo.Create_Token(user.ID, user.Role, user.Nick, os.Getenv("DBKEY"))
	if err != nil {
		return nil, fmt.Errorf("no se puedo generar el Token")
	}

	return &token, nil
}
