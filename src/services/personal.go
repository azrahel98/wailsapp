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

func (s *PersonalService) EditByDni(telf1 string, telf2 string, direccion string, email string, ruc string, dni string) error {
	err := s.repo.EditByDni(context.Background(), telf1, telf2, direccion, email, dni, ruc)
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

func (s *PersonalService) Buscar_Areas(area string) (*[]models.AreaCargoSerch, error) {
	res, err := s.repo.Search_Areas_count(context.Background(), area)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PersonalService) Buscar_Cargos(cargo string) (*[]models.AreaCargoSerch, error) {
	res, err := s.repo.Search_Cargo_count(context.Background(), cargo)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PersonalService) Crear_nuevoTrabajador(info models.Data, existe bool) (*string, error) {
	tx, err := s.repo.TxBegin(context.Background())
	if err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	if existe {
		doc, err := s.repo.Create_documento(tx, info.Documento)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		info.Vinculo.Dingreso = int(*doc)
		info.Vinculo.Dni = info.Persona.Dni
		_, err = s.repo.Crear_Vinculo(tx, info.Vinculo)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		return &info.Persona.Dni, nil
	} else {
		per, err := s.repo.Create_newIUser(tx, info.Persona)
		if err != nil {

			tx.Rollback()
			return nil, err
		}
		doc, err := s.repo.Create_documento(tx, info.Documento)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		info.Vinculo.Dingreso = int(*doc)
		info.Vinculo.Dni = per.Dni
		_, err = s.repo.Crear_Vinculo(tx, info.Vinculo)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		return &info.Persona.Dni, err
	}
}
