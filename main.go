package main

import (
	"fmt"
	"net/http"
)

func main() {
	initDB()

	router := http.NewServeMux()

	// Crear
	router.HandleFunc("POST /personas", crearPersona)

	// Editar
	router.HandleFunc("PUT /personas", actualizarPersona)

	// Obtener
	router.HandleFunc("GET /personas/{id}", obtenerPersona)

	// Obtener todos
	//router.HandleFunc("GET /personas/", obtenerPersona)

	//Eliminar
	router.HandleFunc("DELETE /personas/{id}", eliminarPersona)

	// Iniciar el servidor HTTP en el puerto 8080
	fmt.Println(http.ListenAndServe(":8080", router))
}
