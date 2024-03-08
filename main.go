package main

import (
	"net/http"
)

func main() {

	router := http.NewServeMux()

	// Crear
	router.HandleFunc("POST /personas", crearPersona)

	// Editar
	router.HandleFunc("PUT /personas", actualizarPersona)

	// Obtener
	router.HandleFunc("GET /personas/{id}", obtenerPersona)

	// Iniciar el servidor HTTP en el puerto 8080
	http.ListenAndServe(":8080", router)
}
