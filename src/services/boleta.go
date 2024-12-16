package services

import (
	"context"
	"fmt"
	"os"
	"vesgoapp/src/models"
	"vesgoapp/src/repositories"
)

type BoletaService struct {
	repo repositories.BoletaRepository
	empl repositories.PersonalRepository
}

func NewBoletaService(r repositories.BoletaRepository, e repositories.PersonalRepository) *BoletaService {
	return &BoletaService{repo: r, empl: e}
}

func (s *BoletaService) Listar_Xmls() (*[]models.ListOfFileDirectory, error) {
	lista, err := s.repo.ReadXmls_folder()
	if err != nil {
		return nil, err
	}
	return lista, nil
}

func (s *BoletaService) ReadXmls_folder() (*[]models.Boleta, error) {

	var resultadoFinal []models.Boleta
	lista, err := s.repo.ReadXmls_folder()
	if err != nil {
		return nil, err
	}
	key := os.Getenv("XMLPATH")

	for _, x := range *lista {
		resul, err := s.empl.IsCasbyDni(context.Background(), x.Dni)
		if err != nil {
			return nil, err
		}

		if resul {
			bolstru, err := s.repo.Extraer_Datos(fmt.Sprintf("%s/%s", key, x.Full))
			if err != nil {
				fmt.Println("Aquiiiiiiiiiiiiiiii el errorr")
			}
			resultadoFinal = append(resultadoFinal, *bolstru)

			MakePdf(x.Dni, *bolstru)
		}
	}

	return &resultadoFinal, nil
}
