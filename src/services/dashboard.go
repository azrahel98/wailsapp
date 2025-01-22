package services

import (
	"context"
	"vesgoapp/src/models"
	"vesgoapp/src/repositories"
)

type DashboardService struct {
	repo repositories.DashboardRepository
}

func NewDashboardService(r repositories.DashboardRepository) *DashboardService {
	return &DashboardService{repo: r}
}

func (s *DashboardService) Regimenes_resumen() (*[]models.RegimenesCantidad, error) {
	res, err := s.repo.Cantidad_by_regimen(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DashboardService) Sexo_cantidad() (*[]models.RegimenesCantidad, error) {
	res, err := s.repo.Cantidad_by_sexo(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DashboardService) Cumpleaños(mes int) (*[]models.Cumpleaños_Activos, error) {
	res, err := s.repo.Cumpleaño_by_activos(context.Background(), mes)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DashboardService) Buscar_trabajador(nombre string) (*[]models.Buscar_trabajador, error) {
	res, err := s.repo.Buscar_by_nombre(context.Background(), nombre)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DashboardService) Trabajadore_Activos_Area() (*[]models.RegimenesCantidad, error) {
	res, err := s.repo.Trabajadores_activos(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DashboardService) Resumen_Dashboard() (*[]models.PersonaActivo, error) {
	res, err := s.repo.Cantidad_vincolos_activos(context.Background())
	if err != nil {
		return nil, err
	}
	renuncias, err := s.repo.Cantidad_renuncias_mes(context.Background())
	if err != nil {
		return nil, err
	}

	var result []models.PersonaActivo
	result = append(result, *res)
	result = append(result, *renuncias...)

	return &result, nil
}
