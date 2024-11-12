package user

import "vesgoapp/backend"

type Usuario struct {
	ID   int    `db:"id"`
	Name string `db:"nombre"`
}

func ObtenerUsuarios() ([]Usuario, error) {
	db := backend.GetConnection(backend.Config{}) // No necesitas pasar config de nuevo
	var usuarios []Usuario
	err := db.Select(&usuarios, "SELECT id, nombre FROM usuario")
	return usuarios, err
}
