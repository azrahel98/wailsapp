package models

type Perfil struct {
	Dni        string  `db:"dni"`
	Nombre     string  `db:"nombre"`
	Direccion  *string `db:"direccion"`
	Telf1      *string `db:"telf1"`
	Telf2      *string `db:"telf2"`
	Email      *string `db:"email"`
	Ruc        *string `db:"ruc"`
	Nacimiento string  `db:"fecha_nacimiento"`
	Sexo       *string `db:"sexo"`
	Foto       *string `db:"foto"`
}
