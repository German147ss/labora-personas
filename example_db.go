package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Persona struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Apellido    string `json:"apellido"`
	Edad        int    `json:"edad"`
	CountryCode string `json:"countryCode"`
}

func main() {
	// Configurar la conexión a PostgreSQL
	db, err := sql.Open("postgres", "user=alfred dbname=labora host=localhost sslmode=disable password=4lfr3d port=5431")
	if err != nil {
		fmt.Println("Error en la conexión a la base de datos")
		panic(err)
	}
	defer db.Close()
	fmt.Println("Conexión exitosa")

	var solicitudDeId = 2
	resultRow := db.QueryRow("SELECT * FROM personas where id = $1", solicitudDeId)
	if err != nil {
		fmt.Println("Error en la consulta")
		panic(err)
	}
	var persona Persona
	err = resultRow.Scan(&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Edad, &persona.CountryCode)
	if err != nil {
		fmt.Println("Error en el escaneo")
		panic(err)
	}
	fmt.Println(persona)

	defer db.Close()
}
