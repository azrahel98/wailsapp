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

func (s *PersonalService) EditByDni(telf1 string, telf2 string, direccion string, email string, dni string) error {
	err := s.repo.EditByDni(context.Background(), telf1, telf2, direccion, email, dni)
	if err != nil {
		return err
	}
	return nil
}

func (s *PersonalService) Search_by_dni_vinculos(dni string) (*[]models.Vinculos, error) {
	res, err := s.repo.Search_by_dni_vinculos(context.Background(), dni)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PersonalService) AddRenuncia(doc models.Documento, id int) error {
	_, err := s.repo.AddRenuncia(context.Background(), doc, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PersonalService) Buscar_dni_onlinne(dni string) (*models.PersonDniRequest, error) {

	res, err := s.repo.Search_dni_onlinne(context.Background(), dni)
	if err != nil {
		return nil, err
	}
	return res, nil
}
