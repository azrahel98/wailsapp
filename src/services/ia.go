package services

import (
	"context"
	"fmt"
	"strings"
	"vesgoapp/src/repositories"
)

type IaService struct {
	repo repositories.IaRepository
}

func NewIaRepository(r repositories.IaRepository) *IaService {
	return &IaService{repo: r}
}

func (s *IaService) Consulta_IA(prompt string) (*string, error) {
	res, err := s.repo.Crear_Query(prompt)
	if err != nil {
		return nil, fmt.Errorf("el error es este generando la consulta")
	}

	limpio := strings.TrimPrefix(*res, "```sql\n")
	limpio = strings.TrimSuffix(limpio, "```")
	query, err := s.repo.Ejecutar_Query(context.Background(), limpio)
	if err != nil {
		return nil, fmt.Errorf("el error es: %s y la consulta es: %s", err, *res)
	}

	response, err := s.repo.Humanizar_Response(prompt, query)
	if err != nil {
		return nil, err
	}
	return response, nil

}
