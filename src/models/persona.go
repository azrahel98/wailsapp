package models

type Perfil struct {
	Dni        string  `db:"dni" json:"Dni"`
	Nombre     string  `db:"nombres" json:"Nombre"`
	Aparterno  string  `db:"apaterno" json:"Apaterno"`
	Amaterno   string  `db:"amaterno" json:"Amaterno"`
	Direccion  *string `db:"direccion" json:"Direccion"`
	Telf1      *string `db:"telf1" json:"Telf1"`
	Telf2      *string `db:"telf2" json:"Telf2"`
	Email      *string `db:"email" json:"Email"`
	Ruc        *string `db:"ruc" json:"Ruc"`
	Nacimiento string  `db:"fecha_nacimiento" json:"Nacimiento"`
	Sexo       *string `db:"sexo" json:"Sexo"`
	Foto       *string `db:"foto" json:"Foto"`
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
	Tipo         string  `db:"tipo" json:"tipoDocumento"`
	Numero       *int    `db:"numero" json:"numeroDocumento"`
	Año          *int    `db:"year" json:"añoDocumento"`
	Fecha        string  `db:"fecha" json:"fecha"`
	Fecha_Valida *string `db:"fecha_valida" json:"fechaValida"`
	Conv         *int64  `db:"conv" json:"conv"`
	Descripcion  string  `db:"descripcion" json:"descripcion"`
	Funcion      *int64  `db:"funcion" json:"funcion"`
	Sueldo       float64 `db:"sueldo" json:"sueldo"`
}

type Vinculo struct {
	Id       *int   `db:"id" json:"id"`
	Dni      string `db:"dni" json:"dni"`
	Dingreso int    `db:"doc_ingreso_id" json:"doc_ingreso"`
	Dsalida  *int   `db:"doc_salida_id" json:"doc_salida"`
	Estado   string `db:"estado" json:"estado"`
	Area     int    `db:"area_id" json:"area"`
	Cargo    int    `db:"cargo_id" json:"cargo"`
	Plaza    *int   `db:"plaza_id" json:"plaza"`
	Regimen  int    `db:"regimen" json:"regimen"`
}

type Data struct {
	Persona   Perfil    `json:"persona"`
	Documento Documento `json:"documento"`
	Vinculo   Vinculo   `json:"vinculo"`
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
