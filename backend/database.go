package backend

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	once     sync.Once
	instance *sqlx.DB
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// GetConnection retorna una única instancia de la conexión a la base de datos
func GetConnection(config Config) *sqlx.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.Database,
		)

		db, err := sqlx.Connect("mysql", dsn)
		if err != nil {
			log.Fatalf("Error conectando a la base de datos: %v", err)
		}

		// Configurar el pool de conexiones
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)

		instance = db
	})

	return instance
}

// CloseConnection cierra la conexión a la base de datos
func CloseConnection() {
	if instance != nil {
		instance.Close()
	}
}
