package models

import "database/sql"

type Perfil struct {
	Dni        string         `db:"dni"`
	Nombre     string         `db:"nombre"`
	Direccion  sql.NullString `db:"direccion"`
	Telf1      sql.NullString `db:"telf1"`
	Telf2      sql.NullString `db:"telf2"`
	Email      sql.NullString `db:"email"`
	Ruc        sql.NullString `db:"ruc"`
	Nacimiento string         `db:"fecha_nacimiento"`
	Sexo       sql.NullString `db:"sexo"`
	Foto       sql.NullString `db:"foto"`
}
