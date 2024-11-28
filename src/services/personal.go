package services

import (
	"context"
	"vesgoapp/src/models"
	"vesgoapp/src/repositories"
)

type PersonalService struct {
	repo repositories.PersonalRepository
}

func NewPersonalService(r repositories.PersonalRepository) *PersonalService {
	return &PersonalService{repo: r}
}

func (s *PersonalService) Buscar_persona_by_dni(dni string) (*models.Perfil, error) {
	res, err := s.repo.Search_by_dni_perfil(context.Background(), dni)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PersonalService) Search_by_dni_vinculos(dni string) (*[]models.Vinculos, error) {
	res, err := s.repo.Search_by_dni_vinculos(context.Background(), dni)
	if err != nil {
		return nil, err
	}
	return res, nil
}
