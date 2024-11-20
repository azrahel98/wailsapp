package services

import (
	"context"
	"vesgoapp/src/repositories"
)

type PersonalService struct {
	repo repositories.PersonalRepository
}

func NewPersonalService(r repositories.PersonalRepository) *PersonalService {
	return &PersonalService{repo: r}
}

func (s *PersonalService) Buscar_persona_by_dni(dni string) (*string, error) {
	_, err := s.repo.Search_by_dni_perfil(context.Background(), dni)
	if err != nil {
		return nil, err
	}
	ses := "asdfadsf"
	return &ses, nil
}
