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

func (s *DashboardService) Reporte_Sindicato(sindicato int) (*[]models.Reporte_Trabajadores, error) {
	res, err := s.repo.Reporte_personal_by_sindicato(context.Background(), sindicato)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *DashboardService) Reporte_Activos() (*[]models.Reporte_Trabajadores, error) {
	res, err := s.repo.Reporte_personal_by_activo(context.Background())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *DashboardService) Reporte_Renuncias(año int) (*[]models.Reporte_Trabajadores, error) {
	res, err := s.repo.Reporte_personal_by_renunciasAño(context.Background(), año)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type Organigrama struct {
	Id           int
	Area         string
	Jefe         *string
	Dni          *string
	Subgerencias []Organigrama
}

func (s *DashboardService) Reporte_Organigrama(año int) (*[]Organigrama, error) {
	res, err := s.repo.Report_funcionarios(context.Background())
	if err != nil {
		return nil, err
	}
	var organigrama []Organigrama

	for _, item := range *res {
		if item.Nivel == nil {
			org := Organigrama{
				Id:   item.Id,
				Area: item.Area,
				Dni:  item.Dni,
				Jefe: item.Nombre,
			}
			organigrama = append(organigrama, org)
		}
	}

	for _, item := range *res {
		if item.Nivel != nil {
			for i := range organigrama {
				if organigrama[i].Id == *item.Nivel {
					sub := Organigrama{
						Id:           item.Id,
						Area:         item.Area,
						Dni:          item.Dni,
						Jefe:         item.Nombre,
						Subgerencias: nil,
					}
					organigrama[i].Subgerencias = append(organigrama[i].Subgerencias, sub)
				}
			}
		}
	}

	for i := range organigrama {
		for j := range organigrama[i].Subgerencias {
			for _, item := range *res {
				if item.Nivel != nil && organigrama[i].Subgerencias[j].Id == *item.Nivel {
					organigrama[i].Subgerencias[j].Subgerencias = append(organigrama[i].Subgerencias[j].Subgerencias, Organigrama{
						Id:   item.Id,
						Area: item.Area,
						Dni:  item.Dni,
						Jefe: item.Nombre,
					})
				}
			}
		}
	}

	return &organigrama, nil
}
