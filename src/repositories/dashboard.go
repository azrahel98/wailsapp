package repositories

import (
	"context"
	"vesgoapp/src/models"

	"github.com/jmoiron/sqlx"
)

type DashboardRepository interface {
	Cantidad_by_regimen(ctx context.Context) (*[]models.RegimenesCantidad, error)
	Cantidad_by_sexo(ctx context.Context) (*[]models.RegimenesCantidad, error)
	Cumpleaño_by_activos(ctx context.Context, mes int) (*[]models.Cumpleaños_Activos, error)
	Buscar_by_nombre(ctx context.Context, nombre string) (*[]models.Buscar_trabajador, error)
}
type dashboardRepository struct {
	db *sqlx.DB
}

func (d *dashboardRepository) Cantidad_by_sexo(ctx context.Context) (*[]models.RegimenesCantidad, error) {
	var res []models.RegimenesCantidad

	query := `
       select
  count(v.id) cantidad,
  p.sexo nombre
from
  Vinculo v
  inner join Persona p on p.dni = v.dni
WHERE
  v.estado = "activo"
GROUP by
  p.sexo

    `

	err := d.db.SelectContext(ctx, &res, query)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (d *dashboardRepository) Buscar_by_nombre(ctx context.Context, nombre string) (*[]models.Buscar_trabajador, error) {
	var res []models.Buscar_trabajador
	query := `select
		concat(dg.nombre, ' ', dg.apaterno, ' ', dg.amaterno) nombre,
		v.dni,
		foto,
		MIN(v.estado) AS estado,
		dg.sexo sexo
		from
		Persona dg
		inner join Vinculo v on dg.dni = v.dni 
		WHERE
		concat(dg.apaterno, " ", dg.amaterno, " ", dg.nombre) LIKE ?
		GROUP BY
		v.dni
	`
	nombreLike := "%" + nombre + "%"
	err := d.db.SelectContext(ctx, &res, query, nombreLike)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (d *dashboardRepository) Cumpleaño_by_activos(ctx context.Context, mes int) (*[]models.Cumpleaños_Activos, error) {
	var res []models.Cumpleaños_Activos
	query := `
	select
	p.dni,
	concat(p.apaterno, " ", amaterno, " ", p.nombre) nombre,
	p.fecha_nacimiento nacimiento,
	YEAR(CURRENT_DATE) - year(p.fecha_nacimiento) edad
	from
	Persona p
	inner join Vinculo v on p.dni = v.dni
	where
	v.estado = 'activo' and month(p.fecha_nacimiento) = ?
	GROUP by
	p.dni
	`
	err := d.db.SelectContext(ctx, &res, query, mes)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (d *dashboardRepository) Cantidad_by_regimen(ctx context.Context) (*[]models.RegimenesCantidad, error) {
	var res []models.RegimenesCantidad

	query := `
        select
		count(v.id) cantidad,
		r.decreto nombre
		from
		Vinculo v
		inner join Regimen r on v.regimen = r.id
		WHERE
		v.estado = "activo"
		GROUP by
		r.estructura
		ORDER by
		r.nombre
    `

	err := d.db.SelectContext(ctx, &res, query)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func CreatedashboardRepository(db *sqlx.DB) DashboardRepository {
	return &dashboardRepository{db}
}
