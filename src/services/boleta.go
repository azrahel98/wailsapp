package services

import (
	"context"
	"vesgoapp/src/models"
	"vesgoapp/src/repositories"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type BoletaService struct {
	repo repositories.BoletaRepository
	ctx  context.Context
}

func NewBoletaService(r repositories.BoletaRepository) *BoletaService {
	return &BoletaService{repo: r}
}

func (s *BoletaService) SetContext(ctx context.Context) {
	s.ctx = ctx
}

func (s *BoletaService) Upload_file() (*models.Boleta, error) {
	if s.ctx == nil {
		return nil, context.Canceled
	}

	filePath, err := runtime.OpenFileDialog(
		s.ctx,
		runtime.OpenDialogOptions{
			Title: "Seleccionar el Archivo Excel",
			Filters: []runtime.FileFilter{
				{DisplayName: "Todos los Archivos", Pattern: "*.xml"},
			},
		},
	)

	if err != nil {
		return nil, err
	}

	if filePath == "" {
		return nil, nil
	}

	res, err := s.repo.Upload_file(filePath)
	if err != nil {
		return nil, err
	}

	return res, nil
}
