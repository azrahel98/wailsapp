package services

import (
	"context"
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

func (s *BoletaService) Upload_file() (bool, error) {
	if s.ctx == nil {
		return false, context.Canceled
	}

	filePath, err := runtime.OpenFileDialog(
		s.ctx,
		runtime.OpenDialogOptions{
			Title: "Seleccionar el Archivo Excel",
			Filters: []runtime.FileFilter{
				{DisplayName: "Todos los Archivos", Pattern: "*.xlsx"},
			},
		},
	)

	if err != nil {
		return false, err
	}

	if filePath == "" {
		return false, nil
	}

	res, err := s.repo.Upload_file(filePath)
	if err != nil {
		return false, err
	}

	return res, nil
}
