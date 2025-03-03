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

func (s *DashboardService) Resumen_Dashboard() (*models.ResumenIndicadores, error) {
	res, err := s.repo.Cantidad_vincolos_activos(context.Background())
	if err != nil {
		return nil, err
	}
	renuncias, err := s.repo.Cantidad_renuncias_mes(context.Background())
	if err != nil {
		return nil, err
	}
	regimenes, err := s.repo.Cantidad_by_regimen(context.Background())
	if err != nil {
		return nil, err
	}
	sex, err := s.repo.Cantidad_by_sexo(context.Background())
	if err != nil {
		return nil, err
	}
	sindicatos, err := s.repo.Cantidad_Activos_Sindicato(context.Background())
	if err != nil {
		return nil, err
	}
	var result models.ResumenIndicadores

	result.Personalregistrado = *res
	result.RenunciasMes = *renuncias
	result.Regimenes = *regimenes
	result.Sexo = *sex
	result.Sindicatos = *sindicatos

	return &result, nil
}

func (s *DashboardService) Trabajadores_por_area(nombre string) (*[]models.TrabajadoresArea, error) {
	res, err := s.repo.Trabajadores_area(context.Background(), nombre)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *DashboardService) Resumen_Regiemenes() (*[]models.RegimenesResumen, error) {
	res, err := s.repo.Resumen_regimenes_activos(context.Background())
	if err != nil {
		return nil, err
	}

	return res, nil
}
