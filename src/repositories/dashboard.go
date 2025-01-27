package repositories

import (
	"context"
	"fmt"
	"vesgoapp/src/models"

	"github.com/jmoiron/sqlx"
)

type DashboardRepository interface {
	Cantidad_by_regimen(ctx context.Context) (*[]models.RegimenesCantidad, error)
	Cantidad_by_sexo(ctx context.Context) (*[]models.RegimenesCantidad, error)
	Cumpleaño_by_activos(ctx context.Context, mes int) (*[]models.Cumpleaños_Activos, error)
	Buscar_by_nombre(ctx context.Context, nombre string) (*[]models.Buscar_trabajador, error)
	Trabajadores_activos(ctx context.Context) (*[]models.RegimenesCantidad, error)
	Cantidad_vincolos_activos(ctx context.Context) (*models.PersonaActivo, error)
	Cantidad_renuncias_mes(ctx context.Context) (*[]models.PersonaActivo, error)
	Trabajadores_area(ctx context.Context, nombre string) (*[]models.TrabajadoresArea, error)
}
type dashboardRepository struct {
	db *sqlx.DB
}

// Trabajadores_area implements DashboardRepository.
func (d *dashboardRepository) Trabajadores_area(ctx context.Context, nombre string) (*[]models.TrabajadoresArea, error) {
	query := `select
   cast(p.dni as char) dni,
  concat(p.apaterno,' ',p.amaterno,' ',p.nombre) nombre,
  cr.nombre cargo,
  d.sueldo,
  r.nombre regimen
from
  Vinculo v
  inner join Area ar on v.area_id = ar.id
  inner join Persona p on v.dni = p.dni
  inner join Cargo cr on v.cargo_id = cr.id
  inner join Documento d on v.doc_ingreso_id = d.id
    INNER JOIN Regimen r ON v.regimen = r.id
where
  v.estado = 'activo'
  and ar.nombre = ? order by d.sueldo desc`

	res := []models.TrabajadoresArea{}

	err := d.db.SelectContext(ctx, &res, query, nombre)
	if err != nil {

		return nil, err
	}
	return &res, nil
}

func (d *dashboardRepository) Cantidad_renuncias_mes(ctx context.Context) (*[]models.PersonaActivo, error) {
	query := `select
  count(d.id) cantidad,
  month(d.fecha) activos
from
  Vinculo v
  inner join Documento d on v.doc_ingreso_id = d.id
WHERE
  year(d.fecha) = year(now())
GROUP by
  month(d.fecha)`
	res := []models.PersonaActivo{}

	err := d.db.SelectContext(ctx, &res, query)
	if err != nil {
		fmt.Println("aquii tenemos el error señor")
		fmt.Println(err)
		return nil, err
	}
	return &res, nil
}

func (d *dashboardRepository) Cantidad_vincolos_activos(ctx context.Context) (*models.PersonaActivo, error) {
	query := `select
count(*) cantidad,
(
  select
	count(*) activos
  from
	Vinculo
  where
	estado = 'activo'
) activos
from
Vinculo`
	var res models.PersonaActivo

	err := d.db.GetContext(ctx, &res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (d *dashboardRepository) Trabajadores_activos(ctx context.Context) (*[]models.RegimenesCantidad, error) {
	var res []models.RegimenesCantidad

	query := `
      select
  ar.nombre nombre,
  count(v.area_id) cantidad
from
  Vinculo v
  inner join Area ar on v.area_id = ar.id
WHERE
  v.estado = 'activo'
  and ar.activo = 1
GROUP by
  ar.id order by
  cantidad desc

    `

	err := d.db.SelectContext(ctx, &res, query)
	if err != nil {
		return nil, err
	}

	return &res, nil
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
