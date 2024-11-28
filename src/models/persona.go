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

type Vinculos struct {
	Id            *int    `db:"id"`
	Dni           *string `db:"dni"`
	Doc_i         *string `db:"doc_ingreso"`
	Numero_doc_i  *string `db:"numero_doc_ingreso"`
	Descrip_ingre *string `db:"descrip_ingreso"`
	Fecha_ingreso *string `db:"fecha_ingreso"`
	Area          *string `db:"area"`
	Cargo         *string `db:"cargo"`
	Regimen       *string `db:"regimen"`
	Sueldo        *int    `db:"sueldo"`
	Estado        *string `db:"estado"`
	Doc_s         *string `db:"doc_salida"`
	Fecha_salida  *string `db:"fecha_salida"`
	Descrip_s     *string `db:"descrip_salida"`
	Numero_doc_s  *string `db:"numero_doc_salida"`
}
