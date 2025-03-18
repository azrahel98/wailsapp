package models

import "database/sql"

type RegimenesCantidad struct {
	Cantidad int    `db:"cantidad"`
	Nombre   string `db:"nombre"`
}

type RegimenesResumen struct {
	Cantidad   int     `db:"cantidad" json:"cantidad"`
	Nombre     string  `db:"nombre" json:"nombre"`
	Porcentaje float64 `db:"porcentaje" json:"porcentaje"`
}

type Cumpleaños_Activos struct {
	Dni        string `db:"dni"`
	Nombre     string `db:"nombre"`
	Nacimiento string `db:"nacimiento"`
	Edad       string `db:"edad"`
}

type Buscar_trabajador struct {
	Dni    string         `db:"dni"`
	Nombre string         `db:"nombre"`
	Foto   sql.NullString `db:"foto"`
	Estado string         `db:"estado"`
	Sexo   string         `db:"sexo"`
}

type PersonaActivo struct {
	Cantidad int `db:"cantidad"`
	Activos  int `db:"activos"`
}

type TrabajadoresArea struct {
	Dni       string   `db:"dni" json:"dni"`
	Nombre    string   `db:"nombre" json:"nombre"`
	Cargo     string   `db:"cargo" json:"area"`
	Sueldo    *float64 `db:"sueldo" json:"sueldo"`
	Regimen   string   `db:"regimen" json:"reg"`
	Sindicato *string  `db:"sindicato" json:"sindicato"`
}

type ResumenIndicadores struct {
	Personalregistrado PersonaActivo       `json:"personalregistrado"`
	RenunciasMes       []PersonaActivo     `json:"renunciasmes"`
	Regimenes          []RegimenesCantidad `json:"regimenescantidad"`
	Sexo               []RegimenesCantidad `json:"sexo"`
	Sindicatos         []RegimenesCantidad `json:"sindicatos"`
}

type Reporte_Trabajadores struct {
	Dni       string  `db:"dni" json:"dni"`
	Nombre    string  `db:"nombre" json:"nombre"`
	Ingreso   string  `db:"ingreso" json:"ingreso"`
	Renuncia  *string `db:"renuncia" json:"renuncia"`
	Area      string  `db:"area" json:"area"`
	Cargo     string  `db:"cargo" json:"cargo"`
	Sindicato *string `db:"sindicato" json:"sindicato"`
	Regimen   string  `db:"regimen" json:"regimen"`
}

type Reporte_Funcionarios struct {
	Id     int     `db:"id" json:"id"`
	Area   string  `db:"area" json:"area"`
	Nombre *string `db:"nombre" json:"nombre"`
	Nivel  *int    `db:"nivel" json:"nivel"`
}
