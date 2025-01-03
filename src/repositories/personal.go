package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"vesgoapp/src/models"

	"github.com/jmoiron/sqlx"
)

type PersonalRepository interface {
	Search_dni_onlinne(ctx context.Context, dni string) (*models.PersonDniRequest, error)
	Search_by_dni_perfil(ctx context.Context, dni string) (*models.Perfil, error)
	Search_by_dni_vinculos(ctx context.Context, dni string) (*[]models.Vinculos, error)
	Search_Areas_count(ctx context.Context, area string) (*[]models.AreaCargoSerch, error)
	Search_Cargo_count(ctx context.Context, cargo string) (*[]models.AreaCargoSerch, error)
	IsCasbyDni(ctx context.Context, dni string) (bool, error)
	EditByDni(ctx context.Context, telf1 string, telf2 string, direccion string, emai string, dni string) error
	AddRenuncia(ctx context.Context, doc models.Documento, idvinculo int) (*int64, error)
}

type personalRepository struct {
	db *sqlx.DB
}

func (p *personalRepository) Search_Cargo_count(ctx context.Context, cargo string) (*[]models.AreaCargoSerch, error) {
	var res []models.AreaCargoSerch
	query := `SELECT
	COUNT(*) AS cantidad,
	ar.nombre,
	ar.id
  FROM
	Vinculo v
	INNER JOIN Cargo ar ON v.cargo_id = ar.id
  WHERE
	v.estado = 'activo' AND ar.activo = 1 AND ar.nombre LIKE ?
  GROUP BY
	v.cargo_id`
	nombreLike := "%" + cargo + "%"

	err := p.db.SelectContext(ctx, &res, query, nombreLike)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (p *personalRepository) Search_Areas_count(ctx context.Context, area string) (*[]models.AreaCargoSerch, error) {
	var res []models.AreaCargoSerch
	query := `select
  count(*) as cantidad,
  ar.nombre,
  ar.id
from
  Vinculo v
  inner join Area ar on v.area_id = ar.id
  where v.estado = 'activo' and ar.activo = 1 and ar.nombre like ?
GROUP by
  v.area_id`
	nombreLike := "%" + area + "%"
	err := p.db.SelectContext(ctx, &res, query, nombreLike)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (p *personalRepository) Search_dni_onlinne(ctx context.Context, dni string) (*models.PersonDniRequest, error) {
	url := fmt.Sprintf("https://api.apis.net.pe/v2/reniec/dni?numero=%s", dni)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+models.APIKEY)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error leyendo la respuesta: %w", err)
	}

	fmt.Println("Cuerpo de la respuesta:", string(body))
	var resultado models.PersonDniRequest

	if err := json.Unmarshal(body, &resultado); err != nil {
		return nil, fmt.Errorf("error decodificando la respuesta: %w", err)
	}

	return &resultado, nil
}

func (p *personalRepository) AddRenuncia(ctx context.Context, doc models.Documento, idvinculo int) (*int64, error) {
	query := `insert into Documento (tipo, numero, year, fecha, fecha_valida, descripcion) values (?, ?, ?, ?, ?, ?)`

	last, err := p.db.Exec(query, doc.Tipo, doc.Numero, doc.Año, doc.Fecha, doc.Fecha_Valida, doc.Descripcion)
	if err != nil {
		return nil, err
	}

	id, err := last.LastInsertId()
	if err != nil {
		return nil, err
	}

	query2 := `update Vinculo set estado = 'inactivo',doc_salida_id= ? where   id = ?`

	last, err = p.db.Exec(query2, id, idvinculo)
	if err != nil {
		return nil, err
	}
	id2, err := last.RowsAffected()
	if err != nil {
		return nil, err
	}
	return &id2, nil
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

func (p *personalRepository) Search_by_dni_vinculos(ctx context.Context, dni string) (*[]models.Vinculos, error) {
	var resul []models.Vinculos

	err := p.db.SelectContext(ctx, &resul, "select * from Vinculos_vigentes where dni = ? order by fecha_ingreso desc", dni)
	if err != nil {
		return nil, err
	}
	return &resul, nil
}

func (p *personalRepository) Search_by_dni_perfil(ctx context.Context, dni string) (*models.Perfil, error) {
	var res models.Perfil

	key := os.Getenv("DBKEY")

	query := `
	select
	p.dni,
	p.nombre nombres,
	p.apaterno,
	p.amaterno,
	cast(aes_decrypt(p.direccion,?) as char)  direccion,
	cast(aes_decrypt(p.telf1,?) as char)  telf1,
	cast(aes_decrypt(p.telf2,?) as char)  telf2,
	cast(aes_decrypt(p.email,?) as char)  email,
	p.ruc,
	p.fecha_nacimiento,
	p.sexo,
	null as foto
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
