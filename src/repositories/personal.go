package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"vesgoapp/src/models"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type PersonalRepository interface {
	Search_dni_onlinne(ctx context.Context, dni string) (*models.PersonDniRequest, error)
	Search_by_dni_perfil(ctx context.Context, dni string) (*models.Perfil, error)
	Search_by_dni_vinculos(ctx context.Context, dni string) (*[]models.Vinculos, error)
	Search_Areas_count(ctx context.Context, area string) (*[]models.AreaCargoSerch, error)
	Search_Cargo_count(ctx context.Context, cargo string) (*[]models.AreaCargoSerch, error)
	IsCasbyDni(ctx context.Context, dni string) (bool, error)
	EditByDni(ctx context.Context, telf1 string, telf2 string, direccion string, emai string, dni string, ruc string) error
	AddRenuncia(ctx context.Context, doc models.Documento, idvinculo int) (*int64, error)
	Buscar_Asistencia(ctx context.Context, dni string, mes int, año int) (*[]models.AsistenciaDb, error)
	Create_newIUser(ctx *sqlx.Tx, perfil models.Perfil) (*models.Perfil, error)
	Create_documento(ctx *sqlx.Tx, documento models.Documento) (*int64, error)
	Crear_Vinculo(ctx *sqlx.Tx, vinculo models.Vinculo) (*int64, error)
	TxBegin(ctx context.Context) (*sqlx.Tx, error)
}

type personalRepository struct {
	Db *sqlx.DB
}

func (p *personalRepository) Buscar_Asistencia(ctx context.Context, dni string, mes int, año int) (*[]models.AsistenciaDb, error) {

	var asistencia []models.AsistenciaDb

	query := `select * from asistenciavw where dni = ? and year(fecha) = ? and month(fecha) = ?`

	err := p.Db.SelectContext(ctx, &asistencia, query, dni, año, mes)
	if err != nil {
		return nil, err
	}
	return &asistencia, nil
}

func (p *personalRepository) TxBegin(ctx context.Context) (*sqlx.Tx, error) {
	tx, err := p.Db.BeginTxx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (p *personalRepository) Crear_Vinculo(ctx *sqlx.Tx, vinculo models.Vinculo) (*int64, error) {
	query := `INSERT INTO Vinculo (dni, doc_ingreso_id, doc_salida_id, estado, area_id, cargo_id, regimen) VALUES (?, ?, ?, ?, ?, ?, ?)`

	row, err := ctx.Exec(query, vinculo.Dni, vinculo.Dingreso, vinculo.Dsalida, vinculo.Estado, vinculo.Area, vinculo.Cargo, vinculo.Regimen)
	if err != nil {
		return nil, err
	}
	last, err := row.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &last, nil
}

func (p *personalRepository) Create_documento(ctx *sqlx.Tx, documento models.Documento) (*int64, error) {
	query := `INSERT INTO Documento (tipo, numero, year, fecha, fecha_valida, descripcion,sueldo) VALUES (?, ?, ?, ?, ?, ?, ?)`
	row, err := ctx.Exec(query, documento.Tipo, documento.Numero, documento.Año, documento.Fecha, documento.Fecha_Valida, documento.Descripcion, documento.Sueldo)
	if err != nil {
		if sqlErr, ok := err.(*pq.Error); ok && sqlErr.Code == "1062" {
			return nil, errors.New("el registro ya existe, no pueden existir duplicados")
		}
		return nil, err
	}
	last, err := row.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &last, nil
}

func (p *personalRepository) Create_newIUser(ctx *sqlx.Tx, perfil models.Perfil) (*models.Perfil, error) {
	key := os.Getenv("DBKEY")
	query := `INSERT INTO Persona (dni, nombre, apaterno, amaterno, direccion, telf1, telf2, email, ruc, fecha_nacimiento, sexo, foto) 
              VALUES (?, ?, ?, ?, aes_encrypt(?, ?), aes_encrypt(?, ?), aes_encrypt(?, ?), aes_encrypt(?, ?), ?, ?, ?, ?)`
	_, err := ctx.Exec(query, perfil.Dni, perfil.Nombre, perfil.Aparterno, perfil.Amaterno, perfil.Direccion, key, perfil.Telf1, key, perfil.Telf2, key, perfil.Email, key,
		perfil.Ruc, perfil.Nacimiento, perfil.Sexo, perfil.Foto)
	if err != nil {
		if sqlErr, ok := err.(*pq.Error); ok && sqlErr.Code == "1062" {
			return nil, errors.New("el registro ya existe, no pueden existir duplicados")
		}
		return nil, err
	}
	return &perfil, nil
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

	err := p.Db.SelectContext(ctx, &res, query, nombreLike)
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
	err := p.Db.SelectContext(ctx, &res, query, nombreLike)
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

	var resultado models.PersonDniRequest

	if err := json.Unmarshal(body, &resultado); err != nil {
		return nil, fmt.Errorf("error decodificando la respuesta: %w", err)
	}

	return &resultado, nil
}

func (p *personalRepository) AddRenuncia(ctx context.Context, doc models.Documento, idvinculo int) (*int64, error) {
	query := `insert into Documento (tipo, numero, year, fecha, fecha_valida, descripcion) values (?, ?, ?, ?, ?, ?)`

	last, err := p.Db.Exec(query, doc.Tipo, doc.Numero, doc.Año, doc.Fecha, doc.Fecha_Valida, doc.Descripcion)
	if err != nil {
		return nil, err
	}

	id, err := last.LastInsertId()
	if err != nil {
		return nil, err
	}

	query2 := `update Vinculo set estado = 'inactivo',doc_salida_id= ? where   id = ?`

	last, err = p.Db.Exec(query2, id, idvinculo)
	if err != nil {
		return nil, err
	}
	id2, err := last.RowsAffected()
	if err != nil {
		return nil, err
	}
	return &id2, nil
}

func (p *personalRepository) EditByDni(ctx context.Context, telf1 string, telf2 string, direccion string, emai string, ruc string, dni string) error {
	key := os.Getenv("DBKEY")
	query := `UPDATE Persona SET 
        direccion = aes_encrypt(?, ?), 
        telf1 = aes_encrypt(?, ?), 
        telf2 = aes_encrypt(?, ?), 
        email = aes_encrypt(?, ?), 
        ruc = ? 
        WHERE dni = ?`

	result, err := p.Db.ExecContext(ctx, query, direccion, key, telf1, key, telf2, key, emai, key, ruc, dni)
	if err != nil {

		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows were updated")
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

	err := p.Db.GetContext(ctx, &status, query, dni)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *personalRepository) Search_by_dni_vinculos(ctx context.Context, dni string) (*[]models.Vinculos, error) {
	var resul []models.Vinculos

	err := p.Db.SelectContext(ctx, &resul, "select * from Vinculos_vigentes where dni = ? order by fecha_ingreso desc", dni)
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
	err := p.Db.GetContext(ctx, &res, query, key, key, key, key, dni)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
func CreatePersonalRepository(db *sqlx.DB) PersonalRepository {
	return &personalRepository{db}
}
