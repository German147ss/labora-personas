package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Función para conectar a la base de datos
func conectarDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=alfred dbname=labora host=localhost sslmode=disable password=4lfr3d port=5432")

	if err != nil {
		return nil, err
	}
	return db, nil
}

// Función para obtener una persona por su ID
func obtenerPersonaPorID(db *sql.DB, id int) (Persona, error) {
	var persona Persona
	row := db.QueryRow("SELECT * FROM personas WHERE id = $1", id)
	err := row.Scan(&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Edad, &persona.CountryCode)
	if err != nil {
		return Persona{}, err
	}
	return persona, nil
}
