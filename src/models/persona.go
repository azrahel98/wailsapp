package models

type Perfil struct {
	Dni        string  `db:"dni"`
	Nombre     string  `db:"nombres"`
	Aparterno  string  `db:"apaterno"`
	Amaterno   string  `db:"amaterno"`
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

type Documento struct {
	Id           *int    `db:"id" json:"id"`
	Tipo         string  `db:"tipo" json:"tipo"`
	Numero       *int    `db:"numero" json:"numero"`
	Año          *int    `db:"year" json:"año"`
	Fecha        string  `db:"fecha" json:"fecha"`
	Fecha_Valida *string `db:"fecha_valida" json:"fecha_valida"`
	Conv         *int64  `db:"conv" json:"conv"`
	Descripcion  string  `db:"descripcion" json:"descripcion"`
	Funcion      *int64  `db:"funcion" json:"funcion"`
}

type PersonDniRequest struct {
	Nombres  string `json:"nombres"`
	Apaterno string `json:"apellidoPaterno"`
	Amaterno string `json:"apellidoMaterno"`
}

type AreaCargoSerch struct {
	Cantidad int    `db:"cantidad" json:"cantidad"`
	Nombre   string `db:"nombre" json:"nombre"`
	Id       int    `db:"id" json:"id"`
}
