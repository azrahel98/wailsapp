package models

import "database/sql"

type RegimenesCantidad struct {
	Cantidad int    `db:"cantidad"`
	Nombre   string `db:"nombre"`
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
	Dni     string   `db:"dni" json:"dni"`
	Nombre  string   `db:"nombre" json:"nombre"`
	Cargo   string   `db:"cargo" json:"area"`
	Sueldo  *float64 `db:"sueldo" json:"sueldo"`
	Regimen string   `db:"regimen" json:"reg"`
}
