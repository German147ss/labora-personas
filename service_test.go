package main

import (
	"testing"
)

// Prueba para initDB
func TestInitDB(t *testing.T) {
	// Intenta inicializar la base de datos
	initDB()

	// Verifica si la conexión se estableció correctamente
	if DB == nil {
		t.Errorf("La conexión a la base de datos no se estableció correctamente")
	}

	// Intenta eliminar la persona de prueba si aún existe
	DB.Exec("DELETE FROM personas WHERE id = 1")
}

func TestEditarPersona(t *testing.T) {
	// Simular una base de datos con personas de prueba
	PersonasDB = []Persona{
		{ID: 1, Nombre: "Juan", Apellido: "Gómez", Edad: 30, CountryCode: "MX"},
		{ID: 2, Nombre: "María", Apellido: "López", Edad: 25, CountryCode: "US"},
	}

	// Editar la persona con ID 1
	editarPersona(1, "Pedro", "", 35, "ES")

	// Verificar si los cambios se realizaron correctamente
	if PersonasDB[0].Nombre != "Pedro" {
		t.Errorf("El nombre de la persona con ID 1 no se actualizó correctamente")
	}
	if PersonasDB[0].Edad != 35 {
		t.Errorf("La edad de la persona con ID 1 no se actualizó correctamente")
	}
	if PersonasDB[0].CountryCode != "ES" {
		t.Errorf("El countryCode de la persona con ID 1 no se actualizó correctamente")
	}
	if PersonasDB[0].Apellido != "Gómez" {
		t.Errorf("El Apellido de la persona con ID 1 se actualizó y no debería ser así.")
	}
}
