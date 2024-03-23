package main

import "testing"

func TestObtenerPersonaPorID(t *testing.T) {
	// Conectar a la base de datos
	db, err := conectarDB()
	if err != nil {
		t.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Agregar una persona de prueba a la base de datos
	_, err = db.Exec("INSERT INTO personas (id, nombre, apellido, edad, countryCode) VALUES ($1, $2, $3, $4, $5)", 1, "John", "Doe", 30, "US")
	if err != nil {
		t.Fatalf("Error al agregar una persona de prueba: %v", err)
	}

	// Obtener la persona recién agregada por su ID
	persona, err := obtenerPersonaPorID(db, 1)
	if err != nil {
		t.Fatalf("Error al obtener la persona por ID: %v", err)
	}

	// Verificar que los detalles de la persona sean correctos
	expectedPersona := Persona{ID: 1, Nombre: "John", Apellido: "Doe", Edad: 30, CountryCode: "US"}
	if persona != expectedPersona {
		t.Errorf("Persona obtenida incorrecta. Se esperaba %v pero se obtuvo %v", expectedPersona, persona)
	}

	// Eliminar la persona de prueba de la base de datos
	_, err = db.Exec("DELETE FROM personas WHERE id = $1", 1)
}

func TestConectarDB(t *testing.T) {
	// Conectar a la base de datos
	db, err := conectarDB()
	if err != nil {
		t.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("Error al hacer ping a la base de datos: %v", err)
	}
}
