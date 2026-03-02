package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

func Connect() {

	connString := "server=MSI;database=EvaluacionFinal;trusted_connection=yes"

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error conectando:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a SQL Server:", err)
	}

	DB = db
	fmt.Println("🔥 Conectado a SQL Server correctamente")
}
