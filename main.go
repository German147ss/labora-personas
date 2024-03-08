package main

import (
	"fmt"
	"net/http"
)


func main() {
	p := Persona{
		1,
		"Juan",
		"Arbelaez",
		20,
		"col",
	}
	insertarPersonaEnLaDb(p)

	router := http.NewServeMux()

	// Obtener
	router.HandleFunc("/personas", obtenerPersona)

	// Iniciar el servidor HTTP en el puerto 8080
	fmt.Println(http.ListenAndServe(":8086", router))
}
