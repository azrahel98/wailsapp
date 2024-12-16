package repositories

import (
	"context"
	"fmt"
	"os"
	"vesgoapp/src/models"

	"github.com/jmoiron/sqlx"
)

type PersonalRepository interface {
	Search_by_dni_perfil(ctx context.Context, dni string) (*models.Perfil, error)
	Search_by_dni_vinculos(ctx context.Context, dni string) (*[]models.Vinculos, error)
	IsCasbyDni(ctx context.Context, dni string) (bool, error)
	EditByDni(ctx context.Context, telf1 string, telf2 string, direccion string, emai string, dni string) error
}

type personalRepository struct {
	db *sqlx.DB
}

func (p *personalRepository) EditByDni(ctx context.Context, telf1 string, telf2 string, direccion string, emai string, dni string) error {
	key := os.Getenv("DBKEY")
	query := fmt.Sprintf("update Persona set direccion = aes_encrypt('%s','%s'), telf1= aes_encrypt('%s','%s'), telf2 = aes_encrypt('%s','%s'), email = aes_encrypt('%s','%s') where dni = '%s'", direccion, key, telf1, key, telf2, key, emai, key, dni)

	last, err := p.db.Exec(query)
	if err != nil {
		return err
	}

	_, err = last.RowsAffected()
	if err != nil {
		return err
	}

	return nil

}

func (p *personalRepository) IsCasbyDni(ctx context.Context, dni string) (bool, error) {
	var status string

	query := `
	select
  p.dni
from
  Persona p
  inner join Vinculo v on p.dni = v.dni
  inner join Regimen r on r.id = v.regimen
  inner join Documento d on v.doc_ingreso_id = d.id
where
  v.estado = 'activo'
  and r.estructura <= 2
  or r.estructura = 5 and p.dni = ?
ORDER by
  d.fecha_valida desc
limit
  1`

	err := p.db.GetContext(ctx, &status, query, dni)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Search_by_dni_vinculos implements PersonalRepository.
func (p *personalRepository) Search_by_dni_vinculos(ctx context.Context, dni string) (*[]models.Vinculos, error) {
	var resul []models.Vinculos

	err := p.db.SelectContext(ctx, &resul, "select * from Vinculos_vigentes where dni = ? order by fecha_ingreso desc", dni)
	if err != nil {
		return nil, err
	}
	return &resul, nil
}

// Search_by_dni_perfil implements PersonalRepository.
func (p *personalRepository) Search_by_dni_perfil(ctx context.Context, dni string) (*models.Perfil, error) {
	var res models.Perfil

	key := os.Getenv("DBKEY")

	query := `
	select
	p.dni,
	concat(p.apaterno, " ", p.amaterno, " ", p.nombre) nombre,
	cast(aes_decrypt(p.direccion,?) as char)  direccion,
	cast(aes_decrypt(p.telf1,?) as char)  telf1,
	cast(aes_decrypt(p.telf2,?) as char)  telf2,
	cast(aes_decrypt(p.email,?) as char)  email,
	p.ruc,
	p.fecha_nacimiento,
	p.sexo,
	p.foto
	from
	Vinculo v
	inner join Persona p on v.dni = p.dni
	where p.dni = ?
	GROUP by
	p.dni`
	err := p.db.GetContext(ctx, &res, query, key, key, key, key, dni)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func CreatePersonalRepository(db *sqlx.DB) PersonalRepository {
	return &personalRepository{db}
}
